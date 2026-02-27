package batchapi

import (
	"context"
	"errors"

	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// batchURLTemplate the url template for Service-Now batch API
	batchURLTemplate = "{+baseurl}/api/now/v1/batch"
)

// BatchRequestBuilder constructs batch requests for the specified base URL.
type BatchRequestBuilder struct {
	kiota.RequestBuilder
}

// NewBatchRequestBuilderInternal instantiates a new BatchRequestBuilder with custom parsable for table entries.
func NewBatchRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *BatchRequestBuilder {
	m := &BatchRequestBuilder{
		kiota.NewBaseRequestBuilder(requestAdapter, batchURLTemplate, pathParameters),
	}
	return m
}

// NewBatchRequestBuilder instantiates a new BatchRequestBuilder with custom parsable for table entries.
func NewBatchRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *BatchRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[utils.RawURLKey] = rawURL
	return NewBatchRequestBuilderInternal(urlParams, requestAdapter)
}

// Post produces a batch response using the specified parameters
func (rB *BatchRequestBuilder) Post(ctx context.Context, body BatchRequest, requestConfiguration *BatchRequestBuilderPostRequestConfiguration) (*BatchResponseModel, error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if utils.IsNil(body) {
		return nil, errors.New("body can't be nil")
	}

	requestInfo, err := rB.toPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": model.CreateServiceNowErrorFromDiscriminatorValue,
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
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &kiota.KiotaRequestInformation{RequestInformation: requestInfo}
	if !utils.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !utils.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !utils.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), utils.ContentTypeApplicationJSON)

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), utils.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}

	return kiotaRequestInfo.RequestInformation, nil
}
