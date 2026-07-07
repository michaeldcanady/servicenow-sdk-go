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
	setContact(*string) error
	GetContactRPVariable() (RPVariable, error)
	setContactRPVariable(RPVariable) error
	GetLocation() (*string, error)
	setLocation(*string) error
	GetLocationRPVariable() (RPVariable, error)
	setLocationRPVariable(RPVariable) error
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
		internalSerialization.SerializeStringFunc(contactKey)(m.GetContact),
		internalSerialization.SerializeObjectValueFunc[RPVariable](contactRPVariableKey)(m.GetContactRPVariable),
		internalSerialization.SerializeStringFunc(locationKey)(m.GetLocation),
		internalSerialization.SerializeObjectValueFunc[RPVariable](locationRPVariableKey)(m.GetLocationRPVariable),
	)
}

func (m *FieldMappingModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		contactKey:            internalSerialization.DeserializeStringFunc()(m.setContact),
		contactRPVariableKey:  internalSerialization.DeserializeObjectValueFunc[RPVariable](CreateRPVariableFromDiscriminatorValue)(m.setContactRPVariable),
		locationKey:           internalSerialization.DeserializeStringFunc()(m.setLocation),
		locationRPVariableKey: internalSerialization.DeserializeObjectValueFunc[RPVariable](CreateRPVariableFromDiscriminatorValue)(m.setLocationRPVariable),
	}
}

func (m *FieldMappingModel) GetContact() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldMappingModel, *string](m, contactKey)
}
func (m *FieldMappingModel) setContact(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, contactKey, val)
}
func (m *FieldMappingModel) GetContactRPVariable() (RPVariable, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldMappingModel, RPVariable](m, contactRPVariableKey)
}
func (m *FieldMappingModel) setContactRPVariable(val RPVariable) error {
	return store.DefaultBackedModelMutatorFunc(m, contactRPVariableKey, val)
}
func (m *FieldMappingModel) GetLocation() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldMappingModel, *string](m, locationKey)
}
func (m *FieldMappingModel) setLocation(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, locationKey, val)
}
func (m *FieldMappingModel) GetLocationRPVariable() (RPVariable, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldMappingModel, RPVariable](m, locationRPVariableKey)
}
func (m *FieldMappingModel) setLocationRPVariable(val RPVariable) error {
	return store.DefaultBackedModelMutatorFunc(m, locationRPVariableKey, val)
}
