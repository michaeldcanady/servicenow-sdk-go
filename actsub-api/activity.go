package actsubapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
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
	newInternal.BaseModel
}

func NewActivity() *Activity {
	return &Activity{
		BaseModel: *newInternal.NewBaseModel(),
	}
}

func CreateActivityFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivity(), nil
}

func (m *Activity) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return nil
}

// activityStringFieldDeserializer builds a deserializer for a single string-typed
// field, delegating the actual mutation to the provided setter.
func activityStringFieldDeserializer(setter func(*string) error) func(serialization.ParseNode) error {
	return func(parseNode serialization.ParseNode) error {
		val, err := parseNode.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			return setter(val)
		}
		return nil
	}
}

// activityFieldCollectionDeserializer builds a deserializer for a `[]*Field`
// collection-typed field, delegating the actual mutation to the provided setter.
func activityFieldCollectionDeserializer(setter func([]*Field) error) func(serialization.ParseNode) error {
	return func(parseNode serialization.ParseNode) error {
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
			return setter(res)
		}
		return nil
	}
}

func (m *Activity) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		activityTypeIDKey:     activityStringFieldDeserializer(m.SetActivityTypeID),
		sourceTableNameKey:    activityStringFieldDeserializer(m.SetSourceTableName),
		subObjectTableNameKey: activityStringFieldDeserializer(m.SetSubObjectTableName),
		subObjectSysIDKey:     activityStringFieldDeserializer(m.SetSubObjectSysID),
		titleKey:              activityStringFieldDeserializer(m.SetTitle),
		sysIDKey:              activityStringFieldDeserializer(m.SetSysIDKey),
		contentFieldsKey:      activityFieldCollectionDeserializer(m.SetContentFields),
		subheaderFieldsKey:    activityFieldCollectionDeserializer(m.SetSubheaderFields),
	}
}

func (m *Activity) GetActivityTypeID() (*string, error) {
	if internal.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	val, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, activityTypeIDKey)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (m *Activity) SetActivityTypeID(value *string) error {
	if internal.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, activityTypeIDKey, value)
}

func (m *Activity) GetSourceTableName() (*string, error) {
	if internal.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sourceTableNameKey)
}

func (m *Activity) SetSourceTableName(value *string) error {
	if internal.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, sourceTableNameKey, value)
}

func (m *Activity) GetSubObjectTableName() (*string, error) {
	if internal.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, subObjectTableNameKey)
}

func (m *Activity) SetSubObjectTableName(value *string) error {
	if internal.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, subObjectTableNameKey, value)
}

func (m *Activity) GetSubObjectSysID() (*string, error) {
	if internal.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, subObjectSysIDKey)
}

func (m *Activity) SetSubObjectSysID(value *string) error {
	if internal.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, subObjectSysIDKey, value)
}

func (m *Activity) GetTitle() (*string, error) {
	if internal.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, titleKey)
}

func (m *Activity) SetTitle(value *string) error {
	if internal.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, titleKey, value)
}

func (m *Activity) GetSysIDKey() (*string, error) {
	if internal.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sysIDKey)
}

func (m *Activity) SetSysIDKey(value *string) error {
	if internal.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, sysIDKey, value)
}

func (m *Activity) GetContentFields() ([]*Field, error) {
	if internal.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []*Field](backingStore, contentFieldsKey)
}

func (m *Activity) SetContentFields(value []*Field) error {
	if internal.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, contentFieldsKey, value)
}

func (m *Activity) GetSubheaderFields() ([]*Field, error) {
	if internal.IsNil(m) {
		return nil, errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil, errors.New("store is nil")
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []*Field](backingStore, subheaderFieldsKey)
}

func (m *Activity) SetSubheaderFields(value []*Field) error {
	if internal.IsNil(m) {
		return errors.New("model is nil")
	}

	backingStore := m.GetBackingStore()
	if internal.IsNil(backingStore) {
		return errors.New("store is nil")
	}

	return store.DefaultBackedModelMutatorFunc(backingStore, subheaderFieldsKey, value)
}
