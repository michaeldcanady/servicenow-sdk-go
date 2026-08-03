package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// Reference represents a link and value pair.
type Reference struct {
	core.BaseModel
}

// NewReference instantiates a new Reference.
func NewReference() *Reference {
	return &Reference{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the object's properties to the given writer.
func (m *Reference) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(linkKey, m.GetLink),
		internalSerialization.SerializeStringFunc(valueKey, m.GetValue),
	)
}

// GetFieldDeserializers returns the deserializers for this object's fields.
func (m *Reference) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		linkKey:  internalSerialization.DeserializeStringFunc(m.setLink),
		valueKey: internalSerialization.DeserializeStringFunc(m.setValue),
	}
}

// GetLink returns the link.
func (m *Reference) GetLink() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Reference, *string](m, linkKey)
}
func (m *Reference) setLink(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, linkKey, val)
}

// GetValue returns the value.
func (m *Reference) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Reference, *string](m, valueKey)
}
func (m *Reference) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}

// CreateReferenceFromDiscriminatorValue creates a new Reference from a ParseNode.
func CreateReferenceFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewReference(), nil
}
