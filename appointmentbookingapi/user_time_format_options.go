package appointmentbookingapi // nolint:dupl // shares field-count shape with RpVariableModel by coincidence, not copy-paste; distinct API concept, not worth sacrificing named accessors for

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// UserTimeFormatOptions represents userTimeFormatOptions nested object.
type UserTimeFormatOptions interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetHour() (*string, error)
	SetHour(*string) error
	GetHourCycle() (*string, error)
	SetHourCycle(*string) error
	GetMinute() (*string, error)
	SetMinute(*string) error
}

type UserTimeFormatOptionsModel struct {
	core.BaseModel
}

func NewUserTimeFormatOptions() *UserTimeFormatOptionsModel {
	return &UserTimeFormatOptionsModel{
		BaseModel: *core.NewBaseModel(),
	}
}

func CreateUserTimeFormatOptionsFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUserTimeFormatOptions(), nil
}

func (m *UserTimeFormatOptionsModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(hourKey, m.GetHour),
		internalSerialization.SerializeStringFunc(hourCycleKey, m.GetHourCycle),
		internalSerialization.SerializeStringFunc(minuteKey, m.GetMinute),
	)
}

func (m *UserTimeFormatOptionsModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		hourKey:      internalSerialization.DeserializeStringFunc(m.SetHour),
		hourCycleKey: internalSerialization.DeserializeStringFunc(m.SetHourCycle),
		minuteKey:    internalSerialization.DeserializeStringFunc(m.SetMinute),
	}
}

func (m *UserTimeFormatOptionsModel) GetHour() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserTimeFormatOptionsModel, *string](m, hourKey)
}
func (m *UserTimeFormatOptionsModel) SetHour(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, hourKey, val)
}
func (m *UserTimeFormatOptionsModel) GetHourCycle() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserTimeFormatOptionsModel, *string](m, hourCycleKey)
}
func (m *UserTimeFormatOptionsModel) SetHourCycle(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, hourCycleKey, val)
}
func (m *UserTimeFormatOptionsModel) GetMinute() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserTimeFormatOptionsModel, *string](m, minuteKey)
}
func (m *UserTimeFormatOptionsModel) SetMinute(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, minuteKey, val)
}
