package actsubapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	actSubURLTemplate = "{+baseurl}/api/now/v1/actsub"
)

// ActSubRequestBuilder provides operations to manage Activity Subscriptions.
type ActSubRequestBuilder struct {
	internal.RequestBuilder
}

// NewActSubRequestBuilderInternal instantiates a new ActSubRequestBuilder.
func NewActSubRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ActSubRequestBuilder {
	return &ActSubRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, actSubURLTemplate, pathParameters),
	}
}

// Activities returns an ActivitiesRequestBuilder.
func (rB *ActSubRequestBuilder) Activities() *ActivitiesRequestBuilder {
	return NewActivitiesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Contexts returns a ContextsRequestBuilder.
func (rB *ActSubRequestBuilder) Contexts() *ContextsRequestBuilder {
	return NewContextsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Facets returns a FacetsRequestBuilder.
func (rB *ActSubRequestBuilder) Facets() *FacetsRequestBuilder {
	return NewFacetsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Followings returns a FollowingsRequestBuilder.
func (rB *ActSubRequestBuilder) Followings() *FollowingsRequestBuilder {
	return NewFollowingsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Preferences returns a PreferencesRequestBuilder.
func (rB *ActSubRequestBuilder) Preferences() *PreferencesRequestBuilder {
	return NewPreferencesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// SubObjects returns a SubObjectsRequestBuilder.
func (rB *ActSubRequestBuilder) SubObjects() *SubObjectsRequestBuilder {
	return NewSubObjectsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Subscribers returns a SubscribersRequestBuilder.
func (rB *ActSubRequestBuilder) Subscribers() *SubscribersRequestBuilder {
	return NewSubscribersRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Subscriptions returns a SubscriptionsRequestBuilder.
func (rB *ActSubRequestBuilder) Subscriptions() *SubscriptionsRequestBuilder {
	return NewSubscriptionsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// UserStream returns a UserStreamRequestBuilder.
func (rB *ActSubRequestBuilder) UserStream() *UserStreamRequestBuilder {
	return NewUserStreamRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
