package actsubapi

import (
	"context"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
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

var _ core.CollectionGetRequestBuilder[*ActivitySubscriptionModel, abstractions.DefaultQueryParameters, abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]] = (*collectionGetRequestBuilder)(nil)

// newCollectionGetRequestBuilder instantiates a new collectionGetRequestBuilder.
func newCollectionGetRequestBuilder(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter, urlTemplate string) *collectionGetRequestBuilder {
	return &collectionGetRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, urlTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve the collection.
func (rB *collectionGetRequestBuilder) Get(ctx context.Context, config *abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]) (*core.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowCollectionResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*core.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel]), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *collectionGetRequestBuilder) ToGetRequestInformation(ctx context.Context, config *abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
