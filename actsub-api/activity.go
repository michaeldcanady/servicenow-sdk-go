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
	activityTypeIDKey     = "activity_type_id"
	contentFieldsKey      = "content_fields"
	sourceTableNameKey    = "source_table_name"
	subheaderFieldsKey    = "subheader_fields"
	subObjectSysIDKey     = "subobject_sys_id"
	subObjectTableNameKey = "subobject_table_name"
	sysIDKey              = "sys_id"
	titleKey              = "title"
)

type Activity struct {
	internal.BaseModel
}

func NewActivity() *Activity {
	return &Activity{
		BaseModel: *internal.NewBaseModel(),
	}
}

func CreateActivityFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivity(), nil
}

func (m *Activity) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return nil
}

func (m *Activity) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		activityTypeIDKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetActivityTypeID(val)
			}
			return nil
		},
		sourceTableNameKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetSourceTableName(val)
			}
			return nil
		},
		subObjectTableNameKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetSubObjectTableName(val)
			}
			return nil
		},
		subObjectSysIDKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetSubObjectSysID(val)
			}
			return nil
		},
		titleKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetTitle(val)
			}
			return nil
		},
		sysIDKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetStringValue()
			if err != nil {
				return err
			}
			if val != nil {
				return m.SetSysIDKey(val)
			}
			return nil
		},
		contentFieldsKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetCollectionOfObjectValues(CreateFieldFromDiscriminatorValue)
			if err != nil {
				return err
			}
			if val != nil {
				res := make([]*Field, len(val))
				for i, v := range val {
					if v != nil {
						res[i] = v.(*Field)
					}
				}
				return m.SetContentFields(res)
			}
			return nil
		},
		subheaderFieldsKey: func(parseNode serialization.ParseNode) error {
			val, err := parseNode.GetCollectionOfObjectValues(CreateFieldFromDiscriminatorValue)
			if err != nil {
				return err
			}
			if val != nil {
				res := make([]*Field, len(val))
				for i, v := range val {
					if v != nil {
						res[i] = v.(*Field)
					}
				}
				return m.SetSubheaderFields(res)
			}
			return nil
		},
	}
}

func (m *Activity) GetActivityTypeID() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	val, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, activityTypeIDKey)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (m *Activity) SetActivityTypeID(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, activityTypeIDKey, value)
}

func (m *Activity) GetSourceTableName() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sourceTableNameKey)
}

func (m *Activity) SetSourceTableName(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, sourceTableNameKey, value)
}

func (m *Activity) GetSubObjectTableName() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, subObjectTableNameKey)
}

func (m *Activity) SetSubObjectTableName(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, subObjectTableNameKey, value)
}

func (m *Activity) GetSubObjectSysID() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, subObjectSysIDKey)
}

func (m *Activity) SetSubObjectSysID(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, subObjectSysIDKey, value)
}

func (m *Activity) GetTitle() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, titleKey)
}

func (m *Activity) SetTitle(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, titleKey, value)
}

func (m *Activity) GetSysIDKey() (*string, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sysIDKey)
}

func (m *Activity) SetSysIDKey(value *string) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, sysIDKey, value)
}

func (m *Activity) GetContentFields() ([]*Field, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []*Field](backingStore, contentFieldsKey)
}

func (m *Activity) SetContentFields(value []*Field) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, contentFieldsKey, value)
}

func (m *Activity) GetSubheaderFields() ([]*Field, error) {
	if conversion.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []*Field](backingStore, subheaderFieldsKey)
}

func (m *Activity) SetSubheaderFields(value []*Field) error {
	if conversion.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if conversion.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, subheaderFieldsKey, value)
}
