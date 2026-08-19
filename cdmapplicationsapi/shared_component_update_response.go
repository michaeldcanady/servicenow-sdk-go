package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedComponentUpdateResponse represents a response for shared component updates.
type SharedComponentUpdateResponse = core.ServiceNowItemResponse[*UploadStatusResultModel]

// CreateSharedComponentUpdateResponseFromDiscriminatorValue instantiates a new SharedComponentUpdateResponse.
func CreateSharedComponentUpdateResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*UploadStatusResultModel](CreateUploadStatusResultFromDiscriminatorValue), nil
}
