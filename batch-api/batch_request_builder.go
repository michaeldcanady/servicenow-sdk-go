package batchapi

import (
	"github.com/RecoLabs/servicenow-sdk-go/batch-api/internal"
	"github.com/RecoLabs/servicenow-sdk-go/core"
)

// BatchRequestBuilder constructs batch requests for the specified base URL.
type BatchRequestBuilder struct {
	internal.RequestBuilder
}

// NewBatchRequestBuilder creates a new BatchRequestBuilder.
// It takes a client (core.Client) and pathParameters (map[string]string).
func NewBatchRequestBuilder(client core.Client, pathParameters map[string]string) *BatchRequestBuilder {
	requestBuilder := core.NewRequestBuilder( //nolint:staticcheck
		client,
		// Must be versioned
		"{+baseurl}/v1/batch",
		pathParameters,
	)
	return &BatchRequestBuilder{
		RequestBuilder: requestBuilder,
	}
}

// Post sends a batch request and returns the BatchResponse.
func (rB *BatchRequestBuilder) Post(request BatchRequest) (BatchResponse, error) {
	config := &core.RequestConfiguration{
		Data:     request,
		Response: &batchResponse{},
	}

	err := rB.SendPost3(config)
	if err != nil {
		return nil, err
	}

	return config.Response.(*batchResponse), nil
}
