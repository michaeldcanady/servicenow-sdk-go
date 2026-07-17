package cmdbinstanceapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cmdbClassURLTemplate = "{+baseurl}/api/now/v1/cmdb/instance/{className}{?sysparm_query,sysparm_limit,sysparm_offset}"
)

// CmdbClassRequestBuilder provides operations to manage a specific CMDB class.
type CmdbClassRequestBuilder struct {
	core.RequestBuilder
}

var _ core.ItemPostRequestBuilder[CmdbInstance, abstractions.DefaultQueryParameters, abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]] = (*CmdbClassRequestBuilder)(nil)

// NewCmdbClassRequestBuilderInternal instantiates a new CmdbClassRequestBuilder.
func NewCmdbClassRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CmdbClassRequestBuilder {
	return &CmdbClassRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, cmdbClassURLTemplate, pathParameters),
	}
}

// Get queries records for a CMDB class.
func (rB *CmdbClassRequestBuilder) Get(ctx context.Context, config *CmdbClassRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowCollectionResponse[CmdbInstance], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowCollectionResponseFromDiscriminatorValue[CmdbInstance](CreateCmdbInstanceFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(*core.BaseServiceNowCollectionResponse[CmdbInstance]), nil
}

// Post creates a record with associated relations.
func (rB *CmdbClassRequestBuilder) Post(ctx context.Context, body CmdbInstance, config *CmdbClassRequestBuilderPostRequestConfiguration) (*core.BaseServiceNowItemResponse[CmdbInstance], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[CmdbInstance](CreateCmdbInstanceFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(*core.BaseServiceNowItemResponse[CmdbInstance]), nil
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
	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

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
	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	if !conversion.IsNil(body) {
		err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
		if err != nil {
			return nil, err
		}
	}

	return requestInfo, nil
}

// ByID provides operations to manage a specific CI record.
func (rB *CmdbClassRequestBuilder) ByID(sysID string) *CmdbItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters[sysIDKey] = sysID
	return NewCmdbItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
