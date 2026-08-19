package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegisterServiceResponse represents the response containing the registered CSDM service details.
type RegisterServiceResponse interface {
	core.ServiceNowItemResponse[*RegisterServiceResult]
}

// CreateRegisterServiceResponseFromDiscriminatorValue creates a new RegisterServiceResponse from a ParseNode.
func CreateRegisterServiceResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*RegisterServiceResult](CreateRegisterServiceResultFromDiscriminatorValue), nil
}
