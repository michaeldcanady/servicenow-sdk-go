package batchapi

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/batch-api/internal"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockClient2 struct {
	mock.Mock
}

func (c *mockClient2) SendWithContext(ctx context.Context, info intCore.RequestInformation, mapping intCore.ErrorMapping) (*http.Response, error) {
	args := c.Called(ctx, info, mapping)
	return args.Get(0).(*http.Response), args.Error(1)
}

type mockRequestBuilder2 struct {
	mock.Mock
}

func (rB *mockRequestBuilder2) Send(ctx context.Context, method intCore.HttpMethod, config intCore.RequestConfiguration) (interface{}, error) {
	args := rB.Called(ctx, method, config)
	return args.Get(0), args.Error(1)
}

func TestNewBatchRequestBuilder2(t *testing.T) {
	mockClient := new(mockClient2)
	pathParameters := map[string]string{"param1": "value1"}

	result, _ := NewBatchRequestBuilder2(mockClient, pathParameters)

	assert.Implements(t, (*intCore.Sendable)(nil), result.(*batchRequestBuilder2).Sendable)
	assert.NotNil(t, result)
	assert.IsType(t, &batchRequestBuilder2{}, result)

	builder := result.(*batchRequestBuilder2).Sendable.(intCore.RequestBuilder2)

	assert.Equal(t, "{+baseurl}/v1/batch", builder.GetURLTemplate())
	assert.Equal(t, pathParameters, builder.GetPathParameters())
	assert.Equal(t, mockClient, builder.GetClient())
}

func TestBatchRequestBuilder2_Post(t *testing.T) {
	// Create a new mock request builder
	mockBuilder := new(mockRequestBuilder2)

	// Create a new batch request builder with the mock builder
	rB := &batchRequestBuilder2{
		Sendable: mockBuilder,
	}

	request := &internal.MockBatchRequest{}

	// Define the test cases
	tests := []internal.Test[any]{
		{
			Title: "Successful Post",
			Setup: func() {
				mockBuilder.On("Send", context.Background(), intCore.MethodPost, mock.AnythingOfType("*core.RequestConfigurationImpl")).Return((*batchResponse)(nil), nil)
			},
		},
		{
			Title: "Failed Post",
			Setup: func() {
				// Set up the mock to return an error for a failed post
				mockBuilder.Mock.ExpectedCalls[0].Return((*batchResponse)(nil), errors.New("post error"))
			},
			ExpectedErr: errors.New("post error"),
		},
	}

	// Run the test cases
	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			// Set up the test case
			if test.Setup != nil {
				test.Setup()
			}

			// Call the function under test
			_, err := rB.Post(context.Background(), request)

			// Check the result
			assert.Equal(t, test.ExpectedErr, err)

			// Assert that the mock expectations were met
			mockBuilder.AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}
