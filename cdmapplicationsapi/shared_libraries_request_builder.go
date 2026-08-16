package cdmapplicationsapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const sharedLibrariesURLTemplate = "{+baseurl}/api/sn_cdm/applications/shared_libraries"

// SharedLibrariesRequestBuilder provides operations to manage shared libraries.
type SharedLibrariesRequestBuilder struct {
	core.RequestBuilder
}

// NewSharedLibrariesRequestBuilderInternal instantiates a new [SharedLibrariesRequestBuilder].
func NewSharedLibrariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SharedLibrariesRequestBuilder {
	return &SharedLibrariesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, sharedLibrariesURLTemplate, pathParameters),
	}
}

// Components returns a [SharedLibrariesComponentsRequestBuilder].
func (rB *SharedLibrariesRequestBuilder) Components() *SharedLibrariesComponentsRequestBuilder {
	return NewSharedLibrariesComponentsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
