package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

type ExecuteRuleConditionsResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetDedicatedCapacity() (*bool, error)
	setDedicatedCapacity(*bool) error
	GetFutureMaxBookableDays() (*string, error)
	setFutureMaxBookableDays(*string) error
	GetRuleId() (*string, error)
	setRuleId(*string) error
	GetRuleName() (*string, error)
	setRuleName(*string) error
}

type ExecuteRuleConditionsResultModel struct {
	core.BaseModel
}

func NewExecuteRuleConditionsResult() *ExecuteRuleConditionsResultModel {
	return &ExecuteRuleConditionsResultModel{BaseModel: *core.NewBaseModel()}
}

func CreateExecuteRuleConditionsResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewExecuteRuleConditionsResult(), nil
}

func (m *ExecuteRuleConditionsResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeBoolFunc(dedicatedCapacityKey)(m.GetDedicatedCapacity),
		internalSerialization.SerializeStringFunc(futureMaxBookableDaysKey)(m.GetFutureMaxBookableDays),
		internalSerialization.SerializeStringFunc(ruleIdKey)(m.GetRuleId),
		internalSerialization.SerializeStringFunc(ruleNameKey)(m.GetRuleName),
	)
}

func (m *ExecuteRuleConditionsResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		dedicatedCapacityKey:     internalSerialization.DeserializeBoolFunc()(m.setDedicatedCapacity),
		futureMaxBookableDaysKey: internalSerialization.DeserializeStringFunc()(m.setFutureMaxBookableDays),
		ruleIdKey:                internalSerialization.DeserializeStringFunc()(m.setRuleId),
		ruleNameKey:              internalSerialization.DeserializeStringFunc()(m.setRuleName),
	}
}

func (m *ExecuteRuleConditionsResultModel) GetDedicatedCapacity() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsResultModel, *bool](m, dedicatedCapacityKey)
}
func (m *ExecuteRuleConditionsResultModel) setDedicatedCapacity(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, dedicatedCapacityKey, val)
}
func (m *ExecuteRuleConditionsResultModel) GetFutureMaxBookableDays() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsResultModel, *string](m, futureMaxBookableDaysKey)
}
func (m *ExecuteRuleConditionsResultModel) setFutureMaxBookableDays(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, futureMaxBookableDaysKey, val)
}
func (m *ExecuteRuleConditionsResultModel) GetRuleId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsResultModel, *string](m, ruleIdKey)
}
func (m *ExecuteRuleConditionsResultModel) setRuleId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, ruleIdKey, val)
}
func (m *ExecuteRuleConditionsResultModel) GetRuleName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsResultModel, *string](m, ruleNameKey)
}
func (m *ExecuteRuleConditionsResultModel) setRuleName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, ruleNameKey, val)
}
