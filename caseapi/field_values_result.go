package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

var _ FieldValuesResult = (*FieldValuesResultModel)(nil)

// FieldValuesResult represents field values.
type FieldValuesResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetLabel() (*string, error)
	GetValue() (*string, error)
	GetSequence() (*int32, error)
	GetDependentValue() (*string, error)
}

// FieldValuesResultModel implementation of FieldValuesResult.
type FieldValuesResultModel struct {
	core.BaseModel
}

// NewFieldValuesResult creates a new instance of FieldValuesResult.
func NewFieldValuesResult() *FieldValuesResultModel {
	return &FieldValuesResultModel{BaseModel: *core.NewBaseModel()}
}

// CreateFieldValuesResultFromDiscriminatorValue creates a new instance of FieldValuesResult.
func CreateFieldValuesResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewFieldValuesResult(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *FieldValuesResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(labelKey, m.GetLabel),
		internalSerialization.SerializeStringFunc(valueKey, m.GetValue),
		internalSerialization.SerializeInt32Func(sequenceKey, m.GetSequence),
		internalSerialization.SerializeStringFunc(dependentValueKey, m.GetDependentValue),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *FieldValuesResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		labelKey:          internalSerialization.DeserializeStringFunc(m.setLabel),
		valueKey:          internalSerialization.DeserializeStringFunc(m.setValue),
		sequenceKey:       internalSerialization.DeserializeInt32Func(m.setSequence),
		dependentValueKey: internalSerialization.DeserializeStringFunc(m.setDependentValue),
	}
}

// GetLabel returns the display label for the field value.
func (m *FieldValuesResultModel) GetLabel() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldValuesResultModel, *string](m, labelKey)
}

func (m *FieldValuesResultModel) setLabel(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, labelKey, val)
}

// GetValue returns the value for the field.
func (m *FieldValuesResultModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldValuesResultModel, *string](m, valueKey)
}

func (m *FieldValuesResultModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}

// GetSequence returns the sequence number for ordering field values.
func (m *FieldValuesResultModel) GetSequence() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldValuesResultModel, *int32](m, sequenceKey)
}

func (m *FieldValuesResultModel) setSequence(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m, sequenceKey, val)
}

// GetDependentValue returns the dependent value for the field.
func (m *FieldValuesResultModel) GetDependentValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldValuesResultModel, *string](m, dependentValueKey)
}

func (m *FieldValuesResultModel) setDependentValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, dependentValueKey, val)
}
