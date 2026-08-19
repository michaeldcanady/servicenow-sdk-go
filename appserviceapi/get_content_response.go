package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetContentResponse represents the response containing the content of an application service.
type GetContentResponse interface {
	core.ServiceNowItemResponse[*GetContentResult]
}

// CreateGetContentResponseFromDiscriminatorValue creates a new GetContentResponse from a ParseNode.
func CreateGetContentResponseFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return core.ServiceNowItemResponseFromDiscriminatorValue[*GetContentResult](CreateGetContentResultFromDiscriminatorValue)(parseNode)
}
