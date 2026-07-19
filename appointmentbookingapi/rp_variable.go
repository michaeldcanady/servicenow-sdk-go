package appointmentbookingapi // nolint:dupl // shares field-count shape with UserTimeFormatOptionsModel by coincidence, not copy-paste; distinct API concept, not worth sacrificing named accessors for

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// RPVariable represents the RPVariable nested object.
type RPVariable interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetDisplayName() (*string, error)
	SetDisplayName(*string) error
	GetLabel() (*string, error)
	SetLabel(*string) error
	GetName() (*string, error)
	SetName(*string) error
}

type RPVariableModel struct {
	core.BaseModel
}

func NewRPVariable() *RPVariableModel {
	return &RPVariableModel{
		BaseModel: *core.NewBaseModel(),
	}
}

func CreateRPVariableFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRPVariable(), nil
}

func (m *RPVariableModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(displayNameKey, m.GetDisplayName),
		internalSerialization.SerializeStringFunc(labelKey, m.GetLabel),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
	)
}

func (m *RPVariableModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		displayNameKey: internalSerialization.DeserializeStringFunc(m.SetDisplayName),
		labelKey:       internalSerialization.DeserializeStringFunc(m.SetLabel),
		nameKey:        internalSerialization.DeserializeStringFunc(m.SetName),
	}
}

func (m *RPVariableModel) GetDisplayName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*RPVariableModel, *string](m, displayNameKey)
}
func (m *RPVariableModel) SetDisplayName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, displayNameKey, val)
}
func (m *RPVariableModel) GetLabel() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*RPVariableModel, *string](m, labelKey)
}
func (m *RPVariableModel) SetLabel(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, labelKey, val)
}
func (m *RPVariableModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*RPVariableModel, *string](m, nameKey)
}
func (m *RPVariableModel) SetName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}
