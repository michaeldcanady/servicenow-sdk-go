package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppServiceActionResponse represents the response returned by an app_service action endpoint.
type AppServiceActionResponse = core.ServiceNowItemResponse[*AppServiceActionResult]

// CreateAppServiceActionResponseFromDiscriminatorValue creates a new AppServiceActionResponse from a ParseNode.
func CreateAppServiceActionResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*AppServiceActionResult](CreateAppServiceActionResultFromDiscriminatorValue), nil
}
