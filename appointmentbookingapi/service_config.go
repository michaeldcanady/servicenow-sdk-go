package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// ServiceConfig represents the service_config nested object.
type ServiceConfig interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetActive() (*bool, error)
	SetActive(*bool) error
	GetActiveString() (*string, error)
	SetActiveString(*string) error
	GetAppointmentBookingConfig() (*string, error)
	SetAppointmentBookingConfig(*string) error
	GetAppointmentDuration() (*string, error)
	SetAppointmentDuration(*string) error
	GetAppointmentsPerBookableSlot() (*string, error)
	SetAppointmentsPerBookableSlot(*string) error
	GetBookableDays() (*string, error)
	SetBookableDays(*string) error
	GetCancelByTime() (*string, error)
	SetCancelByTime(*string) error
	GetDefaultTimezone() (*string, error)
	SetDefaultTimezone(*string) error
	GetEnableAdvancedConfig() (*bool, error)
	SetEnableAdvancedConfig(*bool) error
	GetFieldMapping() (FieldMapping, error)
	SetFieldMapping(FieldMapping) error
	GetFutureBookableMaxDays() (*string, error)
	SetFutureBookableMaxDays(*string) error
	GetLeadTime() (*string, error)
	SetLeadTime(*string) error
	GetMandatory() (*string, error)
	SetMandatory(*string) error
	GetUseSlotEndTimeAs() (*string, error)
	SetUseSlotEndTimeAs(*string) error
	GetWorkDuration() (*string, error)
	SetWorkDuration(*string) error
}

type ServiceConfigModel struct {
	core.BaseModel
}

func NewServiceConfig() *ServiceConfigModel {
	return &ServiceConfigModel{
		BaseModel: *core.NewBaseModel(),
	}
}

func CreateServiceConfigFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServiceConfig(), nil
}

func (m *ServiceConfigModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeBoolFunc(activeKey, m.GetActive),
		internalSerialization.SerializeStringFunc(activeStringKey, m.GetActiveString),
		internalSerialization.SerializeStringFunc(appointmentBookingConfigKey, m.GetAppointmentBookingConfig),
		internalSerialization.SerializeStringFunc(appointmentDurationKey, m.GetAppointmentDuration),
		internalSerialization.SerializeStringFunc(appointmentsPerBookableSlotKey, m.GetAppointmentsPerBookableSlot),
		internalSerialization.SerializeStringFunc(bookableDaysKey, m.GetBookableDays),
		internalSerialization.SerializeStringFunc(cancelByTimeKey, m.GetCancelByTime),
		internalSerialization.SerializeStringFunc(defaultTimezoneKey, m.GetDefaultTimezone),
		internalSerialization.SerializeBoolFunc(enableAdvancedConfigKey, m.GetEnableAdvancedConfig),
		internalSerialization.SerializeObjectValueFunc[FieldMapping](fieldMappingKey, m.GetFieldMapping),
		internalSerialization.SerializeStringFunc(futureBookableMaxDaysKey, m.GetFutureBookableMaxDays),
		internalSerialization.SerializeStringFunc(leadTimeKey, m.GetLeadTime),
		internalSerialization.SerializeStringFunc(mandatoryKey, m.GetMandatory),
		internalSerialization.SerializeStringFunc(useSlotEndTimeAsKey, m.GetUseSlotEndTimeAs),
		internalSerialization.SerializeStringFunc(workDurationKey, m.GetWorkDuration),
	)
}

func (m *ServiceConfigModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		activeKey:                      internalSerialization.DeserializeBoolFunc(m.SetActive),
		activeStringKey:                internalSerialization.DeserializeStringFunc(m.SetActiveString),
		appointmentBookingConfigKey:    internalSerialization.DeserializeStringFunc(m.SetAppointmentBookingConfig),
		appointmentDurationKey:         internalSerialization.DeserializeStringFunc(m.SetAppointmentDuration),
		appointmentsPerBookableSlotKey: internalSerialization.DeserializeStringFunc(m.SetAppointmentsPerBookableSlot),
		bookableDaysKey:                internalSerialization.DeserializeStringFunc(m.SetBookableDays),
		cancelByTimeKey:                internalSerialization.DeserializeStringFunc(m.SetCancelByTime),
		defaultTimezoneKey:             internalSerialization.DeserializeStringFunc(m.SetDefaultTimezone),
		enableAdvancedConfigKey:        internalSerialization.DeserializeBoolFunc(m.SetEnableAdvancedConfig),
		fieldMappingKey:                internalSerialization.DeserializeObjectValueFunc(CreateFieldMappingFromDiscriminatorValue, m.SetFieldMapping),
		futureBookableMaxDaysKey:       internalSerialization.DeserializeStringFunc(m.SetFutureBookableMaxDays),
		leadTimeKey:                    internalSerialization.DeserializeStringFunc(m.SetLeadTime),
		mandatoryKey:                   internalSerialization.DeserializeStringFunc(m.SetMandatory),
		useSlotEndTimeAsKey:            internalSerialization.DeserializeStringFunc(m.SetUseSlotEndTimeAs),
		workDurationKey:                internalSerialization.DeserializeStringFunc(m.SetWorkDuration),
	}
}

func (m *ServiceConfigModel) GetActive() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *bool](m, activeKey)
}
func (m *ServiceConfigModel) SetActive(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, activeKey, val)
}
func (m *ServiceConfigModel) GetActiveString() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, activeStringKey)
}
func (m *ServiceConfigModel) SetActiveString(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, activeStringKey, val)
}
func (m *ServiceConfigModel) GetAppointmentBookingConfig() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, appointmentBookingConfigKey)
}
func (m *ServiceConfigModel) SetAppointmentBookingConfig(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appointmentBookingConfigKey, val)
}
func (m *ServiceConfigModel) GetAppointmentDuration() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, appointmentDurationKey)
}
func (m *ServiceConfigModel) SetAppointmentDuration(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appointmentDurationKey, val)
}
func (m *ServiceConfigModel) GetAppointmentsPerBookableSlot() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, appointmentsPerBookableSlotKey)
}
func (m *ServiceConfigModel) SetAppointmentsPerBookableSlot(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appointmentsPerBookableSlotKey, val)
}
func (m *ServiceConfigModel) GetBookableDays() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, bookableDaysKey)
}
func (m *ServiceConfigModel) SetBookableDays(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, bookableDaysKey, val)
}
func (m *ServiceConfigModel) GetCancelByTime() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, cancelByTimeKey)
}
func (m *ServiceConfigModel) SetCancelByTime(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, cancelByTimeKey, val)
}
func (m *ServiceConfigModel) GetDefaultTimezone() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, defaultTimezoneKey)
}
func (m *ServiceConfigModel) SetDefaultTimezone(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, defaultTimezoneKey, val)
}
func (m *ServiceConfigModel) GetEnableAdvancedConfig() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *bool](m, enableAdvancedConfigKey)
}
func (m *ServiceConfigModel) SetEnableAdvancedConfig(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, enableAdvancedConfigKey, val)
}
func (m *ServiceConfigModel) GetFieldMapping() (FieldMapping, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, FieldMapping](m, fieldMappingKey)
}
func (m *ServiceConfigModel) SetFieldMapping(val FieldMapping) error {
	return store.DefaultBackedModelMutatorFunc(m, fieldMappingKey, val)
}
func (m *ServiceConfigModel) GetFutureBookableMaxDays() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, futureBookableMaxDaysKey)
}
func (m *ServiceConfigModel) SetFutureBookableMaxDays(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, futureBookableMaxDaysKey, val)
}
func (m *ServiceConfigModel) GetLeadTime() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, leadTimeKey)
}
func (m *ServiceConfigModel) SetLeadTime(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, leadTimeKey, val)
}
func (m *ServiceConfigModel) GetMandatory() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, mandatoryKey)
}
func (m *ServiceConfigModel) SetMandatory(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, mandatoryKey, val)
}
func (m *ServiceConfigModel) GetUseSlotEndTimeAs() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, useSlotEndTimeAsKey)
}
func (m *ServiceConfigModel) SetUseSlotEndTimeAs(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, useSlotEndTimeAsKey, val)
}
func (m *ServiceConfigModel) GetWorkDuration() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, workDurationKey)
}
func (m *ServiceConfigModel) SetWorkDuration(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, workDurationKey, val)
}
