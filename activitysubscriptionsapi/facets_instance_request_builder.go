package activitysubscriptionsapi

import (
	"context"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const facetsInstanceURLTemplate = "{+baseurl}/api/now/v1/actsub/facets/{activity_context}/{context_instance}"

// FacetsInstanceRequestBuilder provides operations to manage facets for a specific context instance.
type FacetsInstanceRequestBuilder struct {
	core.RequestBuilder
}

// NewFacetsInstanceRequestBuilderInternal instantiates a new FacetsInstanceRequestBuilder.
func NewFacetsInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FacetsInstanceRequestBuilder {
	return &FacetsInstanceRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, facetsInstanceURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve facets.
func (rB *FacetsInstanceRequestBuilder) Get(ctx context.Context, config *FacetsInstanceRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowCollectionResponse[*ActivitySubscription], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowCollectionResponseFromDiscriminatorValue[*ActivitySubscription](CreateActivitySubscriptionFromDiscriminatorValue), core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}

	typedRes, ok := res.(*core.BaseServiceNowCollectionResponse[*ActivitySubscription])
	if !ok {
		return nil, fmt.Errorf("unexpected response type %T", res)
	}

	return typedRes, nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *FacetsInstanceRequestBuilder) ToGetRequestInformation(_ context.Context, config *FacetsInstanceRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
