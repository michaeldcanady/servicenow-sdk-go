package actsubapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
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
	internal.BaseModel
}

func NewField() *Field {
	return &Field{
		BaseModel: *internal.NewBaseModel(),
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
		deepLinkToSubObjectKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetDeepLinkToSubObject(val)
			}
			return nil
		},
		displayAsTimeAgoKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetDisplayAsTimeAgo(val)
			}
			return nil
		},
		labelKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetLabel(val)
			}
			return nil
		},
		showLabelKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetShowLabel(val)
			}
			return nil
		},
		typeKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetType(val)
			}
			return nil
		},
		valueKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetValue(val)
			}
			return nil
		},
	}
}

func (m *Field) GetDeepLinkToSubObject() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, deepLinkToSubObjectKey)
}

func (m *Field) SetDeepLinkToSubObject(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), deepLinkToSubObjectKey, value)
}

func (m *Field) GetDisplayAsTimeAgo() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, displayAsTimeAgoKey)
}

func (m *Field) SetDisplayAsTimeAgo(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), displayAsTimeAgoKey, value)
}

func (m *Field) GetLabel() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, labelKey)
}

func (m *Field) SetLabel(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), labelKey, value)
}

func (m *Field) GetShowLabel() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, showLabelKey)
}

func (m *Field) SetShowLabel(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), showLabelKey, value)
}

func (m *Field) GetType() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, typeKey)
}

func (m *Field) SetType(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), typeKey, value)
}

func (m *Field) GetValue() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, valueKey)
}

func (m *Field) SetValue(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), valueKey, value)
}
