package tableapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/table-api/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRequestBuilder struct {
	mock.Mock
}

func (rB *mockRequestBuilder) SendDelete2(config *core.RequestConfiguration) error {
	args := rB.Called(config)
	return args.Error(0)
}

func (rB *mockRequestBuilder) SendGet2(config *core.RequestConfiguration) error {
	args := rB.Called(config)
	return args.Error(0)
}

func (rB *mockRequestBuilder) SendPost3(config *core.RequestConfiguration) error {
	args := rB.Called(config)
	return args.Error(0)
}

func (rB *mockRequestBuilder) SendPut2(config *core.RequestConfiguration) error {
	args := rB.Called(config)
	return args.Error(0)
}

func (rB *mockRequestBuilder) ToHeadRequestInformation() (*core.RequestInformation, error) {
	args := rB.Called()
	return args.Get(0).(*core.RequestInformation), args.Error(1)
}

func TestNewTableItemRequestBuilder2(t *testing.T) {
	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Input: []interface{}{&mockClient{}, map[string]string{"baseurl": "baseurl", "table": "table", "sysId": "fdafsdfdsa"}},
		},
		{
			Title:       "missing sysId",
			Input:       []interface{}{&mockClient{}, map[string]string{"baseurl": "baseurl", "table": "table"}},
			ExpectedErr: errors.New("missing \"sysId\" parameter"),
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

			_, err := NewTableItemRequestBuilder2(inputs[0].(core.Client), inputs[1].(map[string]string))
			assert.Equal(t, test.ExpectedErr, err)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

//nolint:dupl
func TestTableItemRequestBuilder2_Get(t *testing.T) {
	requestBuilder := &TableItemRequestBuilder2{RequestBuilder: &mockRequestBuilder{}}

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

func TestTableItemRequestBuilder2_Delete(t *testing.T) {
	requestBuilder := &TableItemRequestBuilder2{RequestBuilder: &mockRequestBuilder{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				mockRB := &mockRequestBuilder{}
				mockRB.On("SendDelete2", mock.AnythingOfType("*core.RequestConfiguration")).Return(nil)

				requestBuilder.RequestBuilder = mockRB
			},
			ExpectedErr: nil,
		},
		{
			Title: "Error",
			Setup: func() {
				mockRB := &mockRequestBuilder{}
				mockRB.On("SendDelete2", mock.AnythingOfType("*core.RequestConfiguration")).Return(errors.New("unable to send request"))

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

			err := requestBuilder.Delete(nil)
			assert.Equal(t, test.ExpectedErr, err)
			requestBuilder.RequestBuilder.(*mockRequestBuilder).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

func TestTableItemRequestBuilder2_Put(t *testing.T) {
	requestBuilder := &TableItemRequestBuilder2{RequestBuilder: &mockRequestBuilder{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				mockRB := &mockRequestBuilder{}
				mockRB.On("SendPut2", mock.AnythingOfType("*core.RequestConfiguration")).Return(nil)

				requestBuilder.RequestBuilder = mockRB
			},
			Input:       NewTableEntry(),
			ExpectedErr: nil,
		},
		{
			Title: "Error",
			Setup: func() {
				mockRB := &mockRequestBuilder{}
				mockRB.On("SendPut2", mock.AnythingOfType("*core.RequestConfiguration")).Return(errors.New("unable to send request"))

				requestBuilder.RequestBuilder = mockRB
			},
			Input:       NewTableEntry(),
			ExpectedErr: errors.New("unable to send request"),
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

			_, err := requestBuilder.Put(test.Input, nil)
			assert.Equal(t, test.ExpectedErr, err)
			requestBuilder.RequestBuilder.(*mockRequestBuilder).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}
