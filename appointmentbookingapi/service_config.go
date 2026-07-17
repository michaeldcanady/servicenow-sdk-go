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
	setActive(*bool) error
	GetActiveString() (*string, error)
	setActiveString(*string) error
	GetAppointmentBookingConfig() (*string, error)
	setAppointmentBookingConfig(*string) error
	GetAppointmentDuration() (*string, error)
	setAppointmentDuration(*string) error
	GetAppointmentsPerBookableSlot() (*string, error)
	setAppointmentsPerBookableSlot(*string) error
	GetBookableDays() (*string, error)
	setBookableDays(*string) error
	GetCancelByTime() (*string, error)
	setCancelByTime(*string) error
	GetDefaultTimezone() (*string, error)
	setDefaultTimezone(*string) error
	GetEnableAdvancedConfig() (*bool, error)
	setEnableAdvancedConfig(*bool) error
	GetFieldMapping() (FieldMapping, error)
	setFieldMapping(FieldMapping) error
	GetFutureBookableMaxDays() (*string, error)
	setFutureBookableMaxDays(*string) error
	GetLeadTime() (*string, error)
	setLeadTime(*string) error
	GetMandatory() (*string, error)
	setMandatory(*string) error
	GetUseSlotEndTimeAs() (*string, error)
	setUseSlotEndTimeAs(*string) error
	GetWorkDuration() (*string, error)
	setWorkDuration(*string) error
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
		activeKey:                      internalSerialization.DeserializeBoolFunc(m.setActive),
		activeStringKey:                internalSerialization.DeserializeStringFunc(m.setActiveString),
		appointmentBookingConfigKey:    internalSerialization.DeserializeStringFunc(m.setAppointmentBookingConfig),
		appointmentDurationKey:         internalSerialization.DeserializeStringFunc(m.setAppointmentDuration),
		appointmentsPerBookableSlotKey: internalSerialization.DeserializeStringFunc(m.setAppointmentsPerBookableSlot),
		bookableDaysKey:                internalSerialization.DeserializeStringFunc(m.setBookableDays),
		cancelByTimeKey:                internalSerialization.DeserializeStringFunc(m.setCancelByTime),
		defaultTimezoneKey:             internalSerialization.DeserializeStringFunc(m.setDefaultTimezone),
		enableAdvancedConfigKey:        internalSerialization.DeserializeBoolFunc(m.setEnableAdvancedConfig),
		fieldMappingKey:                internalSerialization.DeserializeObjectValueFunc[FieldMapping](CreateFieldMappingFromDiscriminatorValue, m.setFieldMapping),
		futureBookableMaxDaysKey:       internalSerialization.DeserializeStringFunc(m.setFutureBookableMaxDays),
		leadTimeKey:                    internalSerialization.DeserializeStringFunc(m.setLeadTime),
		mandatoryKey:                   internalSerialization.DeserializeStringFunc(m.setMandatory),
		useSlotEndTimeAsKey:            internalSerialization.DeserializeStringFunc(m.setUseSlotEndTimeAs),
		workDurationKey:                internalSerialization.DeserializeStringFunc(m.setWorkDuration),
	}
}

func (m *ServiceConfigModel) GetActive() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *bool](m, activeKey)
}
func (m *ServiceConfigModel) setActive(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, activeKey, val)
}
func (m *ServiceConfigModel) GetActiveString() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, activeStringKey)
}
func (m *ServiceConfigModel) setActiveString(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, activeStringKey, val)
}
func (m *ServiceConfigModel) GetAppointmentBookingConfig() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, appointmentBookingConfigKey)
}
func (m *ServiceConfigModel) setAppointmentBookingConfig(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appointmentBookingConfigKey, val)
}
func (m *ServiceConfigModel) GetAppointmentDuration() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, appointmentDurationKey)
}
func (m *ServiceConfigModel) setAppointmentDuration(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appointmentDurationKey, val)
}
func (m *ServiceConfigModel) GetAppointmentsPerBookableSlot() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, appointmentsPerBookableSlotKey)
}
func (m *ServiceConfigModel) setAppointmentsPerBookableSlot(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appointmentsPerBookableSlotKey, val)
}
func (m *ServiceConfigModel) GetBookableDays() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, bookableDaysKey)
}
func (m *ServiceConfigModel) setBookableDays(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, bookableDaysKey, val)
}
func (m *ServiceConfigModel) GetCancelByTime() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, cancelByTimeKey)
}
func (m *ServiceConfigModel) setCancelByTime(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, cancelByTimeKey, val)
}
func (m *ServiceConfigModel) GetDefaultTimezone() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, defaultTimezoneKey)
}
func (m *ServiceConfigModel) setDefaultTimezone(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, defaultTimezoneKey, val)
}
func (m *ServiceConfigModel) GetEnableAdvancedConfig() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *bool](m, enableAdvancedConfigKey)
}
func (m *ServiceConfigModel) setEnableAdvancedConfig(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, enableAdvancedConfigKey, val)
}
func (m *ServiceConfigModel) GetFieldMapping() (FieldMapping, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, FieldMapping](m, fieldMappingKey)
}
func (m *ServiceConfigModel) setFieldMapping(val FieldMapping) error {
	return store.DefaultBackedModelMutatorFunc(m, fieldMappingKey, val)
}
func (m *ServiceConfigModel) GetFutureBookableMaxDays() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, futureBookableMaxDaysKey)
}
func (m *ServiceConfigModel) setFutureBookableMaxDays(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, futureBookableMaxDaysKey, val)
}
func (m *ServiceConfigModel) GetLeadTime() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, leadTimeKey)
}
func (m *ServiceConfigModel) setLeadTime(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, leadTimeKey, val)
}
func (m *ServiceConfigModel) GetMandatory() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, mandatoryKey)
}
func (m *ServiceConfigModel) setMandatory(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, mandatoryKey, val)
}
func (m *ServiceConfigModel) GetUseSlotEndTimeAs() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, useSlotEndTimeAsKey)
}
func (m *ServiceConfigModel) setUseSlotEndTimeAs(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, useSlotEndTimeAsKey, val)
}
func (m *ServiceConfigModel) GetWorkDuration() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, workDurationKey)
}
func (m *ServiceConfigModel) setWorkDuration(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, workDurationKey, val)
}
