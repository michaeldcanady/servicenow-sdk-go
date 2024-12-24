package batchapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	intHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// batchURLTemplate the url template for Service-Now batch API
	batchURLTemplate = "{+baseurl}/api/now/v1/batch"
)

// BatchRequestBuilder constructs batch requests for the specified base URL.
type BatchRequestBuilder2 struct {
	abstractions.BaseRequestBuilder
}

// NewAPIV1CompatibleBatchRequestBuilder2Internal converts api v1 compatible elements into api v2 compatible elements
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
	urlParams[intCore.RawURLKey] = rawURL
	return NewBatchRequestBuilder2Internal(urlParams, requestAdapter)
}

// Post produces a batch response using the specified parameters
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

// toPostRequestInformation converts provided parameters into request information
func (rB *BatchRequestBuilder2) toPostRequestInformation(ctx context.Context, body BatchRequestable, requestConfiguration *BatchRequestBuilder2PostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.UrlTemplate, rB.PathParameters)
	kiotaRequestInfo := &intHttp.KiotaRequestInformation{RequestInformation: *requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	kiotaRequestInfo.Headers.TryAdd("Accept", "application/json")

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}

	return &kiotaRequestInfo.RequestInformation, nil
}
