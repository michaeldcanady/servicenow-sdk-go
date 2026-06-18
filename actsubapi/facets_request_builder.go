package actsubapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	facetsURLTemplate = "{+baseurl}/api/now/v1/actsub/facets"
)

// FacetsRequestBuilder provides operations to manage facets.
type FacetsRequestBuilder struct {
	internal.RequestBuilder
}

// NewFacetsRequestBuilderInternal instantiates a new FacetsRequestBuilder.
func NewFacetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FacetsRequestBuilder {
	return &FacetsRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, facetsURLTemplate, pathParameters),
	}
}

// ByContext returns a FacetsContextRequestBuilder.
func (rB *FacetsRequestBuilder) ByContext(activityContext string) *FacetsContextRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["activity_context"] = activityContext
	return NewFacetsContextRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// FacetsContextRequestBuilder provides operations to manage facets for a specific context.
type FacetsContextRequestBuilder struct {
	internal.RequestBuilder
}

const facetsContextURLTemplate = "{+baseurl}/api/now/v1/actsub/facets/{activity_context}"

// NewFacetsContextRequestBuilderInternal instantiates a new FacetsContextRequestBuilder.
func NewFacetsContextRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FacetsContextRequestBuilder {
	return &FacetsContextRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, facetsContextURLTemplate, pathParameters),
	}
}

// ByInstance returns a FacetsInstanceRequestBuilder.
func (rB *FacetsContextRequestBuilder) ByInstance(contextInstance string) *FacetsInstanceRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["context_instance"] = contextInstance
	return NewFacetsInstanceRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// FacetsInstanceRequestBuilder provides operations to manage facets for a specific context instance.
type FacetsInstanceRequestBuilder struct {
	internal.RequestBuilder
}

const facetsInstanceURLTemplate = "{+baseurl}/api/now/v1/actsub/facets/{activity_context}/{context_instance}"

// NewFacetsInstanceRequestBuilderInternal instantiates a new FacetsInstanceRequestBuilder.
func NewFacetsInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FacetsInstanceRequestBuilder {
	return &FacetsInstanceRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, facetsInstanceURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve facets.
func (rB *FacetsInstanceRequestBuilder) Get(ctx context.Context, config *FacetsRequestBuilderGetRequestConfiguration) (*internal.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowCollectionResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), internal.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*internal.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel]), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *FacetsInstanceRequestBuilder) ToGetRequestInformation(ctx context.Context, config *FacetsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return requestInfo, nil
}
