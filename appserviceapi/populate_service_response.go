package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// PopulateServiceResponse represents the response containing populate result details.
type PopulateServiceResponse interface {
	core.ServiceNowItemResponse[*PopulateServiceResult]
}

func CreatePopulateServiceResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*PopulateServiceResult](CreatePopulateServiceResultFromDiscriminatorValue), nil
}
