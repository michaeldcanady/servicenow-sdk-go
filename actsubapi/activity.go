package actsubapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
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
	core.BaseModel
}

func NewActivity() *Activity {
	return &Activity{
		BaseModel: *core.NewBaseModel(),
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
		activityTypeIDKey:     internalSerialization.DeserializeStringFunc(m.SetActivityTypeID),
		sourceTableNameKey:    internalSerialization.DeserializeStringFunc(m.SetSourceTableName),
		subObjectTableNameKey: internalSerialization.DeserializeStringFunc(m.SetSubObjectTableName),
		subObjectSysIDKey:     internalSerialization.DeserializeStringFunc(m.SetSubObjectSysID),
		titleKey:              internalSerialization.DeserializeStringFunc(m.SetTitle),
		sysIDKey:              internalSerialization.DeserializeStringFunc(m.SetSysIDKey),
		contentFieldsKey:      internalSerialization.DeserializeCollectionOfObjectValuesFunc[*Field](CreateFieldFromDiscriminatorValue, m.SetContentFields),
		subheaderFieldsKey:    internalSerialization.DeserializeCollectionOfObjectValuesFunc[*Field](CreateFieldFromDiscriminatorValue, m.SetSubheaderFields),
	}
}

func (m *Activity) GetActivityTypeID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, *string](m, activityTypeIDKey)
}

func (m *Activity) SetActivityTypeID(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, activityTypeIDKey, value)
}

func (m *Activity) GetSourceTableName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, *string](m, sourceTableNameKey)
}

func (m *Activity) SetSourceTableName(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sourceTableNameKey, value)
}

func (m *Activity) GetSubObjectTableName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, *string](m, subObjectTableNameKey)
}

func (m *Activity) SetSubObjectTableName(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, subObjectTableNameKey, value)
}

func (m *Activity) GetSubObjectSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, *string](m, subObjectSysIDKey)
}

func (m *Activity) SetSubObjectSysID(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, subObjectSysIDKey, value)
}

func (m *Activity) GetTitle() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, *string](m, titleKey)
}

func (m *Activity) SetTitle(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, titleKey, value)
}

func (m *Activity) GetSysIDKey() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, *string](m, sysIDKey)
}

func (m *Activity) SetSysIDKey(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, value)
}

func (m *Activity) GetContentFields() ([]*Field, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, []*Field](m, contentFieldsKey)
}

func (m *Activity) SetContentFields(value []*Field) error {
	return store.DefaultBackedModelMutatorFunc(m, contentFieldsKey, value)
}

func (m *Activity) GetSubheaderFields() ([]*Field, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, []*Field](m, subheaderFieldsKey)
}

func (m *Activity) SetSubheaderFields(value []*Field) error {
	return store.DefaultBackedModelMutatorFunc(m, subheaderFieldsKey, value)
}
