package activitysubscriptionsapi

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

// Activity represents the activity.
type Activity struct {
	core.BaseModel
}

// NewActivity creates a new instance of Activity.
func NewActivity() *Activity {
	return &Activity{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateActivityFromDiscriminatorValue creates a new [Activity] from a [serialization.ParseNode].
func CreateActivityFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivity(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *Activity) Serialize(_ serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
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

// GetActivityTypeID returns the activity type id value.
func (m *Activity) GetActivityTypeID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, *string](m, activityTypeIDKey)
}

// SetActivityTypeID sets the activity type id value.
func (m *Activity) SetActivityTypeID(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, activityTypeIDKey, value)
}

// GetSourceTableName returns the source table name value.
func (m *Activity) GetSourceTableName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, *string](m, sourceTableNameKey)
}

// SetSourceTableName sets the source table name value.
func (m *Activity) SetSourceTableName(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sourceTableNameKey, value)
}

// GetSubObjectTableName returns the sub object table name value.
func (m *Activity) GetSubObjectTableName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, *string](m, subObjectTableNameKey)
}

// SetSubObjectTableName sets the sub object table name value.
func (m *Activity) SetSubObjectTableName(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, subObjectTableNameKey, value)
}

// GetSubObjectSysID returns the sub object sys id value.
func (m *Activity) GetSubObjectSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, *string](m, subObjectSysIDKey)
}

// SetSubObjectSysID sets the sub object sys id value.
func (m *Activity) SetSubObjectSysID(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, subObjectSysIDKey, value)
}

// GetTitle returns the title value.
func (m *Activity) GetTitle() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, *string](m, titleKey)
}

// SetTitle sets the title value.
func (m *Activity) SetTitle(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, titleKey, value)
}

// GetSysIDKey returns the sys id key value.
func (m *Activity) GetSysIDKey() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, *string](m, sysIDKey)
}

// SetSysIDKey sets the sys id key value.
func (m *Activity) SetSysIDKey(value *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, value)
}

// GetContentFields returns the content fields value.
func (m *Activity) GetContentFields() ([]*Field, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, []*Field](m, contentFieldsKey)
}

// SetContentFields sets the content fields value.
func (m *Activity) SetContentFields(value []*Field) error {
	return store.DefaultBackedModelMutatorFunc(m, contentFieldsKey, value)
}

// GetSubheaderFields returns the subheader fields value.
func (m *Activity) GetSubheaderFields() ([]*Field, error) {
	return store.DefaultBackedModelAccessorFunc[*Activity, []*Field](m, subheaderFieldsKey)
}

// SetSubheaderFields sets the subheader fields value.
func (m *Activity) SetSubheaderFields(value []*Field) error {
	return store.DefaultBackedModelMutatorFunc(m, subheaderFieldsKey, value)
}
