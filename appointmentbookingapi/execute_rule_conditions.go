package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
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
	core.BaseModel
}

func NewExecuteRuleConditionsRequest() *ExecuteRuleConditionsRequestModel {
	return &ExecuteRuleConditionsRequestModel{BaseModel: *core.NewBaseModel()}
}

func CreateExecuteRuleConditionsRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewExecuteRuleConditionsRequest(), nil
}

func (m *ExecuteRuleConditionsRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(catalogIDKey)(m.GetCatalogId),
		internalSerialization.SerializeAnyFunc(otherInputsKey)(m.GetOtherInputs),
		internalSerialization.SerializeStringFunc(taskIDKey)(m.GetTaskId),
	)
}

func (m *ExecuteRuleConditionsRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		catalogIDKey:   internalSerialization.DeserializeStringFunc()(m.setCatalogId),
		otherInputsKey: internalSerialization.DeserializeAnyFunc()(m.setOtherInputs),
		taskIDKey:      internalSerialization.DeserializeStringFunc()(m.setTaskId),
	}
}

func (m *ExecuteRuleConditionsRequestModel) GetCatalogId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsRequestModel, *string](m, catalogIDKey)
}
func (m *ExecuteRuleConditionsRequestModel) setCatalogId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, catalogIDKey, val)
}
func (m *ExecuteRuleConditionsRequestModel) GetOtherInputs() (any, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsRequestModel, any](m, otherInputsKey)
}
func (m *ExecuteRuleConditionsRequestModel) setOtherInputs(val any) error {
	return store.DefaultBackedModelMutatorFunc(m, otherInputsKey, val)
}
func (m *ExecuteRuleConditionsRequestModel) GetTaskId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsRequestModel, *string](m, taskIDKey)
}
func (m *ExecuteRuleConditionsRequestModel) setTaskId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskIDKey, val)
}
