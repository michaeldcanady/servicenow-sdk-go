package tableapi

import (
	"context"
	"errors"
	"testing"

	intCore "github.com/RecoLabs/servicenow-sdk-go/internal/core"
	"github.com/RecoLabs/servicenow-sdk-go/table-api/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var _ intCore.RequestBuilder2 = (*mockRequestBuilder2)(nil)

type mockRequestBuilder2 struct {
	mock.Mock
}

func (rB *mockRequestBuilder2) Send(ctx context.Context, method intCore.HttpMethod, config intCore.RequestConfiguration) (interface{}, error) {
	args := rB.Called(ctx, method, config)
	return args.Get(0), args.Error(1)
}

func (rB *mockRequestBuilder2) GetPathParameters() map[string]string {
	args := rB.Called()
	return args.Get(0).(map[string]string)
}

func (rB *mockRequestBuilder2) GetClient() intCore.ClientSendable {
	args := rB.Called()
	return args.Get(0).(intCore.ClientSendable)
}

func (rB *mockRequestBuilder2) GetURLTemplate() string {
	args := rB.Called()
	return args.String(0)
}

func TestNewTableRequestBuilder2(t *testing.T) {
	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Input: []interface{}{
				&mockClient2{},
				map[string]string{"baseurl": "baseurl", "table": "table"},
			},
		},
		{
			Title: "missing table",
			Input: []interface{}{
				&mockClient2{},
				map[string]string{"baseurl": "baseurl"},
			},
			ExpectedErr: errors.New("missing \"table\" parameter"),
		},
		{
			Title:       "missing baseurl",
			Input:       []interface{}{&mockClient2{}, map[string]string{}},
			ExpectedErr: errors.New("pathParameters must contain a value for \"baseurl\" for the URL to be built"),
		},
		{
			Title:       "missing client",
			Input:       []interface{}{(*mockClient2)(nil), map[string]string{}},
			ExpectedErr: errors.New("client can't be nil"),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			inputs := test.Input.([]interface{})

			_, err := newTableRequestBuilder2[*TableRecordImpl](inputs[0].(intCore.ClientSendable), inputs[1].(map[string]string))
			assert.Equal(t, test.ExpectedErr, err)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

//nolint:dupl
func TestTableRequestBuilder2_Get(t *testing.T) {
	requestBuilder := &tableRequestBuilder2[*TableRecordImpl]{&mockRequestBuilder2{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				mockRB := &mockRequestBuilder2{}
				mockRB.On(
					"Send",
					context.Background(),
					intCore.MethodGet,
					mock.AnythingOfType("*core.RequestConfigurationImpl"),
				).Return(
					&tableCollectionResponse3[*TableRecordImpl]{},
					nil,
				)

				requestBuilder.RequestBuilder2 = mockRB
			},
			ExpectedErr: nil,
		},
		{
			Title: "Error",
			Setup: func() {
				mockRB := &mockRequestBuilder2{}
				mockRB.On(
					"Send",
					context.Background(),
					intCore.MethodGet,
					mock.AnythingOfType("*core.RequestConfigurationImpl"),
				).Return(
					&tableCollectionResponse3[*TableRecordImpl]{},
					errors.New("unable to send request"),
				)

				requestBuilder.RequestBuilder2 = mockRB
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
			requestBuilder.RequestBuilder2.(*mockRequestBuilder2).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

func TestTableRequestBuilder2_Post(t *testing.T) {
	requestBuilder := &tableRequestBuilder2[*TableRecordImpl]{&mockRequestBuilder2{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				mockRB := &mockRequestBuilder2{}
				mockRB.On(
					"Send",
					context.Background(),
					intCore.MethodPost,
					mock.AnythingOfType("*core.RequestConfigurationImpl"),
				).Return(
					&tableItemResponse3[*TableRecordImpl]{},
					nil,
				)

				requestBuilder.RequestBuilder2 = mockRB
			},
			ExpectedErr: nil,
			Input:       NewTableEntry(),
		},
		{
			Title: "Error",
			Setup: func() {
				mockRB := &mockRequestBuilder2{}
				mockRB.On(
					"Send",
					context.Background(),
					intCore.MethodPost,
					mock.AnythingOfType("*core.RequestConfigurationImpl"),
				).Return(
					&tableItemResponse3[*TableRecordImpl]{},
					errors.New("unable to send request"),
				)

				requestBuilder.RequestBuilder2 = mockRB
			},
			ExpectedErr: errors.New("unable to send request"),
			Input:       NewTableEntry(),
		},
		{
			Title: "Nil tableEntry",
			Setup: func() {
				mockRB := &mockRequestBuilder2{}
				requestBuilder.RequestBuilder2 = mockRB
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

			_, err := requestBuilder.Post(context.Background(), test.Input, nil)
			assert.Equal(t, test.ExpectedErr, err)
			requestBuilder.RequestBuilder2.(*mockRequestBuilder2).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}
