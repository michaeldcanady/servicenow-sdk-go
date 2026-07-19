package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// FieldMapping represents the field_mapping nested object.
type FieldMapping interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetContact() (*string, error)
	SetContact(*string) error
	GetContactRPVariable() (RPVariable, error)
	SetContactRPVariable(RPVariable) error
	GetLocation() (*string, error)
	SetLocation(*string) error
	GetLocationRPVariable() (RPVariable, error)
	SetLocationRPVariable(RPVariable) error
}

type FieldMappingModel struct {
	core.BaseModel
}

func NewFieldMapping() *FieldMappingModel {
	return &FieldMappingModel{
		BaseModel: *core.NewBaseModel(),
	}
}

func CreateFieldMappingFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewFieldMapping(), nil
}

func (m *FieldMappingModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(contactKey, m.GetContact),
		internalSerialization.SerializeObjectValueFunc[RPVariable](contactRPVariableKey, m.GetContactRPVariable),
		internalSerialization.SerializeStringFunc(locationKey, m.GetLocation),
		internalSerialization.SerializeObjectValueFunc[RPVariable](locationRPVariableKey, m.GetLocationRPVariable),
	)
}

func (m *FieldMappingModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		contactKey:            internalSerialization.DeserializeStringFunc(m.SetContact),
		contactRPVariableKey:  internalSerialization.DeserializeObjectValueFunc[RPVariable](CreateRPVariableFromDiscriminatorValue, m.SetContactRPVariable),
		locationKey:           internalSerialization.DeserializeStringFunc(m.SetLocation),
		locationRPVariableKey: internalSerialization.DeserializeObjectValueFunc[RPVariable](CreateRPVariableFromDiscriminatorValue, m.SetLocationRPVariable),
	}
}

func (m *FieldMappingModel) GetContact() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldMappingModel, *string](m, contactKey)
}
func (m *FieldMappingModel) SetContact(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, contactKey, val)
}
func (m *FieldMappingModel) GetContactRPVariable() (RPVariable, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldMappingModel, RPVariable](m, contactRPVariableKey)
}
func (m *FieldMappingModel) SetContactRPVariable(val RPVariable) error {
	return store.DefaultBackedModelMutatorFunc(m, contactRPVariableKey, val)
}
func (m *FieldMappingModel) GetLocation() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldMappingModel, *string](m, locationKey)
}
func (m *FieldMappingModel) SetLocation(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, locationKey, val)
}
func (m *FieldMappingModel) GetLocationRPVariable() (RPVariable, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldMappingModel, RPVariable](m, locationRPVariableKey)
}
func (m *FieldMappingModel) SetLocationRPVariable(val RPVariable) error {
	return store.DefaultBackedModelMutatorFunc(m, locationRPVariableKey, val)
}
