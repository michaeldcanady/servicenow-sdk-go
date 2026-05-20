package cdmapplicationsapi

import (
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// UploadStatusResponse represents a response for the upload status.
type UploadStatusResponse interface {
	newInternal.ServiceNowItemResponse[*UploadStatusResult]
}

// CreateUploadStatusResponseFromDiscriminatorValue instantiates a new UploadStatusResponse.
func CreateUploadStatusResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*UploadStatusResult](CreateUploadStatusResultFromDiscriminatorValue), nil
}
