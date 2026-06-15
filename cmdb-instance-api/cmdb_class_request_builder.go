package cmdbinstanceapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cmdbClassURLTemplate = "{+baseurl}/api/now/v1/cmdb/instance/{className}{?sysparm_query,sysparm_limit,sysparm_offset}"
)

// CmdbClassRequestBuilder provides operations to manage a specific CMDB class.
type CmdbClassRequestBuilder struct {
	internal.RequestBuilder
}

// NewCmdbClassRequestBuilderInternal instantiates a new CmdbClassRequestBuilder.
func NewCmdbClassRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CmdbClassRequestBuilder {
	return &CmdbClassRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, cmdbClassURLTemplate, pathParameters),
	}
}

// Get queries records for a CMDB class.
func (rB *CmdbClassRequestBuilder) Get(ctx context.Context, config *CmdbClassRequestBuilderGetRequestConfiguration) (*internal.BaseServiceNowCollectionResponse[CmdbInstance], error) {
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowCollectionResponseFromDiscriminatorValue[CmdbInstance](CreateCmdbInstanceFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(*internal.BaseServiceNowCollectionResponse[CmdbInstance]), nil
}

// Post creates a record with associated relations.
func (rB *CmdbClassRequestBuilder) Post(ctx context.Context, body CmdbInstance, config *CmdbClassRequestBuilderPostRequestConfiguration) (*internal.BaseServiceNowItemResponse[CmdbInstance], error) {
	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
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
func (rB *CmdbClassRequestBuilder) ToGetRequestInformation(ctx context.Context, config *CmdbClassRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if headers := config.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParameters := config.QueryParameters; !conversion.IsNil(queryParameters) {
			kiotaRequestInfo.AddQueryParameters(queryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return requestInfo, nil
}

// ToPostRequestInformation converts request configurations to Post request information.
func (rB *CmdbClassRequestBuilder) ToPostRequestInformation(ctx context.Context, body CmdbInstance, config *CmdbClassRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
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

// ByID provides operations to manage a specific CI record.
func (rB *CmdbClassRequestBuilder) ByID(sysID string) *CmdbItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["sys_id"] = sysID
	return NewCmdbItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
