package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetContentResponse represents the response containing the content of an application service.
type GetContentResponse = core.ServiceNowItemResponse[*GetContentResult]

// CreateGetContentResponseFromDiscriminatorValue creates a new GetContentResponse from a ParseNode.
func CreateGetContentResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*GetContentResult](CreateGetContentResultFromDiscriminatorValue), nil
}
