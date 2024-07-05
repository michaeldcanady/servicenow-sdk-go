package tableapi

import (
	"context"
	"errors"
	"testing"

	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	"github.com/michaeldcanady/servicenow-sdk-go/table-api/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewTableItemRequestBuilder2(t *testing.T) {
	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Input: []interface{}{&intCore.MockClient2{}, map[string]string{"baseurl": "baseurl", "table": "table", "sysId": "fdafsdfdsa"}},
		},
		{
			Title:       "missing sysId",
			Input:       []interface{}{&intCore.MockClient2{}, map[string]string{"baseurl": "baseurl", "table": "table"}},
			ExpectedErr: errors.New("missing \"sysId\" parameter"),
		},
		{
			Title:       "missing table",
			Input:       []interface{}{&intCore.MockClient2{}, map[string]string{"baseurl": "baseurl"}},
			ExpectedErr: errors.New("missing \"table\" parameter"),
		},
		{
			Title:       "missing baseurl",
			Input:       []interface{}{&intCore.MockClient2{}, map[string]string{}},
			ExpectedErr: errors.New("pathParameters must contain a value for \"baseurl\" for the URL to be built"),
		},
		{
			Title:       "missing client",
			Input:       []interface{}{(*intCore.MockClient2)(nil), map[string]string{}},
			ExpectedErr: errors.New("client can't be nil"),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			inputs := test.Input.([]interface{})

			_, err := NewTableItemRequestBuilder2(inputs[0].(intCore.Client2), inputs[1].(map[string]string))
			assert.Equal(t, test.ExpectedErr, err)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

func TestTableItemRequestBuilder2_Get(t *testing.T) {
	requestBuilder := &TableItemRequestBuilder2{RequestBuilder2: &intCore.MockRequestBuilder2{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.GET, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return(&TableItemResponse2[TableEntry]{}, nil).Once()
			},
			ExpectedErr: nil,
		},
		{
			Title: "Error",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.GET, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return((*TableItemResponse2[TableEntry])(nil), errors.New("unable to send request")).Once()
			},
			ExpectedErr: errors.New("unable to send request"),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			_, err := requestBuilder.Get(context.Background(), nil)
			assert.Equal(t, test.ExpectedErr, err)
			requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

func TestTableItemRequestBuilder2_Delete(t *testing.T) {
	requestBuilder := &TableItemRequestBuilder2{RequestBuilder2: &intCore.MockRequestBuilder2{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.DELETE, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return(nil, nil).Once()
			},
			ExpectedErr: nil,
		},
		{
			Title: "Error",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.DELETE, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return(nil, errors.New("unable to send request")).Once()
			},
			ExpectedErr: errors.New("unable to send request"),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			err := requestBuilder.Delete(context.Background(), nil)
			assert.Equal(t, test.ExpectedErr, err)
			requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

func TestTableItemRequestBuilder2_Put(t *testing.T) {
	requestBuilder := &TableItemRequestBuilder2{RequestBuilder2: &intCore.MockRequestBuilder2{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.PUT, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return(&TableItemResponse{}, nil).Once()
			},
			Input:       []intCore.RequestConfigurationOption{intCore.WithData(NewTableEntry())},
			ExpectedErr: nil,
		},
		{
			Title: "Error",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.PUT, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return((*TableItemResponse)(nil), errors.New("unable to send request")).Once()
			},
			Input:       []intCore.RequestConfigurationOption{intCore.WithData(NewTableEntry())},
			ExpectedErr: errors.New("unable to send request"),
		},
		{
			Title: "Nil tableEntry",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.PUT, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return((*TableItemResponse)(nil), errors.New("entry is nil")).Once()
			},
			Input:       []intCore.RequestConfigurationOption{},
			ExpectedErr: errors.New("entry is nil"),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			_, err := requestBuilder.Put(context.Background(), test.Input.([]intCore.RequestConfigurationOption)...)
			assert.Equal(t, test.ExpectedErr, err)
			requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}
