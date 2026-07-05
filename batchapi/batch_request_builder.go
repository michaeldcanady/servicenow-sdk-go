package batchapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

const (
	// batchURLTemplate the url template for Service-Now batch API
	batchURLTemplate = "{+baseurl}/api/now/v1/batch"
)

// BatchRequestBuilder constructs batch requests for the specified base URL.
type BatchRequestBuilder struct {
	core.RequestBuilder
}

// NewBatchRequestBuilderInternal instantiates a new BatchRequestBuilder with custom parsable for table entries.
func NewBatchRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *BatchRequestBuilder {
	m := &BatchRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, batchURLTemplate, pathParameters),
	}
	return m
}

// NewBatchRequestBuilder instantiates a new BatchRequestBuilder with custom parsable for table entries.
func NewBatchRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *BatchRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[internal.RawURLKey] = rawURL
	return NewBatchRequestBuilderInternal(urlParams, requestAdapter)
}

// Head sends an HTTP HEAD request and returns the response headers.
func (rB *BatchRequestBuilder) Head(ctx context.Context, requestConfiguration *BatchRequestBuilderGetRequestConfiguration) (*abstractions.ResponseHeaders, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if conversion.IsNil(requestConfiguration) {
		requestConfiguration = &BatchRequestBuilderGetRequestConfiguration{}
	}

	headerOpt := nethttplibrary.NewHeadersInspectionOptions()
	headerOpt.InspectResponseHeaders = true
	requestConfiguration.Options = append(requestConfiguration.Options, headerOpt)

	requestInfo, err := rB.ToHeadRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()

	if err = rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping); err != nil {
		return nil, err
	}

	return headerOpt.GetResponseHeaders(), nil
}

// ToHeadRequestInformation converts provided parameters into request information
func (rB *BatchRequestBuilder) ToHeadRequestInformation(_ context.Context, requestConfiguration *BatchRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.HEAD, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return kiotaRequestInfo.RequestInformation, nil
}

// Post produces a batch response using the specified parameters
func (rB *BatchRequestBuilder) Post(ctx context.Context, body BatchRequest, requestConfiguration *BatchRequestBuilderPostRequestConfiguration) (*BatchResponseModel, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if conversion.IsNil(body) {
		return nil, snerrors.NewValidationError("body")
	}

	requestInfo, err := rB.toPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()

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
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}

	return kiotaRequestInfo.RequestInformation, nil
}
