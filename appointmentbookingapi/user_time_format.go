package appointmentbookingapi // nolint:dupl // shares field-count shape with CalendarResponseModel by coincidence, not copy-paste; distinct API concept, not worth sacrificing named accessors for

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// UserTimeFormat represents userTimeFormat nested object.
type UserTimeFormat interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetType() (*TimeFormat, error)
	SetType(*TimeFormat) error
	GetValue() (*string, error)
	SetValue(*string) error
}

// UserTimeFormatModel represents the user time format model.
type UserTimeFormatModel struct {
	core.BaseModel
}

// NewUserTimeFormat creates a new instance of UserTimeFormatModel.
func NewUserTimeFormat() *UserTimeFormatModel {
	return &UserTimeFormatModel{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateUserTimeFormatFromDiscriminatorValue creates a new UserTimeFormat from a ParseNode.
func CreateUserTimeFormatFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUserTimeFormat(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *UserTimeFormatModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeEnumFunc(typeKey, m.GetType),
		internalSerialization.SerializeStringFunc(valueKey, m.GetValue),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *UserTimeFormatModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		typeKey:  internalSerialization.DeserializeEnumFunc(ParseTimeFormat, m.SetType),
		valueKey: internalSerialization.DeserializeStringFunc(m.SetValue),
	}
}

// GetType returns the type value.
func (m *UserTimeFormatModel) GetType() (*TimeFormat, error) {
	return store.DefaultBackedModelAccessorFunc[*UserTimeFormatModel, *TimeFormat](m, typeKey)
}

// SetType sets the type value.
func (m *UserTimeFormatModel) SetType(val *TimeFormat) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}

// GetValue returns the value value.
func (m *UserTimeFormatModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserTimeFormatModel, *string](m, valueKey)
}

// SetValue sets the value value.
func (m *UserTimeFormatModel) SetValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}
