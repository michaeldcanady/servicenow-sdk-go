package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExecuteRuleConditionsResult represents the execute rule conditions result.
type ExecuteRuleConditionsResult struct {
	core.BaseModel
}

// NewExecuteRuleConditionsResult creates a new instance of ExecuteRuleConditionsResult.
func NewExecuteRuleConditionsResult() *ExecuteRuleConditionsResult {
	return &ExecuteRuleConditionsResult{BaseModel: *core.NewBaseModel()}
}

// CreateExecuteRuleConditionsResultFromDiscriminatorValue creates a new ExecuteRuleConditionsResult from a ParseNode.
func CreateExecuteRuleConditionsResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewExecuteRuleConditionsResult(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *ExecuteRuleConditionsResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeBoolFunc(dedicatedCapacityKey, m.GetDedicatedCapacity),
		internalSerialization.SerializeStringFunc(futureMaxBookableDaysKey, m.GetFutureMaxBookableDays),
		internalSerialization.SerializeStringFunc(ruleIDKey, m.GetRuleID),
		internalSerialization.SerializeStringFunc(ruleNameKey, m.GetRuleName),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ExecuteRuleConditionsResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		dedicatedCapacityKey:     internalSerialization.DeserializeBoolFunc(m.SetDedicatedCapacity),
		futureMaxBookableDaysKey: internalSerialization.DeserializeStringFunc(m.SetFutureMaxBookableDays),
		ruleIDKey:                internalSerialization.DeserializeStringFunc(m.SetRuleID),
		ruleNameKey:              internalSerialization.DeserializeStringFunc(m.SetRuleName),
	}
}

// GetDedicatedCapacity returns the dedicated capacity value.
func (m *ExecuteRuleConditionsResult) GetDedicatedCapacity() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsResult, *bool](m, dedicatedCapacityKey)
}

// SetDedicatedCapacity sets the dedicated capacity value.
func (m *ExecuteRuleConditionsResult) SetDedicatedCapacity(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, dedicatedCapacityKey, val)
}

// GetFutureMaxBookableDays returns the future max bookable days value.
func (m *ExecuteRuleConditionsResult) GetFutureMaxBookableDays() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsResult, *string](m, futureMaxBookableDaysKey)
}

// SetFutureMaxBookableDays sets the future max bookable days value.
func (m *ExecuteRuleConditionsResult) SetFutureMaxBookableDays(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, futureMaxBookableDaysKey, val)
}

// GetRuleID returns the rule id value.
func (m *ExecuteRuleConditionsResult) GetRuleID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsResult, *string](m, ruleIDKey)
}

// SetRuleID sets the rule id value.
func (m *ExecuteRuleConditionsResult) SetRuleID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, ruleIDKey, val)
}

// GetRuleName returns the rule name value.
func (m *ExecuteRuleConditionsResult) GetRuleName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsResult, *string](m, ruleNameKey)
}

// SetRuleName sets the rule name value.
func (m *ExecuteRuleConditionsResult) SetRuleName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, ruleNameKey, val)
}
