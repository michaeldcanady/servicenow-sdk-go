package cdmapplicationsapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const applicationsURLTemplate = "{+baseurl}/api/sn_cdm/applications"

// ApplicationsRequestBuilder provides operations to manage applications.
type ApplicationsRequestBuilder struct {
	core.RequestBuilder
}

// NewApplicationsRequestBuilderInternal instantiates a new [ApplicationsRequestBuilder].
func NewApplicationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ApplicationsRequestBuilder {
	return &ApplicationsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, applicationsURLTemplate, pathParameters),
	}
}

// Deployables returns a [DeployablesRequestBuilder].
func (rB *ApplicationsRequestBuilder) Deployables() *DeployablesRequestBuilder {
	return NewDeployablesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// SharedComponents returns a [SharedComponentsRequestBuilder].
func (rB *ApplicationsRequestBuilder) SharedComponents() *SharedComponentsRequestBuilder {
	return NewSharedComponentsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// UploadStatus returns a [UploadStatusRequestBuilder].
func (rB *ApplicationsRequestBuilder) UploadStatus() *UploadStatusRequestBuilder {
	return NewUploadStatusRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// SharedLibraries returns a [SharedLibrariesRequestBuilder].
func (rB *ApplicationsRequestBuilder) SharedLibraries() *SharedLibrariesRequestBuilder {
	return NewSharedLibrariesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Uploads returns a [UploadsRequestBuilder].
func (rB *ApplicationsRequestBuilder) Uploads() *UploadsRequestBuilder {
	return NewUploadsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
