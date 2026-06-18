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
	followingsURLTemplate = "{+baseurl}/api/now/v1/actsub/followings"
)

// FollowingsRequestBuilder provides operations to manage followings.
type FollowingsRequestBuilder struct {
	internal.RequestBuilder
}

// NewFollowingsRequestBuilderInternal instantiates a new FollowingsRequestBuilder.
func NewFollowingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FollowingsRequestBuilder {
	return &FollowingsRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, followingsURLTemplate, pathParameters),
	}
}

// ByFollower returns a FollowingItemRequestBuilder.
func (rB *FollowingsRequestBuilder) ByFollower(follower string) *FollowingItemRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["follower"] = follower
	return NewFollowingItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// FollowingItemRequestBuilder provides operations to manage following for a specific user.
type FollowingItemRequestBuilder struct {
	internal.RequestBuilder
}

const followingItemURLTemplate = "{+baseurl}/api/now/v1/actsub/followings/{follower}"

// NewFollowingItemRequestBuilderInternal instantiates a new FollowingItemRequestBuilder.
func NewFollowingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FollowingItemRequestBuilder {
	return &FollowingItemRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, followingItemURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve following.
func (rB *FollowingItemRequestBuilder) Get(ctx context.Context, config *FollowingsRequestBuilderGetRequestConfiguration) (*internal.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel], error) {
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
func (rB *FollowingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, config *FollowingsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return requestInfo, nil
}
