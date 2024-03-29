package batchapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

// Test for NewBatchRequest
func TestNewBatchRequest(t *testing.T) {
	client := new(MockClient)
	batchReq := NewBatchRequest(client)
	assert.NotNil(t, batchReq)
}

// Test for AddRequest
func TestAddRequest(t *testing.T) {
	client := new(MockClient)

	client.On("GetBaseURL").Return("https://instance.service-now.com")

	batchReq := NewBatchRequest(client)

	mockReqInfo := new(mocking.MockRequestInformation)
	mockReqInfo.On("Url").Return("http://example.com", nil)
	mockReqInfo.On("GetContent").Return([]byte(`{"key":"value"}`))
	mockReqInfo.On("GetMethod").Return("GET")

	headers := internal.NewRequestHeader()
	headers.Set("Content-Type", "application/json")

	mockReqInfo.On("GetHeaders").Return(headers)

	err := batchReq.AddRequest(mockReqInfo, false)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(batchReq.(*batchRequest).Requests))
}

// Test for toBatchItem
func TestToBatchItem(t *testing.T) {
	client := new(MockClient)
	batchReq := NewBatchRequest(client).(*batchRequest)

	var mockReqInfo *mocking.MockRequestInformation
	var headers core.RequestHeader

	tests := []test[[]interface{}]{
		{
			title: "Successful",
			setup: func() {
				client.On("GetBaseURL").Return("https://instance.service-now.com")
				mockReqInfo = new(mocking.MockRequestInformation)
				mockReqInfo.On("Url").Return("table/test", nil)
				mockReqInfo.On("GetContent").Return([]byte(`{"key":"value"}`))
				mockReqInfo.On("GetMethod").Return("GET")

				headers = internal.NewRequestHeader()
				headers.Set("Content-Type", "application/json")

				mockReqInfo.On("GetHeaders").Return(headers)
			},
			expected:    []interface{}{"GET", "/api/table/test"},
			expectedErr: nil,
		},
		{
			title: "Url returns error",
			setup: func() {
				mockReqInfo = new(mocking.MockRequestInformation)
				mockReqInfo.On("Url").Return("", errors.New("url error"))
			},
			expected:    nil,
			expectedErr: errors.New("url error"),
		},
		{
			title: "GetContent returns invalid JSON",
			setup: func() {
				mockReqInfo = new(mocking.MockRequestInformation)
				mockReqInfo.On("Url").Return("table/test", nil)
				mockReqInfo.On("GetContent").Return([]byte(`{"key":`)) // invalid JSON
				mockReqInfo.On("GetMethod").Return("GET")

				headers = internal.NewRequestHeader()
				headers.Set("Content-Type", "application/json")

				mockReqInfo.On("GetHeaders").Return(headers)
			},
			expected:    nil,
			expectedErr: errors.New("unexpected end of JSON input"), // error message from json.Unmarshal
		},
		{
			title: "Uri has base url as prefix",
			setup: func() {
				mockReqInfo = new(mocking.MockRequestInformation)
				mockReqInfo.On("Url").Return(client.GetBaseURL()+"/table/test", nil)
				mockReqInfo.On("GetContent").Return([]byte(`{"key":"value"}`))
				mockReqInfo.On("GetMethod").Return("GET")

				headers = internal.NewRequestHeader()
				headers.Set("Content-Type", "application/json")

				mockReqInfo.On("GetHeaders").Return(headers)
			},
			expected:    []interface{}{"GET", "/api/table/test"},
			expectedErr: nil,
		},
		{
			title: "bad base url",
			setup: func() {
				client.On("GetBaseURL").Unset()
				client.On("GetBaseURL").Return("http://a b.com")

				mockReqInfo = new(mocking.MockRequestInformation)
				mockReqInfo.On("Url").Return("table/test", nil)
				mockReqInfo.On("GetContent").Return([]byte(`{"key":"value"}`))
				mockReqInfo.On("GetMethod").Return("GET")

				headers = internal.NewRequestHeader()
				headers.Set("Content-Type", "application/json")

				mockReqInfo.On("GetHeaders").Return(headers)
			},
			expected:    nil,
			expectedErr: errors.New("parse \"http://a b.com\": invalid character \" \" in host name"),
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			if test.setup != nil {
				test.setup()
			}

			item, err := batchReq.toBatchItem(mockReqInfo, false)

			if test.expectedErr != nil {
				assert.ErrorContains(t, err, test.expectedErr.Error())
			} else {
				assert.Nil(t, test.expectedErr, err)
			}

			if test.expected != nil {
				method := test.expected[0].(string)
				url := test.expected[1].(string)

				assert.NotNil(t, item)
				assert.Equal(t, method, *item.GetMethod())
				assert.Equal(t, url, *item.GetURL())
				assert.Equal(t, headers, item.GetHeaders())
			} else {
				assert.Nil(t, item)
			}

			if test.cleanup != nil {
				test.cleanup()
			}
		})
	}
}
