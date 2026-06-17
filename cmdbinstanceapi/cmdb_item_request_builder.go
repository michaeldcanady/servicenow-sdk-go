package cmdbinstanceapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cmdbItemURLTemplate = "{+baseurl}/api/now/v1/cmdb/instance/{className}/{sys_id}"
)

// CmdbItemRequestBuilder provides operations to manage a specific CI record.
type CmdbItemRequestBuilder struct {
	internal.RequestBuilder
}

// NewCmdbItemRequestBuilderInternal instantiates a new CmdbItemRequestBuilder.
func NewCmdbItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CmdbItemRequestBuilder {
	return &CmdbItemRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, cmdbItemURLTemplate, pathParameters),
	}
}

// Get queries attributes and relationship information for a specific record.
func (rB *CmdbItemRequestBuilder) Get(ctx context.Context, config *CmdbItemRequestBuilderGetRequestConfiguration) (*internal.BaseServiceNowItemResponse[CmdbInstance], error) {
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"400": internal.CreateBadRequestErrorFromDiscriminatorValue,
		"401": internal.CreateUnauthorizedErrorFromDiscriminatorValue,
		"403": internal.CreateForbiddenErrorFromDiscriminatorValue,
		"404": internal.CreateNotFoundErrorFromDiscriminatorValue,
		"429": internal.CreateTooManyRequestsErrorFromDiscriminatorValue,
		"5XX": internal.CreateServerErrorFromDiscriminatorValue,
		"XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowItemResponseFromDiscriminatorValue[CmdbInstance](CreateCmdbInstanceFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(*internal.BaseServiceNowItemResponse[CmdbInstance]), nil
}

// Put replaces a CI record.
func (rB *CmdbItemRequestBuilder) Put(ctx context.Context, body CmdbInstance, config *CmdbItemRequestBuilderPutRequestConfiguration) (*internal.BaseServiceNowItemResponse[CmdbInstance], error) {
	requestInfo, err := rB.ToPutRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"400": internal.CreateBadRequestErrorFromDiscriminatorValue,
		"401": internal.CreateUnauthorizedErrorFromDiscriminatorValue,
		"403": internal.CreateForbiddenErrorFromDiscriminatorValue,
		"404": internal.CreateNotFoundErrorFromDiscriminatorValue,
		"429": internal.CreateTooManyRequestsErrorFromDiscriminatorValue,
		"5XX": internal.CreateServerErrorFromDiscriminatorValue,
		"XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowItemResponseFromDiscriminatorValue[CmdbInstance](CreateCmdbInstanceFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(*internal.BaseServiceNowItemResponse[CmdbInstance]), nil
}

// Patch updates a CI record.
func (rB *CmdbItemRequestBuilder) Patch(ctx context.Context, body CmdbInstance, config *CmdbItemRequestBuilderPatchRequestConfiguration) (*internal.BaseServiceNowItemResponse[CmdbInstance], error) {
	requestInfo, err := rB.ToPatchRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"400": internal.CreateBadRequestErrorFromDiscriminatorValue,
		"401": internal.CreateUnauthorizedErrorFromDiscriminatorValue,
		"403": internal.CreateForbiddenErrorFromDiscriminatorValue,
		"404": internal.CreateNotFoundErrorFromDiscriminatorValue,
		"429": internal.CreateTooManyRequestsErrorFromDiscriminatorValue,
		"5XX": internal.CreateServerErrorFromDiscriminatorValue,
		"XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowItemResponseFromDiscriminatorValue[CmdbInstance](CreateCmdbInstanceFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(*internal.BaseServiceNowItemResponse[CmdbInstance]), nil
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *CmdbItemRequestBuilder) ToGetRequestInformation(ctx context.Context, config *CmdbItemRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if headers := config.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return requestInfo, nil
}

// ToPutRequestInformation converts request configurations to Put request information.
func (rB *CmdbItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body CmdbInstance, config *CmdbItemRequestBuilderPutRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if headers := config.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	if !conversion.IsNil(body) {
		err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, body)
		if err != nil {
			return nil, err
		}
	}

	return requestInfo, nil
}

// ToPatchRequestInformation converts request configurations to Patch request information.
func (rB *CmdbItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body CmdbInstance, config *CmdbItemRequestBuilderPatchRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PATCH, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if headers := config.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	if !conversion.IsNil(body) {
		err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, body)
		if err != nil {
			return nil, err
		}
	}

	return requestInfo, nil
}

// Relation provides operations to manage CI relationships.
func (rB *CmdbItemRequestBuilder) Relation() *CmdbRelationRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	return NewCmdbRelationRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
