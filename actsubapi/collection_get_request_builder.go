package actsubapi

import (
	"context"
	"fmt"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// collectionGetRequestBuilder is the shared implementation behind every actsubapi
// request builder whose only operation is a GET returning a collection of
// ActivitySubscriptionModel (Activities, Contexts, SubObjects, FollowingItem,
// SubscriberItem) - they differ only in URL template.
type collectionGetRequestBuilder struct {
	core.RequestBuilder
}

var _ core.CollectionGetRequestBuilder[*ActivitySubscription, abstractions.DefaultQueryParameters, abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]] = (*collectionGetRequestBuilder)(nil)

// TODO: what is this?
// newCollectionGetRequestBuilder instantiates a new collectionGetRequestBuilder.
func newCollectionGetRequestBuilder(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter, urlTemplate string) *collectionGetRequestBuilder {
	return &collectionGetRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, urlTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve the collection.
func (rB *collectionGetRequestBuilder) Get(ctx context.Context, config *abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]) (*core.BaseServiceNowCollectionResponse[*ActivitySubscription], error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowCollectionResponseFromDiscriminatorValue[*ActivitySubscription](CreateActivitySubscriptionModelFromDiscriminatorValue), core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	typedRes, ok := res.(*core.BaseServiceNowCollectionResponse[*ActivitySubscription])
	if !ok {
		return nil, fmt.Errorf("unexpected response type %T", res)
	}

	return typedRes, nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *collectionGetRequestBuilder) ToGetRequestInformation(_ context.Context, config *abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
