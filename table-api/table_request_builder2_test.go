package tableapi

import (
	"context"
	"errors"
	"maps"
	"testing"

	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	"github.com/michaeldcanady/servicenow-sdk-go/table-api/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewTableRequestBuilder2(t *testing.T) {
	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Input: []interface{}{&intCore.MockClient2{}, map[string]string{"baseurl": "baseurl", "table": "table"}},
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

			_, err := NewTableRequestBuilder2(inputs[0].(intCore.Client2), inputs[1].(map[string]string))
			assert.Equal(t, test.ExpectedErr, err)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

func TestTableRequestBuilder2_ByID(t *testing.T) {
	requestBuilder := &TableRequestBuilder2{RequestBuilder2: &intCore.MockRequestBuilder2{}}
	mockClient := &intCore.MockClient2{}
	mockPathParams := map[string]string{"baseurl": "baseurl", "table": "test_table"}
	retValue := (*TableItemRequestBuilder2)(nil)

	tests := []internal.Test[*TableItemRequestBuilder2]{
		{
			Title: "Successful",
			Input: "sys_id",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("GetClient").Return(mockClient).Once()
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("GetPathParameters").Return(mockPathParams).Once()
				expectedParams := maps.Clone(mockPathParams)
				expectedParams["sysId"] = "sys_id"
				retValue, _ = NewTableItemRequestBuilder2(mockClient, expectedParams)
			},
			Expected: retValue,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
				test.Expected = retValue
			}

			itemReqBuilder, err := requestBuilder.ByID(test.Input.(string))
			assert.Equal(t, *test.Expected, *itemReqBuilder)
			assert.Equal(t, test.ExpectedErr, err)
			requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

//nolint:dupl
func TestTableRequestBuilder2_Get(t *testing.T) {
	requestBuilder := &TableRequestBuilder2{RequestBuilder2: &intCore.MockRequestBuilder2{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.GET, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return(&TableCollectionResponse2[TableEntry]{}, nil).Once()
			},
			ExpectedErr: nil,
		},
		{
			Title: "Error",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.GET, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return((*TableCollectionResponse2[TableEntry])(nil), errors.New("unable to send request")).Once()
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

func TestTableRequestBuilder2_Post(t *testing.T) {
	requestBuilder := &TableRequestBuilder2{RequestBuilder2: &intCore.MockRequestBuilder2{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.POST, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return(&TableItemResponse2[TableEntry]{}, nil).Once()
			},
			ExpectedErr: nil,
			Input:       []intCore.RequestConfigurationOption{intCore.WithData(NewTableEntry())},
		},
		{
			Title: "Error",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.POST, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return((*TableItemResponse2[TableEntry])(nil), errors.New("unable to send request")).Once()
			},
			ExpectedErr: errors.New("unable to send request"),
			Input:       []intCore.RequestConfigurationOption{intCore.WithData(NewTableEntry())},
		},
		{
			Title: "Nil tableEntry",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.POST, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return((*TableItemResponse2[TableEntry])(nil), errors.New("entry is nil")).Once()
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

			_, err := requestBuilder.Post(context.Background(), test.Input.([]intCore.RequestConfigurationOption)...)
			assert.Equal(t, test.ExpectedErr, err)
			requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

func TestTableRequestBuilder2_Head(t *testing.T) {
	requestBuilder := &TableRequestBuilder2{RequestBuilder2: &intCore.MockRequestBuilder2{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.HEAD, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return(&TableCollectionResponse{}, nil).Once()
			},
			ExpectedErr: nil,
			Input:       []intCore.RequestConfigurationOption{},
		},
		{
			Title: "Error",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.HEAD, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return((*TableCollectionResponse)(nil), errors.New("unable to send request")).Once()
			},
			ExpectedErr: errors.New("unable to send request"),
			Input:       []intCore.RequestConfigurationOption{},
		},
		{
			Title: "Nil tableEntry",
			Setup: func() {
				requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).On("Send", context.Background(), intCore.HEAD, mock.AnythingOfType("[]core.RequestConfigurationOption")).Return((*TableCollectionResponse)(nil), errors.New("entry is nil")).Once()
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

			_, err := requestBuilder.Head(context.Background(), test.Input.([]intCore.RequestConfigurationOption)...)
			assert.Equal(t, test.ExpectedErr, err)
			requestBuilder.RequestBuilder2.(*intCore.MockRequestBuilder2).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}
