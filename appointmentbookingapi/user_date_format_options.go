package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// UserDateFormatOptions represents userDateFormatOptions nested object.
type UserDateFormatOptions interface {
	serialization.Parsable
	kiotaStore.BackedModel

	// TODO: string representation of Numeric (Values or 1-31)
	GetDay() (*string, error)
	SetDay(*string) error
	GetMonth() (*ShortMonth, error)
	SetMonth(*ShortMonth) error
	// TODO: string representation of Numeric (Values of 1-5)
	GetWeek() (*string, error)
	SetWeek(*string) error
	GetWeekday() (*ShortWeekday, error)
	SetWeekday(*ShortWeekday) error
}

// UserDateFormatOptionsModel represents the user date format options model.
type UserDateFormatOptionsModel struct {
	core.BaseModel
}

// NewUserDateFormatOptions creates a new instance of UserDateFormatOptionsModel.
func NewUserDateFormatOptions() *UserDateFormatOptionsModel {
	return &UserDateFormatOptionsModel{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateUserDateFormatOptionsFromDiscriminatorValue creates a new UserDateFormatOptions from a ParseNode.
func CreateUserDateFormatOptionsFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUserDateFormatOptions(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *UserDateFormatOptionsModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(dayKey, m.GetDay),
		internalSerialization.SerializeEnumFunc(monthKey, m.GetMonth),
		internalSerialization.SerializeStringFunc(weekKey, m.GetWeek),
		internalSerialization.SerializeEnumFunc(weekdayKey, m.GetWeekday),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *UserDateFormatOptionsModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		dayKey:     internalSerialization.DeserializeStringFunc(m.SetDay),
		monthKey:   internalSerialization.DeserializeEnumFunc(ParseShortMonth, m.SetMonth),
		weekKey:    internalSerialization.DeserializeStringFunc(m.SetWeek),
		weekdayKey: internalSerialization.DeserializeEnumFunc(ParseShortWeekday, m.SetWeekday),
	}
}

// GetDay returns the day value.
func (m *UserDateFormatOptionsModel) GetDay() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserDateFormatOptionsModel, *string](m, dayKey)
}

// SetDay sets the day value.
func (m *UserDateFormatOptionsModel) SetDay(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, dayKey, val)
}

// GetMonth returns the month value.
func (m *UserDateFormatOptionsModel) GetMonth() (*ShortMonth, error) {
	return store.DefaultBackedModelAccessorFunc[*UserDateFormatOptionsModel, *ShortMonth](m, monthKey)
}

// SetMonth sets the month value.
func (m *UserDateFormatOptionsModel) SetMonth(val *ShortMonth) error {
	return store.DefaultBackedModelMutatorFunc(m, monthKey, val)
}

// GetWeek returns the week value.
func (m *UserDateFormatOptionsModel) GetWeek() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserDateFormatOptionsModel, *string](m, weekKey)
}

// SetWeek sets the week value.
func (m *UserDateFormatOptionsModel) SetWeek(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, weekKey, val)
}

// GetWeekday returns the weekday value.
func (m *UserDateFormatOptionsModel) GetWeekday() (*ShortWeekday, error) {
	return store.DefaultBackedModelAccessorFunc[*UserDateFormatOptionsModel, *ShortWeekday](m, weekdayKey)
}

// SetWeekday sets the weekday value.
func (m *UserDateFormatOptionsModel) SetWeekday(val *ShortWeekday) error {
	return store.DefaultBackedModelMutatorFunc(m, weekdayKey, val)
}
