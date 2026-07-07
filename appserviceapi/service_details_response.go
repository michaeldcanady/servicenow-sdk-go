package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceDetailsResponse represents the response containing service details update status.
type ServiceDetailsResponse interface {
	core.ServiceNowItemResponse[*ServiceDetailsResult]
}

func CreateServiceDetailsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*ServiceDetailsResult](CreateServiceDetailsResultFromDiscriminatorValue), nil
}
