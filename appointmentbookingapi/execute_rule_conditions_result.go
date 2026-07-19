package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExecuteRuleConditionsResult struct {
	core.BaseModel
}

func NewExecuteRuleConditionsResult() *ExecuteRuleConditionsResult {
	return &ExecuteRuleConditionsResult{BaseModel: *core.NewBaseModel()}
}

func CreateExecuteRuleConditionsResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewExecuteRuleConditionsResult(), nil
}

func (m *ExecuteRuleConditionsResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeBoolFunc(dedicatedCapacityKey, m.GetDedicatedCapacity),
		internalSerialization.SerializeStringFunc(futureMaxBookableDaysKey, m.GetFutureMaxBookableDays),
		internalSerialization.SerializeStringFunc(ruleIdKey, m.GetRuleId),
		internalSerialization.SerializeStringFunc(ruleNameKey, m.GetRuleName),
	)
}

func (m *ExecuteRuleConditionsResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		dedicatedCapacityKey:     internalSerialization.DeserializeBoolFunc(m.SetDedicatedCapacity),
		futureMaxBookableDaysKey: internalSerialization.DeserializeStringFunc(m.SetFutureMaxBookableDays),
		ruleIdKey:                internalSerialization.DeserializeStringFunc(m.SetRuleId),
		ruleNameKey:              internalSerialization.DeserializeStringFunc(m.SetRuleName),
	}
}

func (m *ExecuteRuleConditionsResult) GetDedicatedCapacity() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsResult, *bool](m, dedicatedCapacityKey)
}
func (m *ExecuteRuleConditionsResult) SetDedicatedCapacity(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, dedicatedCapacityKey, val)
}
func (m *ExecuteRuleConditionsResult) GetFutureMaxBookableDays() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsResult, *string](m, futureMaxBookableDaysKey)
}
func (m *ExecuteRuleConditionsResult) SetFutureMaxBookableDays(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, futureMaxBookableDaysKey, val)
}
func (m *ExecuteRuleConditionsResult) GetRuleId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsResult, *string](m, ruleIdKey)
}
func (m *ExecuteRuleConditionsResult) SetRuleId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, ruleIdKey, val)
}
func (m *ExecuteRuleConditionsResult) GetRuleName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsResult, *string](m, ruleNameKey)
}
func (m *ExecuteRuleConditionsResult) SetRuleName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, ruleNameKey, val)
}
