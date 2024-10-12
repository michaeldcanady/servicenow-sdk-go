package batchapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

const (
	batchURLTemplate = "{+baseurl}/v1/batch" // Must be versioned
)

// BatchRequestBuilder2 defines the operations available for the batch api as a whole.
type BatchRequestBuilder2 interface {
	Post(ctx context.Context, request BatchRequest) (BatchResponse, error)
}

// batchRequestBuilder2 constructs batch requests for the specified base URL.
type batchRequestBuilder2 struct {
	intCore.Sendable
}

// NewBatchRequestBuilder creates a new BatchRequestBuilder.
// It takes a client (core.Client) and pathParameters (map[string]string).
func NewBatchRequestBuilder2(client intCore.ClientSendable, pathParameters map[string]string) (BatchRequestBuilder2, error) {
	requestBuilder := intCore.NewRequestBuilder2(
		client,
		batchURLTemplate,
		pathParameters,
	)
	return &batchRequestBuilder2{
		requestBuilder,
	}, nil
}

// Post sends a batch request and returns the BatchResponse.
func (rB *batchRequestBuilder2) Post(ctx context.Context, request BatchRequest) (BatchResponse, error) {
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
