package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExportStatusResponse represents a response containing an export status result.
type ExportStatusResponse interface {
	core.ServiceNowItemResponse[*ExportStatusResult]
}

// CreateExportStatusResponseFromDiscriminatorValue instantiates a new ExportStatusResponse.
func CreateExportStatusResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*ExportStatusResult](CreateExportStatusResultFromDiscriminatorValue), nil
}
