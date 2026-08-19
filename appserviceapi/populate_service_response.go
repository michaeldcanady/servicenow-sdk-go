package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// PopulateServiceResponse represents the response containing populate result details.
type PopulateServiceResponse = core.ServiceNowItemResponse[*PopulateServiceResult]

// CreatePopulateServiceResponseFromDiscriminatorValue creates a new PopulateServiceResponse from a ParseNode.
func CreatePopulateServiceResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*PopulateServiceResult](CreatePopulateServiceResultFromDiscriminatorValue), nil
}
