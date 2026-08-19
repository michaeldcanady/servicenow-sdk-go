package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeployableUpdateResponse represents a response for deployable updates.
type DeployableUpdateResponse = core.ServiceNowItemResponse[*UploadStatusResultModel]

// CreateDeployableUpdateResponseFromDiscriminatorValue instantiates a new DeployableUpdateResponse.
func CreateDeployableUpdateResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*UploadStatusResultModel](CreateUploadStatusResultFromDiscriminatorValue), nil
}
