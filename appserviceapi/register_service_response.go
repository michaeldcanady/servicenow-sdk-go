package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegisterServiceResponse represents the response containing the registered CSDM service details.
type RegisterServiceResponse interface {
	core.ServiceNowItemResponse[*RegisterServiceResult]
}

func CreateRegisterServiceResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*RegisterServiceResult](CreateRegisterServiceResultFromDiscriminatorValue), nil
}