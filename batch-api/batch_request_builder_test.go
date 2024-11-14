package batchapi

import (
	"errors"
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/batch-api/internal"
	"github.com/RecoLabs/servicenow-sdk-go/core"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewBatchRequestBuilder(t *testing.T) {
	mockClient := new(internal.MockClient)
	pathParameters := map[string]string{"param1": "value1"}

	result := NewBatchRequestBuilder(mockClient, pathParameters)

	assert.IsType(t, &core.RequestBuilder{}, result.RequestBuilder) //nolint:staticcheck
	assert.NotNil(t, result)
	assert.IsType(t, &BatchRequestBuilder{}, result)

	builder := result.RequestBuilder.(*core.RequestBuilder) //nolint:staticcheck

	assert.Equal(t, "{+baseurl}/v1/batch", builder.UrlTemplate)
	assert.Equal(t, pathParameters, builder.PathParameters)
	assert.Equal(t, mockClient, builder.Client)
}

func TestBatchRequestBuilder_Post(t *testing.T) {
	// Create a new mock request builder
	mockBuilder := new(internal.MockRequestBuilder)

	// Create a new batch request builder with the mock builder
	rB := &BatchRequestBuilder{
		RequestBuilder: mockBuilder,
	}

	request := &internal.MockBatchRequest{}

	// Define the test cases
	tests := []internal.Test[any]{
		{
			Title: "Successful Post",
			Setup: func() {
				mockBuilder.On("SendPost3", mock.AnythingOfType("*core.RequestConfigurationImpl")).Return(nil)
			},
		},
		{
			Title: "Failed Post",
			Setup: func() {
				// Set up the mock to return an error for a failed post
				mockBuilder.Mock.ExpectedCalls[0].Return(errors.New("post error"))
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
			_, err := rB.Post(request)

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
