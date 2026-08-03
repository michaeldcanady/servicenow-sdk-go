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

// ServiceConfigModel represents the service config model.
type ServiceConfigModel struct {
	core.BaseModel
}

// NewServiceConfig creates a new instance of ServiceConfigModel.
func NewServiceConfig() *ServiceConfigModel {
	return &ServiceConfigModel{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateServiceConfigFromDiscriminatorValue creates a new ServiceConfig from a ParseNode.
func CreateServiceConfigFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServiceConfig(), nil
}

// Serialize writes the objects properties to the current writer.
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

// GetFieldDeserializers returns the deserialization information for this object.
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

// GetActive returns the active value.
func (m *ServiceConfigModel) GetActive() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *bool](m, activeKey)
}

// SetActive sets the active value.
func (m *ServiceConfigModel) SetActive(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, activeKey, val)
}

// GetActiveString returns the active string value.
func (m *ServiceConfigModel) GetActiveString() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, activeStringKey)
}

// SetActiveString sets the active string value.
func (m *ServiceConfigModel) SetActiveString(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, activeStringKey, val)
}

// GetAppointmentBookingConfig returns the appointment booking config value.
func (m *ServiceConfigModel) GetAppointmentBookingConfig() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, appointmentBookingConfigKey)
}

// SetAppointmentBookingConfig sets the appointment booking config value.
func (m *ServiceConfigModel) SetAppointmentBookingConfig(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appointmentBookingConfigKey, val)
}

// GetAppointmentDuration returns the appointment duration value.
func (m *ServiceConfigModel) GetAppointmentDuration() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, appointmentDurationKey)
}

// SetAppointmentDuration sets the appointment duration value.
func (m *ServiceConfigModel) SetAppointmentDuration(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appointmentDurationKey, val)
}

// GetAppointmentsPerBookableSlot returns the appointments per bookable slot value.
func (m *ServiceConfigModel) GetAppointmentsPerBookableSlot() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, appointmentsPerBookableSlotKey)
}

// SetAppointmentsPerBookableSlot sets the appointments per bookable slot value.
func (m *ServiceConfigModel) SetAppointmentsPerBookableSlot(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appointmentsPerBookableSlotKey, val)
}

// GetBookableDays returns the bookable days value.
func (m *ServiceConfigModel) GetBookableDays() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, bookableDaysKey)
}

// SetBookableDays sets the bookable days value.
func (m *ServiceConfigModel) SetBookableDays(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, bookableDaysKey, val)
}

// GetCancelByTime returns the cancel by time value.
func (m *ServiceConfigModel) GetCancelByTime() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, cancelByTimeKey)
}

// SetCancelByTime sets the cancel by time value.
func (m *ServiceConfigModel) SetCancelByTime(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, cancelByTimeKey, val)
}

// GetDefaultTimezone returns the default timezone value.
func (m *ServiceConfigModel) GetDefaultTimezone() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, defaultTimezoneKey)
}

// SetDefaultTimezone sets the default timezone value.
func (m *ServiceConfigModel) SetDefaultTimezone(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, defaultTimezoneKey, val)
}

// GetEnableAdvancedConfig returns the enable advanced config value.
func (m *ServiceConfigModel) GetEnableAdvancedConfig() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *bool](m, enableAdvancedConfigKey)
}

// SetEnableAdvancedConfig sets the enable advanced config value.
func (m *ServiceConfigModel) SetEnableAdvancedConfig(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, enableAdvancedConfigKey, val)
}

// GetFieldMapping returns the field mapping value.
func (m *ServiceConfigModel) GetFieldMapping() (FieldMapping, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, FieldMapping](m, fieldMappingKey)
}

// SetFieldMapping sets the field mapping value.
func (m *ServiceConfigModel) SetFieldMapping(val FieldMapping) error {
	return store.DefaultBackedModelMutatorFunc(m, fieldMappingKey, val)
}

// GetFutureBookableMaxDays returns the future bookable max days value.
func (m *ServiceConfigModel) GetFutureBookableMaxDays() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, futureBookableMaxDaysKey)
}

// SetFutureBookableMaxDays sets the future bookable max days value.
func (m *ServiceConfigModel) SetFutureBookableMaxDays(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, futureBookableMaxDaysKey, val)
}

// GetLeadTime returns the lead time value.
func (m *ServiceConfigModel) GetLeadTime() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, leadTimeKey)
}

// SetLeadTime sets the lead time value.
func (m *ServiceConfigModel) SetLeadTime(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, leadTimeKey, val)
}

// GetMandatory returns the mandatory value.
func (m *ServiceConfigModel) GetMandatory() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, mandatoryKey)
}

// SetMandatory sets the mandatory value.
func (m *ServiceConfigModel) SetMandatory(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, mandatoryKey, val)
}

// GetUseSlotEndTimeAs returns the use slot end time as value.
func (m *ServiceConfigModel) GetUseSlotEndTimeAs() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, useSlotEndTimeAsKey)
}

// SetUseSlotEndTimeAs sets the use slot end time as value.
func (m *ServiceConfigModel) SetUseSlotEndTimeAs(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, useSlotEndTimeAsKey, val)
}

// GetWorkDuration returns the work duration value.
func (m *ServiceConfigModel) GetWorkDuration() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceConfigModel, *string](m, workDurationKey)
}

// SetWorkDuration sets the work duration value.
func (m *ServiceConfigModel) SetWorkDuration(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, workDurationKey, val)
}
