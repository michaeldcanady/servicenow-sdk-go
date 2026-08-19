package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// UploadStatusResponse represents a response for the upload status.
type UploadStatusResponse interface {
	core.ServiceNowItemResponse[*UploadStatusResultModel]
}

// CreateUploadStatusResponseFromDiscriminatorValue instantiates a new UploadStatusResponse.
func CreateUploadStatusResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*UploadStatusResultModel](CreateUploadStatusResultFromDiscriminatorValue), nil
}
