package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExecuteRuleConditionsResponse represents the response from execute_rule_conditions.
type ExecuteRuleConditionsResponse = core.ServiceNowItemResponse[*ExecuteRuleConditionsResult]

// CreateExecuteRuleConditionsResponseFromDiscriminatorValue is a factory.
func CreateExecuteRuleConditionsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*ExecuteRuleConditionsResult](CreateExecuteRuleConditionsResultFromDiscriminatorValue), nil
}
