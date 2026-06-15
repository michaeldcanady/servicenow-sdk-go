package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

type ExecuteRuleConditionsRequest interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetCatalogId() (*string, error)
	setCatalogId(*string) error
	GetOtherInputs() (any, error)
	setOtherInputs(any) error
	GetTaskId() (*string, error)
	setTaskId(*string) error
}

type ExecuteRuleConditionsRequestModel struct {
	internal.BaseModel
}

func NewExecuteRuleConditionsRequest() *ExecuteRuleConditionsRequestModel {
	return &ExecuteRuleConditionsRequestModel{BaseModel: *internal.NewBaseModel()}
}

func CreateExecuteRuleConditionsRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewExecuteRuleConditionsRequest(), nil
}

func (m *ExecuteRuleConditionsRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(catalogIdKey)(m.GetCatalogId),
		internalSerialization.SerializeAnyFunc(otherInputsKey)(m.GetOtherInputs),
		internalSerialization.SerializeStringFunc(taskIdKey)(m.GetTaskId),
	)
}

func (m *ExecuteRuleConditionsRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		catalogIdKey:   internalSerialization.DeserializeStringFunc()(m.setCatalogId),
		otherInputsKey: internalSerialization.DeserializeAnyFunc()(m.setOtherInputs),
		taskIdKey:      internalSerialization.DeserializeStringFunc()(m.setTaskId),
	}
}

func (m *ExecuteRuleConditionsRequestModel) GetCatalogId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), catalogIdKey)
}
func (m *ExecuteRuleConditionsRequestModel) setCatalogId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), catalogIdKey, val)
}
func (m *ExecuteRuleConditionsRequestModel) GetOtherInputs() (any, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, any](m.GetBackingStore(), otherInputsKey)
}
func (m *ExecuteRuleConditionsRequestModel) setOtherInputs(val any) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), otherInputsKey, val)
}
func (m *ExecuteRuleConditionsRequestModel) GetTaskId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), taskIdKey)
}
func (m *ExecuteRuleConditionsRequestModel) setTaskId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), taskIdKey, val)
}

// ExecuteRuleConditionsResponse represents the response from execute_rule_conditions.
type ExecuteRuleConditionsResponse = internal.ServiceNowItemResponse[*ExecuteRuleConditionsResultModel]

// CreateExecuteRuleConditionsResponseFromDiscriminatorValue is a factory.
func CreateExecuteRuleConditionsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowItemResponse[*ExecuteRuleConditionsResultModel](CreateExecuteRuleConditionsResultFromDiscriminatorValue), nil
}

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
	internal.BaseModel
}

func NewExecuteRuleConditionsResult() *ExecuteRuleConditionsResultModel {
	return &ExecuteRuleConditionsResultModel{BaseModel: *internal.NewBaseModel()}
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
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), dedicatedCapacityKey)
}
func (m *ExecuteRuleConditionsResultModel) setDedicatedCapacity(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), dedicatedCapacityKey, val)
}
func (m *ExecuteRuleConditionsResultModel) GetFutureMaxBookableDays() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), futureMaxBookableDaysKey)
}
func (m *ExecuteRuleConditionsResultModel) setFutureMaxBookableDays(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), futureMaxBookableDaysKey, val)
}
func (m *ExecuteRuleConditionsResultModel) GetRuleId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), ruleIdKey)
}
func (m *ExecuteRuleConditionsResultModel) setRuleId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), ruleIdKey, val)
}
func (m *ExecuteRuleConditionsResultModel) GetRuleName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), ruleNameKey)
}
func (m *ExecuteRuleConditionsResultModel) setRuleName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), ruleNameKey, val)
}
