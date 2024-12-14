package batchapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	batchURLTemplate = "{+baseurl}/api/now/v1/batch"
)

// BatchRequestBuilder constructs batch requests for the specified base URL.
type BatchRequestBuilder2 struct {
	abstractions.BaseRequestBuilder
}

// NewAPIV1CompatibleBatchRequestBuilder2Internal ...
func NewAPIV1CompatibleBatchRequestBuilder2Internal(
	pathParameters map[string]string,
	client core.Client,
) *BatchRequestBuilder2 {
	reqAdapter, _ := internal.NewServiceNowRequestAdapterBase(core.NewAPIV1ClientAdapter(client))

	return NewBatchRequestBuilder2Internal(
		pathParameters,
		reqAdapter,
	)
}

// NewBatchRequestBuilder2Internal instantiates a new BatchRequestBuilder2 with custom parsable for table entries.
func NewBatchRequestBuilder2Internal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *BatchRequestBuilder2 {
	m := &BatchRequestBuilder2{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, batchURLTemplate, pathParameters),
	}
	return m
}

// NewBatchRequestBuilder2 instantiates a new BatchRequestBuilder2 with custom parsable for table entries.
func NewBatchRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *BatchRequestBuilder2 {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawURL
	return NewBatchRequestBuilder2Internal(urlParams, requestAdapter)
}

// Post
func (rB *BatchRequestBuilder2) Post(ctx context.Context, body BatchRequestable, requestConfiguration *BatchRequestBuilder2PostRequestConfiguration) (BatchResponseable, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	if internal.IsNil(body) {
		return nil, errors.New("body can't be nil")
	}

	requestInfo, err := rB.toPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}

	// TODO: implement error mapping
	errorMapping := abstractions.ErrorMappings{}

	resp, err := rB.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateBatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	typedResp, ok := resp.(BatchResponseable)
	if !ok {
		return nil, errors.New("resp is not BatchResponseable")
	}

	return typedResp, nil
}

func (rB *BatchRequestBuilder2) toPostRequestInformation(_ context.Context, _ BatchRequestable, _ *BatchRequestBuilder2PostRequestConfiguration) (*abstractions.RequestInformation, error) {
	return nil, errors.New("toPostRequestInformation is not implemented")
}
