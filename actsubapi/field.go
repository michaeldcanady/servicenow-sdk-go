package actsubapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	deepLinkToSubObjectKey = "deeplink_to_subobject"
	displayAsTimeAgoKey    = "display_as_timeago"
	labelKey               = "label"
	showLabelKey           = "show_label"
	typeKey                = "type"
	valueKey               = "value"
)

type Field struct {
	core.BaseModel
}

func NewField() *Field {
	return &Field{
		BaseModel: *core.NewBaseModel(),
	}
}

func CreateFieldFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewField(), nil
}

func (m *Field) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return nil
}

func (m *Field) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		deepLinkToSubObjectKey: internalSerialization.DeserializeStringFunc(m.SetDeepLinkToSubObject),
		displayAsTimeAgoKey:    internalSerialization.DeserializeStringFunc(m.SetDisplayAsTimeAgo),
		labelKey:               internalSerialization.DeserializeStringFunc(m.SetLabel),
		showLabelKey:           internalSerialization.DeserializeStringFunc(m.SetShowLabel),
		typeKey:                internalSerialization.DeserializeStringFunc(m.SetType),
		valueKey:               internalSerialization.DeserializeStringFunc(m.SetValue),
	}
}

func (m *Field) GetDeepLinkToSubObject() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Field, *string](m, deepLinkToSubObjectKey)
}

func (m *Field) SetDeepLinkToSubObject(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, deepLinkToSubObjectKey, value)
}

func (m *Field) GetDisplayAsTimeAgo() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Field, *string](m, displayAsTimeAgoKey)
}

func (m *Field) SetDisplayAsTimeAgo(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, displayAsTimeAgoKey, value)
}

func (m *Field) GetLabel() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Field, *string](m, labelKey)
}

func (m *Field) SetLabel(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, labelKey, value)
}

func (m *Field) GetShowLabel() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Field, *string](m, showLabelKey)
}

func (m *Field) SetShowLabel(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, showLabelKey, value)
}

func (m *Field) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Field, *string](m, typeKey)
}

func (m *Field) SetType(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, value)
}

func (m *Field) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Field, *string](m, valueKey)
}

func (m *Field) SetValue(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, value)
}
