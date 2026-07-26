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

	GetCatalogID() (*string, error)
	SetCatalogID(*string) error
	GetOtherInputs() (any, error)
	SetOtherInputs(any) error
	GetTaskID() (*string, error)
	SetTaskID(*string) error
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
		internalSerialization.SerializeStringFunc(catalogIDKey, m.GetCatalogID),
		internalSerialization.SerializeAnyFunc(otherInputsKey, m.GetOtherInputs),
		internalSerialization.SerializeStringFunc(taskIDKey, m.GetTaskID),
	)
}

func (m *ExecuteRuleConditionsRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		catalogIDKey:   internalSerialization.DeserializeStringFunc(m.SetCatalogID),
		otherInputsKey: internalSerialization.DeserializeAnyFunc(m.SetOtherInputs),
		taskIDKey:      internalSerialization.DeserializeStringFunc(m.SetTaskID),
	}
}

func (m *ExecuteRuleConditionsRequestModel) GetCatalogID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsRequestModel, *string](m, catalogIDKey)
}
func (m *ExecuteRuleConditionsRequestModel) SetCatalogID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, catalogIDKey, val)
}
func (m *ExecuteRuleConditionsRequestModel) GetOtherInputs() (any, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsRequestModel, any](m, otherInputsKey)
}
func (m *ExecuteRuleConditionsRequestModel) SetOtherInputs(val any) error {
	return store.DefaultBackedModelMutatorFunc(m, otherInputsKey, val)
}
func (m *ExecuteRuleConditionsRequestModel) GetTaskID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsRequestModel, *string](m, taskIDKey)
}
func (m *ExecuteRuleConditionsRequestModel) SetTaskID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskIDKey, val)
}
