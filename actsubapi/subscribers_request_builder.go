package actsubapi // nolint:dupl // FollowingItemRequestBuilder/SubscriberItemRequestBuilder's nil-guard wrapper methods are inherently near-identical (each depends on which specific outer type is nil, so it can't be extracted into the shared collectionGetRequestBuilder)

import (
	"context"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	subscribersURLTemplate = "{+baseurl}/api/now/v1/actsub/subscribers"
)

// SubscribersRequestBuilder provides operations to manage subscribers.
type SubscribersRequestBuilder struct {
	core.RequestBuilder
}

// NewSubscribersRequestBuilderInternal instantiates a new SubscribersRequestBuilder.
func NewSubscribersRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SubscribersRequestBuilder {
	return &SubscribersRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, subscribersURLTemplate, pathParameters),
	}
}

// BySubObject returns a SubscriberItemRequestBuilder.
func (rB *SubscribersRequestBuilder) BySubObject(subObject string) *SubscriberItemRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["sub_object"] = subObject
	return NewSubscriberItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// SubscriberItemRequestBuilder provides operations to manage subscribers for a specific object.
type SubscriberItemRequestBuilder struct {
	*collectionGetRequestBuilder
}

const subscriberItemURLTemplate = "{+baseurl}/api/now/v1/actsub/subscribers/{sub_object}"

// NewSubscriberItemRequestBuilderInternal instantiates a new SubscriberItemRequestBuilder.
func NewSubscriberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SubscriberItemRequestBuilder {
	return &SubscriberItemRequestBuilder{
		newCollectionGetRequestBuilder(pathParameters, requestAdapter, subscriberItemURLTemplate),
	}
}

// Get sends a GET request to retrieve subscribers.
func (rB *SubscriberItemRequestBuilder) Get(ctx context.Context, config *SubscribersRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.collectionGetRequestBuilder.Get(ctx, config)
}

// ToGetRequestInformation creates a RequestInformation object for a GET request to retrieve subscribers.
func (rB *SubscriberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, config *SubscribersRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.collectionGetRequestBuilder.ToGetRequestInformation(ctx, config)
}
