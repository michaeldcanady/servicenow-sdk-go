package cdmapplicationsapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const uploadsURLTemplate = "{+baseurl}/api/sn_cdm/applications/uploads"

// UploadsRequestBuilder provides operations to manage uploads.
type UploadsRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadsRequestBuilderInternal instantiates a new [UploadsRequestBuilder].
func NewUploadsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadsRequestBuilder {
	return &UploadsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadsURLTemplate, pathParameters),
	}
}

// Components returns a [UploadsComponentsRequestBuilder].
func (rB *UploadsRequestBuilder) Components() *UploadsComponentsRequestBuilder {
	return NewUploadsComponentsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Collections returns a [UploadsCollectionsRequestBuilder].
func (rB *UploadsRequestBuilder) Collections() *UploadsCollectionsRequestBuilder {
	return NewUploadsCollectionsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Deployables returns a [UploadsDeployablesRequestBuilder].
func (rB *UploadsRequestBuilder) Deployables() *UploadsDeployablesRequestBuilder {
	return NewUploadsDeployablesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
