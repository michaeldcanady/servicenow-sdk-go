package batchapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// batchURLTemplate the url template for Service-Now batch API
	batchURLTemplate = "{+baseurl}/api/now/v1/batch"
)

// BatchRequestBuilder constructs batch requests for the specified base URL.
type BatchRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewBatchRequestBuilder2Internal instantiates a new BatchRequestBuilder with custom parsable for table entries.
func NewBatchRequestBuilder2Internal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *BatchRequestBuilder {
	m := &BatchRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, batchURLTemplate, pathParameters),
	}
	return m
}

// NewBatchRequestBuilder2 instantiates a new BatchRequestBuilder with custom parsable for table entries.
func NewBatchRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *BatchRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
	return NewBatchRequestBuilder2Internal(urlParams, requestAdapter)
}

// Post produces a batch response using the specified parameters
func (rB *BatchRequestBuilder) Post(ctx context.Context, body BatchRequestable, requestConfiguration *BatchRequestBuilderPostRequestConfiguration) (BatchResponseable, error) {
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
	errorMapping := abstractions.ErrorMappings{
		//"401":
		//"500":
	}

	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateBatchResponseFromDiscriminatorValue, errorMapping)
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
func (rB *BatchRequestBuilder) toPostRequestInformation(ctx context.Context, body BatchRequestable, requestConfiguration *BatchRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	kiotaRequestInfo.Headers.TryAdd(newInternal.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), newInternal.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}

	return kiotaRequestInfo.RequestInformation, nil
}
