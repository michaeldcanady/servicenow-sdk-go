package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
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

type ReferenceModel struct {
	core.BaseModel
}

func NewReference() *ReferenceModel {
	return &ReferenceModel{BaseModel: *core.NewBaseModel()}
}

func (m *ReferenceModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(linkKey, m.GetLink),
		internalSerialization.SerializeStringFunc(valueKey, m.GetValue),
	)
}

func (m *ReferenceModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		linkKey:  internalSerialization.DeserializeStringFunc(m.setLink),
		valueKey: internalSerialization.DeserializeStringFunc(m.setValue),
	}
}

func (m *ReferenceModel) GetLink() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ReferenceModel, *string](m, linkKey)
}
func (m *ReferenceModel) setLink(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, linkKey, val)
}
func (m *ReferenceModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ReferenceModel, *string](m, valueKey)
}
func (m *ReferenceModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}

func CreateReferenceFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewReference(), nil
}
