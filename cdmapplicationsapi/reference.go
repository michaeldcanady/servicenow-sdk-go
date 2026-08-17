package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	linkKey  = "link"
	valueKey = "value"
)

// Reference represents a reference field object.
type Reference interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetLink() (*string, error)
	setLink(*string) error
	GetValue() (*string, error)
	setValue(*string) error
}

// ReferenceModel is the default implementation of Reference.
type ReferenceModel struct {
	core.BaseModel
}

// NewReference creates a new instance of ReferenceModel.
func NewReference() *ReferenceModel {
	return &ReferenceModel{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the object's properties to the current writer.
func (m *ReferenceModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(linkKey, m.GetLink),
		internalSerialization.SerializeStringFunc(valueKey, m.GetValue),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ReferenceModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		linkKey:  internalSerialization.DeserializeStringFunc(m.setLink),
		valueKey: internalSerialization.DeserializeStringFunc(m.setValue),
	}
}

// GetLink returns the reference link.
func (m *ReferenceModel) GetLink() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ReferenceModel, *string](m, linkKey)
}

func (m *ReferenceModel) setLink(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, linkKey, val)
}

// GetValue returns the reference value.
func (m *ReferenceModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ReferenceModel, *string](m, valueKey)
}

func (m *ReferenceModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}

// CreateReferenceFromDiscriminatorValue creates a new instance of Reference.
func CreateReferenceFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewReference(), nil
}
