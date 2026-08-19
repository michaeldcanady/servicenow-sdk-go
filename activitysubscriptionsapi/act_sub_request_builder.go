package activitysubscriptionsapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	actSubURLTemplate = "{+baseurl}/api/now/v1/actsub"
)

// ActSubRequestBuilder provides operations to manage Activity Subscriptions.
type ActSubRequestBuilder struct {
	core.RequestBuilder
}

// NewActSubRequestBuilderInternal instantiates a new [ActSubRequestBuilder].
func NewActSubRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ActSubRequestBuilder {
	return &ActSubRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, actSubURLTemplate, pathParameters),
	}
}

// NewActSubRequestBuilder instantiates a new [ActSubRequestBuilder] with a raw URL and request adapter.
func NewActSubRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *ActSubRequestBuilder {
	return NewActSubRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Activities returns an [ActivitiesRequestBuilder].
func (rB *ActSubRequestBuilder) Activities() *ActivitiesRequestBuilder {
	return NewActivitiesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Facets returns a [FacetsRequestBuilder].
func (rB *ActSubRequestBuilder) Facets() *FacetsRequestBuilder {
	return NewFacetsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
