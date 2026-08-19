package activitysubscriptionsapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	facetsURLTemplate = "{+baseurl}/api/now/v1/actsub/facets"
)

// FacetsRequestBuilder provides operations to manage facets.
type FacetsRequestBuilder struct {
	core.RequestBuilder
}

// NewFacetsRequestBuilderInternal instantiates a new [FacetsRequestBuilder].
func NewFacetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FacetsRequestBuilder {
	return &FacetsRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, facetsURLTemplate, pathParameters),
	}
}

// NewFacetsRequestBuilder instantiates a new [FacetsRequestBuilder] with a raw URL and request adapter.
func NewFacetsRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *FacetsRequestBuilder {
	return NewFacetsRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
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
