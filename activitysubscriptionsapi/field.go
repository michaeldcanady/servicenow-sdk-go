package activitysubscriptionsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
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

// Field represents the field.
type Field struct {
	core.BaseModel
}

// NewField creates a new instance of Field.
func NewField() *Field {
	return &Field{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateFieldFromDiscriminatorValue creates a new [Field] from a [serialization.ParseNode].
func CreateFieldFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewField(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *Field) Serialize(_ serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
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

// GetDeepLinkToSubObject returns the deep link to sub object value.
func (m *Field) GetDeepLinkToSubObject() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Field, *string](m, deepLinkToSubObjectKey)
}

// SetDeepLinkToSubObject sets the deep link to sub object value.
func (m *Field) SetDeepLinkToSubObject(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, deepLinkToSubObjectKey, value)
}

// GetDisplayAsTimeAgo returns the display as time ago value.
func (m *Field) GetDisplayAsTimeAgo() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Field, *string](m, displayAsTimeAgoKey)
}

// SetDisplayAsTimeAgo sets the display as time ago value.
func (m *Field) SetDisplayAsTimeAgo(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, displayAsTimeAgoKey, value)
}

// GetLabel returns the label value.
func (m *Field) GetLabel() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Field, *string](m, labelKey)
}

// SetLabel sets the label value.
func (m *Field) SetLabel(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, labelKey, value)
}

// GetShowLabel returns the show label value.
func (m *Field) GetShowLabel() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Field, *string](m, showLabelKey)
}

// SetShowLabel sets the show label value.
func (m *Field) SetShowLabel(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, showLabelKey, value)
}

// GetType returns the type value.
func (m *Field) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Field, *string](m, typeKey)
}

// SetType sets the type value.
func (m *Field) SetType(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, value)
}

// GetValue returns the value value.
func (m *Field) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Field, *string](m, valueKey)
}

// SetValue sets the value value.
func (m *Field) SetValue(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, value)
}
