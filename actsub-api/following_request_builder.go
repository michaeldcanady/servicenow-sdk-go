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
	followingsURLTemplate = "{+baseurl}/api/now/v1/actsub/followings"
)

// FollowingsRequestBuilder provides operations to manage followings.
type FollowingsRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewFollowingsRequestBuilderInternal instantiates a new FollowingsRequestBuilder.
func NewFollowingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FollowingsRequestBuilder {
	return &FollowingsRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, followingsURLTemplate, pathParameters),
	}
}

// ByFollower returns a FollowingItemRequestBuilder.
func (rB *FollowingsRequestBuilder) ByFollower(follower string) *FollowingItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["follower"] = follower
	return NewFollowingItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// FollowingItemRequestBuilder provides operations to manage following for a specific user.
type FollowingItemRequestBuilder struct {
	newInternal.RequestBuilder
}

const followingItemURLTemplate = "{+baseurl}/api/now/v1/actsub/followings/{follower}"

// NewFollowingItemRequestBuilderInternal instantiates a new FollowingItemRequestBuilder.
func NewFollowingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FollowingItemRequestBuilder {
	return &FollowingItemRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, followingItemURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve following.
func (rB *FollowingItemRequestBuilder) Get(ctx context.Context, config *FollowingsRequestBuilderGetRequestConfiguration) (*newInternal.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel], error) {
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
func (rB *FollowingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, config *FollowingsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
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
