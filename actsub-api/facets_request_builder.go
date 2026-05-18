package actsubapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	facetsURLTemplate = "{+baseurl}/api/now/v1/actsub/facets"
)

// FacetsRequestBuilder provides operations to manage facets.
type FacetsRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewFacetsRequestBuilderInternal instantiates a new FacetsRequestBuilder.
func NewFacetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FacetsRequestBuilder {
	return &FacetsRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, facetsURLTemplate, pathParameters),
	}
}

// ByContext returns a FacetsContextRequestBuilder.
func (rB *FacetsRequestBuilder) ByContext(activityContext string) *FacetsContextRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["activity_context"] = activityContext
	return NewFacetsContextRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// FacetsContextRequestBuilder provides operations to manage facets for a specific context.
type FacetsContextRequestBuilder struct {
	newInternal.RequestBuilder
}

const facetsContextURLTemplate = "{+baseurl}/api/now/v1/actsub/facets/{activity_context}"

// NewFacetsContextRequestBuilderInternal instantiates a new FacetsContextRequestBuilder.
func NewFacetsContextRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FacetsContextRequestBuilder {
	return &FacetsContextRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, facetsContextURLTemplate, pathParameters),
	}
}

// ByInstance returns a FacetsInstanceRequestBuilder.
func (rB *FacetsContextRequestBuilder) ByInstance(contextInstance string) *FacetsInstanceRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["context_instance"] = contextInstance
	return NewFacetsInstanceRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// FacetsInstanceRequestBuilder provides operations to manage facets for a specific context instance.
type FacetsInstanceRequestBuilder struct {
	newInternal.RequestBuilder
}

const facetsInstanceURLTemplate = "{+baseurl}/api/now/v1/actsub/facets/{activity_context}/{context_instance}"

// NewFacetsInstanceRequestBuilderInternal instantiates a new FacetsInstanceRequestBuilder.
func NewFacetsInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FacetsInstanceRequestBuilder {
	return &FacetsInstanceRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, facetsInstanceURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve facets.
func (rB *FacetsInstanceRequestBuilder) Get(ctx context.Context, config *FacetsRequestBuilderGetRequestConfiguration) (*newInternal.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel], error) {
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, newInternal.ServiceNowCollectionResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), nil)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*newInternal.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel]), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *FacetsInstanceRequestBuilder) ToGetRequestInformation(ctx context.Context, config *FacetsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if headers := config.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParameters := config.QueryParameters; !internal.IsNil(queryParameters) {
			kiotaRequestInfo.AddQueryParameters(queryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return requestInfo, nil
}
