package actsubapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	followingsURLTemplate = "{+baseurl}/api/now/v1/actsub/followings"
)

// FollowingsRequestBuilder provides operations to manage followings.
type FollowingsRequestBuilder struct {
	core.RequestBuilder
}

// NewFollowingsRequestBuilderInternal instantiates a new FollowingsRequestBuilder.
func NewFollowingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FollowingsRequestBuilder {
	return &FollowingsRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, followingsURLTemplate, pathParameters),
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
	*collectionGetRequestBuilder
}

const followingItemURLTemplate = "{+baseurl}/api/now/v1/actsub/followings/{follower}"

// NewFollowingItemRequestBuilderInternal instantiates a new FollowingItemRequestBuilder.
func NewFollowingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FollowingItemRequestBuilder {
	return &FollowingItemRequestBuilder{
		newCollectionGetRequestBuilder(pathParameters, requestAdapter, followingItemURLTemplate),
	}
}
