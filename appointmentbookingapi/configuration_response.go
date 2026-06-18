package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// ConfigurationResponse represents the configuration response.
type ConfigurationResponse = core.ServiceNowItemResponse[*ConfigurationResultModel]

// CreateConfigurationResponseFromDiscriminatorValue is a factory for creating a ConfigurationResponse.
func CreateConfigurationResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*ConfigurationResultModel](CreateConfigurationResultFromDiscriminatorValue), nil
}

// ConfigurationResponseModel is the implementation of ConfigurationResponse.
type ConfigurationResponseModel = core.BaseServiceNowItemResponse[*ConfigurationResultModel]

// ConfigurationResult represents the result object in configuration response.
type ConfigurationResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetActive() (*bool, error)
	setActive(*bool) error
	GetActiveString() (*string, error)
	setActiveString(*string) error
	GetAdvancedCalendarViewPortal() (*bool, error)
	setAdvancedCalendarViewPortal(*bool) error
	GetAutoAcceptance() (*bool, error)
	setAutoAcceptance(*bool) error
	GetLocaleLanguage() (*string, error)
	setLocaleLanguage(*string) error
	GetServiceConfig() (ServiceConfig, error)
	setServiceConfig(ServiceConfig) error
	GetTaskTable() (*string, error)
	setTaskTable(*string) error
	GetTranslations() (map[string]interface{}, error)
	setTranslations(map[string]interface{}) error
	GetUserDateFormatOptions() (UserDateFormatOptions, error)
	setUserDateFormatOptions(UserDateFormatOptions) error
	GetUseRR() (*bool, error)
	setUseRR(*bool) error
	GetUserTimeFormat() (UserTimeFormat, error)
	setUserTimeFormat(UserTimeFormat) error
	GetUserTimeFormatOptions() (UserTimeFormatOptions, error)
	setUserTimeFormatOptions(UserTimeFormatOptions) error
	GetViewScale() (*string, error)
	setViewScale(*string) error
}

// ConfigurationResultModel implementation of ConfigurationResult
type ConfigurationResultModel struct {
	core.BaseModel
}

// NewConfigurationResult creates a new instance of ConfigurationResultModel
func NewConfigurationResult() *ConfigurationResultModel {
	return &ConfigurationResultModel{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateConfigurationResultFromDiscriminatorValue is a factory for creating a ConfigurationResult model.
func CreateConfigurationResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewConfigurationResult(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *ConfigurationResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeBoolFunc(activeKey)(m.GetActive),
		internalSerialization.SerializeStringFunc(activeStringKey)(m.GetActiveString),
		internalSerialization.SerializeBoolFunc(advancedCalendarViewPortalKey)(m.GetAdvancedCalendarViewPortal),
		internalSerialization.SerializeBoolFunc(autoAcceptanceKey)(m.GetAutoAcceptance),
		internalSerialization.SerializeStringFunc(localeLanguageKey)(m.GetLocaleLanguage),
		internalSerialization.SerializeObjectValueFunc[ServiceConfig](serviceConfigKey)(m.GetServiceConfig),
		internalSerialization.SerializeStringFunc(taskTableKey)(m.GetTaskTable),
		internalSerialization.SerializeAnyFunc(translationsKey)(m.GetTranslations),
		internalSerialization.SerializeObjectValueFunc[UserDateFormatOptions](userDateFormatOptionsKey)(m.GetUserDateFormatOptions),
		internalSerialization.SerializeBoolFunc(useRRKey)(m.GetUseRR),
		internalSerialization.SerializeObjectValueFunc[UserTimeFormat](userTimeFormatKey)(m.GetUserTimeFormat),
		internalSerialization.SerializeObjectValueFunc[UserTimeFormatOptions](userTimeFormatOptionsKey)(m.GetUserTimeFormatOptions),
		internalSerialization.SerializeStringFunc(viewScaleKey)(m.GetViewScale),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ConfigurationResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		activeKey:                     internalSerialization.DeserializeBoolFunc()(m.setActive),
		activeStringKey:               internalSerialization.DeserializeStringFunc()(m.setActiveString),
		advancedCalendarViewPortalKey: internalSerialization.DeserializeBoolFunc()(m.setAdvancedCalendarViewPortal),
		autoAcceptanceKey:             internalSerialization.DeserializeBoolFunc()(m.setAutoAcceptance),
		localeLanguageKey:             internalSerialization.DeserializeStringFunc()(m.setLocaleLanguage),
		serviceConfigKey:              internalSerialization.DeserializeObjectValueFunc[ServiceConfig](CreateServiceConfigFromDiscriminatorValue)(m.setServiceConfig),
		taskTableKey:                  internalSerialization.DeserializeStringFunc()(m.setTaskTable),
		translationsKey:               internalSerialization.DeserializeAnyFunc()(m.setTranslations),
		userDateFormatOptionsKey:      internalSerialization.DeserializeObjectValueFunc[UserDateFormatOptions](CreateUserDateFormatOptionsFromDiscriminatorValue)(m.setUserDateFormatOptions),
		useRRKey:                      internalSerialization.DeserializeBoolFunc()(m.setUseRR),
		userTimeFormatKey:             internalSerialization.DeserializeObjectValueFunc[UserTimeFormat](CreateUserTimeFormatFromDiscriminatorValue)(m.setUserTimeFormat),
		userTimeFormatOptionsKey:      internalSerialization.DeserializeObjectValueFunc[UserTimeFormatOptions](CreateUserTimeFormatOptionsFromDiscriminatorValue)(m.setUserTimeFormatOptions),
		viewScaleKey:                  internalSerialization.DeserializeStringFunc()(m.setViewScale),
	}
}

func (m *ConfigurationResultModel) GetActive() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), activeKey)
}
func (m *ConfigurationResultModel) setActive(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), activeKey, val)
}
func (m *ConfigurationResultModel) GetActiveString() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), activeStringKey)
}
func (m *ConfigurationResultModel) setActiveString(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), activeStringKey, val)
}
func (m *ConfigurationResultModel) GetAdvancedCalendarViewPortal() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), advancedCalendarViewPortalKey)
}
func (m *ConfigurationResultModel) setAdvancedCalendarViewPortal(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), advancedCalendarViewPortalKey, val)
}
func (m *ConfigurationResultModel) GetAutoAcceptance() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), autoAcceptanceKey)
}
func (m *ConfigurationResultModel) setAutoAcceptance(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), autoAcceptanceKey, val)
}
func (m *ConfigurationResultModel) GetLocaleLanguage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), localeLanguageKey)
}
func (m *ConfigurationResultModel) setLocaleLanguage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), localeLanguageKey, val)
}
func (m *ConfigurationResultModel) GetServiceConfig() (ServiceConfig, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, ServiceConfig](m.GetBackingStore(), serviceConfigKey)
}
func (m *ConfigurationResultModel) setServiceConfig(val ServiceConfig) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), serviceConfigKey, val)
}
func (m *ConfigurationResultModel) GetTaskTable() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), taskTableKey)
}
func (m *ConfigurationResultModel) setTaskTable(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), taskTableKey, val)
}
func (m *ConfigurationResultModel) GetTranslations() (any, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, any](m.GetBackingStore(), translationsKey)
}
func (m *ConfigurationResultModel) setTranslations(val any) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), translationsKey, val)
}
func (m *ConfigurationResultModel) GetUserDateFormatOptions() (UserDateFormatOptions, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, UserDateFormatOptions](m.GetBackingStore(), userDateFormatOptionsKey)
}
func (m *ConfigurationResultModel) setUserDateFormatOptions(val UserDateFormatOptions) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), userDateFormatOptionsKey, val)
}
func (m *ConfigurationResultModel) GetUseRR() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), useRRKey)
}
func (m *ConfigurationResultModel) setUseRR(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), useRRKey, val)
}
func (m *ConfigurationResultModel) GetUserTimeFormat() (UserTimeFormat, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, UserTimeFormat](m.GetBackingStore(), userTimeFormatKey)
}
func (m *ConfigurationResultModel) setUserTimeFormat(val UserTimeFormat) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), userTimeFormatKey, val)
}
func (m *ConfigurationResultModel) GetUserTimeFormatOptions() (UserTimeFormatOptions, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, UserTimeFormatOptions](m.GetBackingStore(), userTimeFormatOptionsKey)
}
func (m *ConfigurationResultModel) setUserTimeFormatOptions(val UserTimeFormatOptions) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), userTimeFormatOptionsKey, val)
}
func (m *ConfigurationResultModel) GetViewScale() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), viewScaleKey)
}
func (m *ConfigurationResultModel) setViewScale(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), viewScaleKey, val)
}

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
		internalSerialization.SerializeBoolFunc(activeKey)(m.GetActive),
		internalSerialization.SerializeStringFunc(activeStringKey)(m.GetActiveString),
		internalSerialization.SerializeStringFunc(appointmentBookingConfigKey)(m.GetAppointmentBookingConfig),
		internalSerialization.SerializeStringFunc(appointmentDurationKey)(m.GetAppointmentDuration),
		internalSerialization.SerializeStringFunc(appointmentsPerBookableSlotKey)(m.GetAppointmentsPerBookableSlot),
		internalSerialization.SerializeStringFunc(bookableDaysKey)(m.GetBookableDays),
		internalSerialization.SerializeStringFunc(cancelByTimeKey)(m.GetCancelByTime),
		internalSerialization.SerializeStringFunc(defaultTimezoneKey)(m.GetDefaultTimezone),
		internalSerialization.SerializeBoolFunc(enableAdvancedConfigKey)(m.GetEnableAdvancedConfig),
		internalSerialization.SerializeObjectValueFunc[FieldMapping](fieldMappingKey)(m.GetFieldMapping),
		internalSerialization.SerializeStringFunc(futureBookableMaxDaysKey)(m.GetFutureBookableMaxDays),
		internalSerialization.SerializeStringFunc(leadTimeKey)(m.GetLeadTime),
		internalSerialization.SerializeStringFunc(mandatoryKey)(m.GetMandatory),
		internalSerialization.SerializeStringFunc(useSlotEndTimeAsKey)(m.GetUseSlotEndTimeAs),
		internalSerialization.SerializeStringFunc(workDurationKey)(m.GetWorkDuration),
	)
}

func (m *ServiceConfigModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		activeKey:                      internalSerialization.DeserializeBoolFunc()(m.setActive),
		activeStringKey:                internalSerialization.DeserializeStringFunc()(m.setActiveString),
		appointmentBookingConfigKey:    internalSerialization.DeserializeStringFunc()(m.setAppointmentBookingConfig),
		appointmentDurationKey:         internalSerialization.DeserializeStringFunc()(m.setAppointmentDuration),
		appointmentsPerBookableSlotKey: internalSerialization.DeserializeStringFunc()(m.setAppointmentsPerBookableSlot),
		bookableDaysKey:                internalSerialization.DeserializeStringFunc()(m.setBookableDays),
		cancelByTimeKey:                internalSerialization.DeserializeStringFunc()(m.setCancelByTime),
		defaultTimezoneKey:             internalSerialization.DeserializeStringFunc()(m.setDefaultTimezone),
		enableAdvancedConfigKey:        internalSerialization.DeserializeBoolFunc()(m.setEnableAdvancedConfig),
		fieldMappingKey:                internalSerialization.DeserializeObjectValueFunc[FieldMapping](CreateFieldMappingFromDiscriminatorValue)(m.setFieldMapping),
		futureBookableMaxDaysKey:       internalSerialization.DeserializeStringFunc()(m.setFutureBookableMaxDays),
		leadTimeKey:                    internalSerialization.DeserializeStringFunc()(m.setLeadTime),
		mandatoryKey:                   internalSerialization.DeserializeStringFunc()(m.setMandatory),
		useSlotEndTimeAsKey:            internalSerialization.DeserializeStringFunc()(m.setUseSlotEndTimeAs),
		workDurationKey:                internalSerialization.DeserializeStringFunc()(m.setWorkDuration),
	}
}

func (m *ServiceConfigModel) GetActive() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), activeKey)
}
func (m *ServiceConfigModel) setActive(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), activeKey, val)
}
func (m *ServiceConfigModel) GetActiveString() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), activeStringKey)
}
func (m *ServiceConfigModel) setActiveString(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), activeStringKey, val)
}
func (m *ServiceConfigModel) GetAppointmentBookingConfig() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), appointmentBookingConfigKey)
}
func (m *ServiceConfigModel) setAppointmentBookingConfig(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), appointmentBookingConfigKey, val)
}
func (m *ServiceConfigModel) GetAppointmentDuration() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), appointmentDurationKey)
}
func (m *ServiceConfigModel) setAppointmentDuration(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), appointmentDurationKey, val)
}
func (m *ServiceConfigModel) GetAppointmentsPerBookableSlot() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), appointmentsPerBookableSlotKey)
}
func (m *ServiceConfigModel) setAppointmentsPerBookableSlot(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), appointmentsPerBookableSlotKey, val)
}
func (m *ServiceConfigModel) GetBookableDays() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), bookableDaysKey)
}
func (m *ServiceConfigModel) setBookableDays(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), bookableDaysKey, val)
}
func (m *ServiceConfigModel) GetCancelByTime() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), cancelByTimeKey)
}
func (m *ServiceConfigModel) setCancelByTime(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), cancelByTimeKey, val)
}
func (m *ServiceConfigModel) GetDefaultTimezone() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), defaultTimezoneKey)
}
func (m *ServiceConfigModel) setDefaultTimezone(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), defaultTimezoneKey, val)
}
func (m *ServiceConfigModel) GetEnableAdvancedConfig() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), enableAdvancedConfigKey)
}
func (m *ServiceConfigModel) setEnableAdvancedConfig(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), enableAdvancedConfigKey, val)
}
func (m *ServiceConfigModel) GetFieldMapping() (FieldMapping, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, FieldMapping](m.GetBackingStore(), fieldMappingKey)
}
func (m *ServiceConfigModel) setFieldMapping(val FieldMapping) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), fieldMappingKey, val)
}
func (m *ServiceConfigModel) GetFutureBookableMaxDays() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), futureBookableMaxDaysKey)
}
func (m *ServiceConfigModel) setFutureBookableMaxDays(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), futureBookableMaxDaysKey, val)
}
func (m *ServiceConfigModel) GetLeadTime() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), leadTimeKey)
}
func (m *ServiceConfigModel) setLeadTime(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), leadTimeKey, val)
}
func (m *ServiceConfigModel) GetMandatory() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), mandatoryKey)
}
func (m *ServiceConfigModel) setMandatory(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), mandatoryKey, val)
}
func (m *ServiceConfigModel) GetUseSlotEndTimeAs() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), useSlotEndTimeAsKey)
}
func (m *ServiceConfigModel) setUseSlotEndTimeAs(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), useSlotEndTimeAsKey, val)
}
func (m *ServiceConfigModel) GetWorkDuration() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), workDurationKey)
}
func (m *ServiceConfigModel) setWorkDuration(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), workDurationKey, val)
}

// FieldMapping represents the field_mapping nested object.
type FieldMapping interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetContact() (*string, error)
	setContact(*string) error
	GetContactRPVariable() (RPVariable, error)
	setContactRPVariable(RPVariable) error
	GetLocation() (*string, error)
	setLocation(*string) error
	GetLocationRPVariable() (RPVariable, error)
	setLocationRPVariable(RPVariable) error
}

type FieldMappingModel struct {
	core.BaseModel
}

func NewFieldMapping() *FieldMappingModel {
	return &FieldMappingModel{
		BaseModel: *core.NewBaseModel(),
	}
}

func CreateFieldMappingFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewFieldMapping(), nil
}

func (m *FieldMappingModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(contactKey)(m.GetContact),
		internalSerialization.SerializeObjectValueFunc[RPVariable](contactRPVariableKey)(m.GetContactRPVariable),
		internalSerialization.SerializeStringFunc(locationKey)(m.GetLocation),
		internalSerialization.SerializeObjectValueFunc[RPVariable](locationRPVariableKey)(m.GetLocationRPVariable),
	)
}

func (m *FieldMappingModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		contactKey:            internalSerialization.DeserializeStringFunc()(m.setContact),
		contactRPVariableKey:  internalSerialization.DeserializeObjectValueFunc[RPVariable](CreateRPVariableFromDiscriminatorValue)(m.setContactRPVariable),
		locationKey:           internalSerialization.DeserializeStringFunc()(m.setLocation),
		locationRPVariableKey: internalSerialization.DeserializeObjectValueFunc[RPVariable](CreateRPVariableFromDiscriminatorValue)(m.setLocationRPVariable),
	}
}

func (m *FieldMappingModel) GetContact() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), contactKey)
}
func (m *FieldMappingModel) setContact(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), contactKey, val)
}
func (m *FieldMappingModel) GetContactRPVariable() (RPVariable, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, RPVariable](m.GetBackingStore(), contactRPVariableKey)
}
func (m *FieldMappingModel) setContactRPVariable(val RPVariable) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), contactRPVariableKey, val)
}
func (m *FieldMappingModel) GetLocation() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), locationKey)
}
func (m *FieldMappingModel) setLocation(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), locationKey, val)
}
func (m *FieldMappingModel) GetLocationRPVariable() (RPVariable, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, RPVariable](m.GetBackingStore(), locationRPVariableKey)
}
func (m *FieldMappingModel) setLocationRPVariable(val RPVariable) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), locationRPVariableKey, val)
}

// RPVariable represents the RPVariable nested object.
type RPVariable interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetDisplayName() (*string, error)
	setDisplayName(*string) error
	GetLabel() (*string, error)
	setLabel(*string) error
	GetName() (*string, error)
	setName(*string) error
}

type RPVariableModel struct {
	core.BaseModel
}

func NewRPVariable() *RPVariableModel {
	return &RPVariableModel{
		BaseModel: *core.NewBaseModel(),
	}
}

func CreateRPVariableFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRPVariable(), nil
}

func (m *RPVariableModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(displayNameKey)(m.GetDisplayName),
		internalSerialization.SerializeStringFunc(labelKey)(m.GetLabel),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
	)
}

func (m *RPVariableModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		displayNameKey: internalSerialization.DeserializeStringFunc()(m.setDisplayName),
		labelKey:       internalSerialization.DeserializeStringFunc()(m.setLabel),
		nameKey:        internalSerialization.DeserializeStringFunc()(m.setName),
	}
}

func (m *RPVariableModel) GetDisplayName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), displayNameKey)
}
func (m *RPVariableModel) setDisplayName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), displayNameKey, val)
}
func (m *RPVariableModel) GetLabel() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), labelKey)
}
func (m *RPVariableModel) setLabel(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), labelKey, val)
}
func (m *RPVariableModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}
func (m *RPVariableModel) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}

// UserDateFormatOptions represents userDateFormatOptions nested object.
type UserDateFormatOptions interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetDay() (*string, error)
	setDay(*string) error
	GetMonth() (*string, error)
	setMonth(*string) error
	GetWeek() (*string, error)
	setWeek(*string) error
	GetWeekday() (*string, error)
	setWeekday(*string) error
}

type UserDateFormatOptionsModel struct {
	core.BaseModel
}

func NewUserDateFormatOptions() *UserDateFormatOptionsModel {
	return &UserDateFormatOptionsModel{
		BaseModel: *core.NewBaseModel(),
	}
}

func CreateUserDateFormatOptionsFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUserDateFormatOptions(), nil
}

func (m *UserDateFormatOptionsModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(dayKey)(m.GetDay),
		internalSerialization.SerializeStringFunc(monthKey)(m.GetMonth),
		internalSerialization.SerializeStringFunc(weekKey)(m.GetWeek),
		internalSerialization.SerializeStringFunc(weekdayKey)(m.GetWeekday),
	)
}

func (m *UserDateFormatOptionsModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		dayKey:     internalSerialization.DeserializeStringFunc()(m.setDay),
		monthKey:   internalSerialization.DeserializeStringFunc()(m.setMonth),
		weekKey:    internalSerialization.DeserializeStringFunc()(m.setWeek),
		weekdayKey: internalSerialization.DeserializeStringFunc()(m.setWeekday),
	}
}

func (m *UserDateFormatOptionsModel) GetDay() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), dayKey)
}
func (m *UserDateFormatOptionsModel) setDay(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), dayKey, val)
}
func (m *UserDateFormatOptionsModel) GetMonth() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), monthKey)
}
func (m *UserDateFormatOptionsModel) setMonth(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), monthKey, val)
}
func (m *UserDateFormatOptionsModel) GetWeek() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), weekKey)
}
func (m *UserDateFormatOptionsModel) setWeek(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), weekKey, val)
}
func (m *UserDateFormatOptionsModel) GetWeekday() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), weekdayKey)
}
func (m *UserDateFormatOptionsModel) setWeekday(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), weekdayKey, val)
}

// UserTimeFormat represents userTimeFormat nested object.
type UserTimeFormat interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetType() (*string, error)
	setType(*string) error
	GetValue() (*string, error)
	setValue(*string) error
}

type UserTimeFormatModel struct {
	core.BaseModel
}

func NewUserTimeFormat() *UserTimeFormatModel {
	return &UserTimeFormatModel{
		BaseModel: *core.NewBaseModel(),
	}
}

func CreateUserTimeFormatFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUserTimeFormat(), nil
}

func (m *UserTimeFormatModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(typeKey)(m.GetType),
		internalSerialization.SerializeStringFunc(valueKey)(m.GetValue),
	)
}

func (m *UserTimeFormatModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		typeKey:  internalSerialization.DeserializeStringFunc()(m.setType),
		valueKey: internalSerialization.DeserializeStringFunc()(m.setValue),
	}
}

func (m *UserTimeFormatModel) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), typeKey)
}
func (m *UserTimeFormatModel) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), typeKey, val)
}
func (m *UserTimeFormatModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), valueKey)
}
func (m *UserTimeFormatModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), valueKey, val)
}

// UserTimeFormatOptions represents userTimeFormatOptions nested object.
type UserTimeFormatOptions interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetHour() (*string, error)
	setHour(*string) error
	GetHourCycle() (*string, error)
	setHourCycle(*string) error
	GetMinute() (*string, error)
	setMinute(*string) error
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
		internalSerialization.SerializeStringFunc(hourKey)(m.GetHour),
		internalSerialization.SerializeStringFunc(hourCycleKey)(m.GetHourCycle),
		internalSerialization.SerializeStringFunc(minuteKey)(m.GetMinute),
	)
}

func (m *UserTimeFormatOptionsModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		hourKey:      internalSerialization.DeserializeStringFunc()(m.setHour),
		hourCycleKey: internalSerialization.DeserializeStringFunc()(m.setHourCycle),
		minuteKey:    internalSerialization.DeserializeStringFunc()(m.setMinute),
	}
}

func (m *UserTimeFormatOptionsModel) GetHour() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), hourKey)
}
func (m *UserTimeFormatOptionsModel) setHour(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), hourKey, val)
}
func (m *UserTimeFormatOptionsModel) GetHourCycle() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), hourCycleKey)
}
func (m *UserTimeFormatOptionsModel) setHourCycle(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), hourCycleKey, val)
}
func (m *UserTimeFormatOptionsModel) GetMinute() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), minuteKey)
}
func (m *UserTimeFormatOptionsModel) setMinute(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), minuteKey, val)
}
