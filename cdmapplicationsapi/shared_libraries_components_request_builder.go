package cdmapplicationsapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const sharedLibrariesComponentsURLTemplate = "{+baseurl}/api/sn_cdm/applications/shared_libraries/components"

// SharedLibrariesComponentsRequestBuilder provides operations to manage shared library components.
type SharedLibrariesComponentsRequestBuilder struct {
	core.RequestBuilder
}

// NewSharedLibrariesComponentsRequestBuilderInternal instantiates a new [SharedLibrariesComponentsRequestBuilder].
func NewSharedLibrariesComponentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SharedLibrariesComponentsRequestBuilder {
	return &SharedLibrariesComponentsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, sharedLibrariesComponentsURLTemplate, pathParameters),
	}
}

// Applications returns a [SharedLibrariesComponentsApplicationsRequestBuilder].
func (rB *SharedLibrariesComponentsRequestBuilder) Applications() *SharedLibrariesComponentsApplicationsRequestBuilder {
	return NewSharedLibrariesComponentsApplicationsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
