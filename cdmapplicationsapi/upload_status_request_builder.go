package cdmapplicationsapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const uploadStatusURLTemplate = "{+baseurl}/api/sn_cdm/applications/upload-status/{upload_id}"

// UploadStatusRequestBuilder provides operations to access upload status.
type UploadStatusRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadStatusRequestBuilderInternal instantiates a new [UploadStatusRequestBuilder].
func NewUploadStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadStatusRequestBuilder {
	return &UploadStatusRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadStatusURLTemplate, pathParameters),
	}
}

// ByID returns a [UploadStatusItemRequestBuilder].
func (rB *UploadStatusRequestBuilder) ByID(uploadID string) *UploadStatusItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["upload_id"] = uploadID
	return NewUploadStatusItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
