package cdmapplicationsapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const exportItemURLTemplate = "{+baseurl}/api/sn_cdm/applications/deployables/exports/{export_id}"

// ExportItemRequestBuilder provides operations to manage a specific deployable export.
type ExportItemRequestBuilder struct {
	core.RequestBuilder
}

// NewExportItemRequestBuilderInternal instantiates a new [ExportItemRequestBuilder].
func NewExportItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ExportItemRequestBuilder {
	return &ExportItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, exportItemURLTemplate, pathParameters),
	}
}

// Content returns a [ExportItemContentRequestBuilder].
func (rB *ExportItemRequestBuilder) Content() *ExportItemContentRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewExportItemContentRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Status returns a [ExportItemStatusRequestBuilder].
func (rB *ExportItemRequestBuilder) Status() *ExportItemStatusRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewExportItemStatusRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
