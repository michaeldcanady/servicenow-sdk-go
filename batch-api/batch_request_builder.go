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

// NewBatchRequestBuilderInternal instantiates a new BatchRequestBuilder with custom parsable for table entries.
func NewBatchRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *BatchRequestBuilder {
	m := &BatchRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, batchURLTemplate, pathParameters),
	}
	return m
}

// NewBatchRequestBuilder instantiates a new BatchRequestBuilder with custom parsable for table entries.
func NewBatchRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *BatchRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
	return NewBatchRequestBuilderInternal(urlParams, requestAdapter)
}

// Post produces a batch response using the specified parameters
func (rB *BatchRequestBuilder) Post(ctx context.Context, body BatchRequest, requestConfiguration *BatchRequestBuilderPostRequestConfiguration) (*BatchResponseModel, error) {
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

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateBatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	typedResp, ok := resp.(*BatchResponseModel)
	if !ok {
		return nil, errors.New("resp is not *BatchResponse")
	}

	return typedResp, nil
}

// toPostRequestInformation converts provided parameters into request information
func (rB *BatchRequestBuilder) toPostRequestInformation(ctx context.Context, body BatchRequest, requestConfiguration *BatchRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	// BUG: method should be POST not PUT
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(newInternal.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), newInternal.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}

	return kiotaRequestInfo.RequestInformation, nil
}
