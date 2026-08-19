package cdmapplicationsapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
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
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewUploadsComponentsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Collections returns a [UploadsCollectionsRequestBuilder].
func (rB *UploadsRequestBuilder) Collections() *UploadsCollectionsRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewUploadsCollectionsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Deployables returns a [UploadsDeployablesRequestBuilder].
func (rB *UploadsRequestBuilder) Deployables() *UploadsDeployablesRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewUploadsDeployablesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
