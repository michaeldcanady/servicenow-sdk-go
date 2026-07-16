package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// FindServiceResponse represents the response containing the found application service details.
type FindServiceResponse interface {
	core.ServiceNowItemResponse[*FindServiceResult]
}

func CreateFindServiceResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*FindServiceResult](CreateFindServiceResultFromDiscriminatorValue), nil
}
