package tableapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/table-api/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewTableRequestBuilder2(t *testing.T) {
	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Input: []interface{}{&mockClient{}, map[string]string{"baseurl": "baseurl", "table": "table"}},
		},
		{
			Title:       "missing table",
			Input:       []interface{}{&mockClient{}, map[string]string{"baseurl": "baseurl"}},
			ExpectedErr: errors.New("missing \"table\" parameter"),
		},
		{
			Title:       "missing baseurl",
			Input:       []interface{}{&mockClient{}, map[string]string{}},
			ExpectedErr: errors.New("pathParameters must contain a value for \"baseurl\" for the URL to be built"),
		},
		{
			Title:       "missing client",
			Input:       []interface{}{(*mockClient)(nil), map[string]string{}},
			ExpectedErr: errors.New("client can't be nil"),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			inputs := test.Input.([]interface{})

			_, err := NewTableRequestBuilder2(inputs[0].(core.Client), inputs[1].(map[string]string))
			assert.Equal(t, test.ExpectedErr, err)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

//nolint:dupl
func TestTableRequestBuilder2_Get(t *testing.T) {
	requestBuilder := &tableRequestBuilder2{RequestBuilder: &mockRequestBuilder{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				mockRB := &mockRequestBuilder{}
				mockRB.On("SendGet2", mock.AnythingOfType("*core.RequestConfiguration")).Return(nil)

				requestBuilder.RequestBuilder = mockRB
			},
			ExpectedErr: nil,
		},
		{
			Title: "Error",
			Setup: func() {
				mockRB := &mockRequestBuilder{}
				mockRB.On("SendGet2", mock.AnythingOfType("*core.RequestConfiguration")).Return(errors.New("unable to send request"))

				requestBuilder.RequestBuilder = mockRB
			},
			ExpectedErr: errors.New("unable to send request"),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			_, err := requestBuilder.Get(nil)
			assert.Equal(t, test.ExpectedErr, err)
			requestBuilder.RequestBuilder.(*mockRequestBuilder).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

func TestTableRequestBuilder2_Post(t *testing.T) {
	requestBuilder := &tableRequestBuilder2{RequestBuilder: &mockRequestBuilder{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				mockRB := &mockRequestBuilder{}
				mockRB.On("SendPost3", mock.AnythingOfType("*core.RequestConfiguration")).Return(nil)

				requestBuilder.RequestBuilder = mockRB
			},
			ExpectedErr: nil,
			Input:       NewTableEntry(),
		},
		{
			Title: "Error",
			Setup: func() {
				mockRB := &mockRequestBuilder{}
				mockRB.On("SendPost3", mock.AnythingOfType("*core.RequestConfiguration")).Return(errors.New("unable to send request"))

				requestBuilder.RequestBuilder = mockRB
			},
			ExpectedErr: errors.New("unable to send request"),
			Input:       NewTableEntry(),
		},
		{
			Title: "Nil tableEntry",
			Setup: func() {
				mockRB := &mockRequestBuilder{}

				requestBuilder.RequestBuilder = mockRB
			},
			Input:       nil,
			ExpectedErr: errors.New("entry is nil"),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			_, err := requestBuilder.Post(test.Input, nil)
			assert.Equal(t, test.ExpectedErr, err)
			requestBuilder.RequestBuilder.(*mockRequestBuilder).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}
