package batchapi

import (
	"errors"
	"testing"

	intBatch "github.com/RecoLabs/servicenow-sdk-go/batch-api/internal"
	"github.com/RecoLabs/servicenow-sdk-go/internal"
	"github.com/RecoLabs/servicenow-sdk-go/internal/core"
	"github.com/RecoLabs/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

// Test for NewBatchRequest
func TestNewBatchRequest(t *testing.T) {
	client := new(intBatch.MockClient)
	batchReq := NewBatchRequest(client)
	assert.NotNil(t, batchReq)
}

// Test for AddRequest
func TestAddRequest(t *testing.T) {
	client := new(intBatch.MockClient)
	mockReqInfo := new(mocking.MockRequestInformation)

	batchReq := NewBatchRequest(client)

	tests := []intBatch.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				client.On("GetBaseURL").Return("https://instance.service-now.com")

				headers := internal.NewRequestHeader()
				headers.Set("Content-Type", "application/json")

				mockReqInfo.On("Url").Return("http://example.com", nil)
				mockReqInfo.On("GetContent").Return([]byte(`{"key":"value"}`))
				mockReqInfo.On("GetMethod").Return("GET")
				mockReqInfo.On("GetHeaders").Return(headers)
			},
			Cleanup: func() {
				intBatch.ResetCalls(append(mockReqInfo.ExpectedCalls, client.ExpectedCalls...)...)
			},
			ExpectedErr: nil,
		},
		{
			Title: "URL Error",
			Setup: func() {
				client.On("GetBaseURL").Return("https://instance.service-now.com")

				headers := internal.NewRequestHeader()
				headers.Set("Content-Type", "application/json")

				mockReqInfo.On("Url").Return("", errors.New("unable to parse URL"))
				mockReqInfo.On("GetContent").Return([]byte(`{"key":"value"}`))
				mockReqInfo.On("GetMethod").Return("GET")
				mockReqInfo.On("GetHeaders").Return(headers)
			},
			Cleanup: func() {
				intBatch.ResetCalls(append(mockReqInfo.ExpectedCalls, client.ExpectedCalls...)...)
			},
			ExpectedErr: errors.New("unable to parse URL"),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			err := batchReq.AddRequest(mockReqInfo, false)
			assert.Equal(t, test.ExpectedErr, err)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

// Test for toBatchItem
func TestToBatchItem(t *testing.T) {
	client := new(intBatch.MockClient)
	batchReq := NewBatchRequest(client).(*batchRequest)

	var mockReqInfo *mocking.MockRequestInformation
	var headers core.RequestHeader

	tests := []intBatch.Test[[]interface{}]{
		{
			Title: "Successful",
			Setup: func() {
				client.On("GetBaseURL").Return("https://instance.service-now.com")
				mockReqInfo = new(mocking.MockRequestInformation)
				mockReqInfo.On("Url").Return("table/test", nil)
				mockReqInfo.On("GetContent").Return([]byte(`{"key":"value"}`))
				mockReqInfo.On("GetMethod").Return("GET")

				headers = internal.NewRequestHeader()
				headers.Set("Content-Type", "application/json")

				mockReqInfo.On("GetHeaders").Return(headers)
			},
			Expected:    []interface{}{"GET", "/api/table/test"},
			ExpectedErr: nil,
		},
		{
			Title: "Url returns error",
			Setup: func() {
				mockReqInfo = new(mocking.MockRequestInformation)
				mockReqInfo.On("Url").Return("", errors.New("url error"))
			},
			Expected:    nil,
			ExpectedErr: errors.New("url error"),
		},
		{
			Title: "GetContent returns invalid JSON",
			Setup: func() {
				mockReqInfo = new(mocking.MockRequestInformation)
				mockReqInfo.On("Url").Return("table/test", nil)
				mockReqInfo.On("GetContent").Return([]byte(`{"key":`)) // invalid JSON
				mockReqInfo.On("GetMethod").Return("GET")

				headers = internal.NewRequestHeader()
				headers.Set("Content-Type", "application/json")

				mockReqInfo.On("GetHeaders").Return(headers)
			},
			Expected:    nil,
			ExpectedErr: errors.New("unexpected end of JSON input"), // error message from json.Unmarshal
		},
		{
			Title: "Uri has base url as prefix",
			Setup: func() {
				mockReqInfo = new(mocking.MockRequestInformation)
				mockReqInfo.On("Url").Return(client.GetBaseURL()+"/table/test", nil)
				mockReqInfo.On("GetContent").Return([]byte(`{"key":"value"}`))
				mockReqInfo.On("GetMethod").Return("GET")

				headers = internal.NewRequestHeader()
				headers.Set("Content-Type", "application/json")

				mockReqInfo.On("GetHeaders").Return(headers)
			},
			Expected:    []interface{}{"GET", "/api/table/test"},
			ExpectedErr: nil,
		},
		{
			Title: "bad base url",
			Setup: func() {
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
			Expected:    nil,
			ExpectedErr: errors.New("parse \"http://a b.com\": invalid character \" \" in host name"),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			item, err := batchReq.toBatchItem(mockReqInfo, false)

			if test.ExpectedErr != nil {
				assert.ErrorContains(t, err, test.ExpectedErr.Error())
			} else {
				assert.Nil(t, test.ExpectedErr, err)
			}

			if test.Expected != nil {
				method := test.Expected[0].(string)
				url := test.Expected[1].(string)

				assert.NotNil(t, item)
				assert.Equal(t, method, *item.GetMethod())
				assert.Equal(t, url, *item.GetURL())
				assert.Equal(t, headers, item.GetHeaders())
			} else {
				assert.Nil(t, item)
			}

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}
