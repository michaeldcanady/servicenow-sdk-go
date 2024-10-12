package batchapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

const (
	batchURLTemplate = "{+baseurl}/v1/batch" // Must be versioned
)

// BatchRequestBuilder constructs batch requests for the specified base URL.
type BatchRequestBuilder2 struct {
	intCore.Sendable
}

// NewBatchRequestBuilder creates a new BatchRequestBuilder.
// It takes a client (core.Client) and pathParameters (map[string]string).
func NewBatchRequestBuilder2(client intCore.ClientSendable, pathParameters map[string]string) *BatchRequestBuilder2 {
	requestBuilder := intCore.NewRequestBuilder2(
		client,
		batchURLTemplate,
		pathParameters,
	)
	return &BatchRequestBuilder2{
		requestBuilder,
	}
}

// Post sends a batch request and returns the BatchResponse.
func (rB *BatchRequestBuilder2) Post(ctx context.Context, request BatchRequest) (BatchResponse, error) {
	config := &core.RequestConfiguration{
		Data:     request,
		Response: &batchResponse{},
	}

	resp, err := rB.Send(ctx, intCore.MethodPost, config)
	if err != nil {
		return nil, err
	}

	return resp.(*batchResponse), nil
}
