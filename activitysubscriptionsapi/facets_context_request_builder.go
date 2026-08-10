package actsubapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const facetsContextURLTemplate = "{+baseurl}/api/now/v1/actsub/facets/{activity_context}"

// FacetsContextRequestBuilder provides operations to manage facets for a specific context.
type FacetsContextRequestBuilder struct {
	core.RequestBuilder
}

// NewFacetsContextRequestBuilderInternal instantiates a new FacetsContextRequestBuilder.
func NewFacetsContextRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FacetsContextRequestBuilder {
	return &FacetsContextRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, facetsContextURLTemplate, pathParameters),
	}
}

// ByInstance returns a FacetsInstanceRequestBuilder.
func (rB *FacetsContextRequestBuilder) ByInstance(contextInstance string) *FacetsInstanceRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["context_instance"] = contextInstance
	return NewFacetsInstanceRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
