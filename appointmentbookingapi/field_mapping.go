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

	GetContact() (*string, error) // V
	SetContact(*string) error
	GetContactRPVariable() (RPVariable, error) // V
	SetContactRPVariable(RPVariable) error
	GetLocation() (*string, error) // V
	SetLocation(*string) error
	GetLocationRPVariable() (RPVariable, error) // V
	SetLocationRPVariable(RPVariable) error
}

// FieldMappingModel represents the field mapping model.
type FieldMappingModel struct {
	core.BaseModel
}

// NewFieldMapping creates a new instance of FieldMappingModel.
func NewFieldMapping() *FieldMappingModel {
	return &FieldMappingModel{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateFieldMappingFromDiscriminatorValue creates a new FieldMapping from a ParseNode.
func CreateFieldMappingFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewFieldMapping(), nil
}

// Serialize writes the objects properties to the current writer.
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

// GetFieldDeserializers returns the deserialization information for this object.
func (m *FieldMappingModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		contactKey:            internalSerialization.DeserializeStringFunc(m.SetContact),
		contactRPVariableKey:  internalSerialization.DeserializeObjectValueFunc[RPVariable](CreateRPVariableFromDiscriminatorValue, m.SetContactRPVariable),
		locationKey:           internalSerialization.DeserializeStringFunc(m.SetLocation),
		locationRPVariableKey: internalSerialization.DeserializeObjectValueFunc[RPVariable](CreateRPVariableFromDiscriminatorValue, m.SetLocationRPVariable),
	}
}

// GetContact returns the contact value.
func (m *FieldMappingModel) GetContact() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldMappingModel, *string](m, contactKey)
}

// SetContact sets the contact value.
func (m *FieldMappingModel) SetContact(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, contactKey, val)
}

// GetContactRPVariable returns the contact rp variable value.
func (m *FieldMappingModel) GetContactRPVariable() (RPVariable, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldMappingModel, RPVariable](m, contactRPVariableKey)
}

// SetContactRPVariable sets the contact rp variable value.
func (m *FieldMappingModel) SetContactRPVariable(val RPVariable) error {
	return store.DefaultBackedModelMutatorFunc(m, contactRPVariableKey, val)
}

// GetLocation returns the location value.
func (m *FieldMappingModel) GetLocation() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldMappingModel, *string](m, locationKey)
}

// SetLocation sets the location value.
func (m *FieldMappingModel) SetLocation(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, locationKey, val)
}

// GetLocationRPVariable returns the location rp variable value.
func (m *FieldMappingModel) GetLocationRPVariable() (RPVariable, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldMappingModel, RPVariable](m, locationRPVariableKey)
}

// SetLocationRPVariable sets the location rp variable value.
func (m *FieldMappingModel) SetLocationRPVariable(val RPVariable) error {
	return store.DefaultBackedModelMutatorFunc(m, locationRPVariableKey, val)
}
