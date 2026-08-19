package cdmapplicationsapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
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
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewDeployablesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// SharedComponents returns a [SharedComponentsRequestBuilder].
func (rB *ApplicationsRequestBuilder) SharedComponents() *SharedComponentsRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewSharedComponentsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// UploadStatus returns a [UploadStatusRequestBuilder].
func (rB *ApplicationsRequestBuilder) UploadStatus() *UploadStatusRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewUploadStatusRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// SharedLibraries returns a [SharedLibrariesRequestBuilder].
func (rB *ApplicationsRequestBuilder) SharedLibraries() *SharedLibrariesRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewSharedLibrariesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Uploads returns a [UploadsRequestBuilder].
func (rB *ApplicationsRequestBuilder) Uploads() *UploadsRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewUploadsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
