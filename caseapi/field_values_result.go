package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// FieldValuesResult represents field values.
type FieldValuesResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetLabel() (*string, error)
	setLabel(*string) error
	GetValue() (*string, error)
	setValue(*string) error
	GetSequence() (*int32, error)
	setSequence(*int32) error
	GetDependentValue() (*string, error)
	setDependentValue(*string) error
}

type FieldValuesResultModel struct {
	core.BaseModel
}

func NewFieldValuesResult() *FieldValuesResultModel {
	return &FieldValuesResultModel{BaseModel: *core.NewBaseModel()}
}

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

func (m *FieldValuesResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		labelKey:          internalSerialization.DeserializeStringFunc(m.setLabel),
		valueKey:          internalSerialization.DeserializeStringFunc(m.setValue),
		sequenceKey:       internalSerialization.DeserializeInt32Func(m.setSequence),
		dependentValueKey: internalSerialization.DeserializeStringFunc(m.setDependentValue),
	}
}

func (m *FieldValuesResultModel) GetLabel() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldValuesResultModel, *string](m, labelKey)
}
func (m *FieldValuesResultModel) setLabel(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, labelKey, val)
}
func (m *FieldValuesResultModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldValuesResultModel, *string](m, valueKey)
}
func (m *FieldValuesResultModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}
func (m *FieldValuesResultModel) GetSequence() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldValuesResultModel, *int32](m, sequenceKey)
}
func (m *FieldValuesResultModel) setSequence(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m, sequenceKey, val)
}
func (m *FieldValuesResultModel) GetDependentValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FieldValuesResultModel, *string](m, dependentValueKey)
}
func (m *FieldValuesResultModel) setDependentValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, dependentValueKey, val)
}

func CreateFieldValuesResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewFieldValuesResult(), nil
}
