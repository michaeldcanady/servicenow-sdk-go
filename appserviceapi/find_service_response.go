package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// FindServiceResponse represents the response containing the found application service details.
type FindServiceResponse = core.ServiceNowItemResponse[*FindServiceResult]

// CreateFindServiceResponseFromDiscriminatorValue creates a new FindServiceResponse from a ParseNode.
func CreateFindServiceResponseFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return core.ServiceNowItemResponseFromDiscriminatorValue[*FindServiceResult](CreateFindServiceResultFromDiscriminatorValue)(parseNode)
}
