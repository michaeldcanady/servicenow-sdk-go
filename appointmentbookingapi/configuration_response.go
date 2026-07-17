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
		internalSerialization.SerializeBoolFunc(activeKey, m.GetActive),
		internalSerialization.SerializeStringFunc(activeStringKey, m.GetActiveString),
		internalSerialization.SerializeBoolFunc(advancedCalendarViewPortalKey, m.GetAdvancedCalendarViewPortal),
		internalSerialization.SerializeBoolFunc(autoAcceptanceKey, m.GetAutoAcceptance),
		internalSerialization.SerializeStringFunc(localeLanguageKey, m.GetLocaleLanguage),
		internalSerialization.SerializeObjectValueFunc[ServiceConfig](serviceConfigKey, m.GetServiceConfig),
		internalSerialization.SerializeStringFunc(taskTableKey, m.GetTaskTable),
		internalSerialization.SerializeAnyFunc(translationsKey, m.GetTranslations),
		internalSerialization.SerializeObjectValueFunc[UserDateFormatOptions](userDateFormatOptionsKey, m.GetUserDateFormatOptions),
		internalSerialization.SerializeBoolFunc(useRRKey, m.GetUseRR),
		internalSerialization.SerializeObjectValueFunc[UserTimeFormat](userTimeFormatKey, m.GetUserTimeFormat),
		internalSerialization.SerializeObjectValueFunc[UserTimeFormatOptions](userTimeFormatOptionsKey, m.GetUserTimeFormatOptions),
		internalSerialization.SerializeStringFunc(viewScaleKey, m.GetViewScale),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ConfigurationResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		activeKey:                     internalSerialization.DeserializeBoolFunc(m.setActive),
		activeStringKey:               internalSerialization.DeserializeStringFunc(m.setActiveString),
		advancedCalendarViewPortalKey: internalSerialization.DeserializeBoolFunc(m.setAdvancedCalendarViewPortal),
		autoAcceptanceKey:             internalSerialization.DeserializeBoolFunc(m.setAutoAcceptance),
		localeLanguageKey:             internalSerialization.DeserializeStringFunc(m.setLocaleLanguage),
		serviceConfigKey:              internalSerialization.DeserializeObjectValueFunc[ServiceConfig](CreateServiceConfigFromDiscriminatorValue, m.setServiceConfig),
		taskTableKey:                  internalSerialization.DeserializeStringFunc(m.setTaskTable),
		translationsKey:               internalSerialization.DeserializeAnyFunc(m.setTranslations),
		userDateFormatOptionsKey:      internalSerialization.DeserializeObjectValueFunc[UserDateFormatOptions](CreateUserDateFormatOptionsFromDiscriminatorValue, m.setUserDateFormatOptions),
		useRRKey:                      internalSerialization.DeserializeBoolFunc(m.setUseRR),
		userTimeFormatKey:             internalSerialization.DeserializeObjectValueFunc[UserTimeFormat](CreateUserTimeFormatFromDiscriminatorValue, m.setUserTimeFormat),
		userTimeFormatOptionsKey:      internalSerialization.DeserializeObjectValueFunc[UserTimeFormatOptions](CreateUserTimeFormatOptionsFromDiscriminatorValue, m.setUserTimeFormatOptions),
		viewScaleKey:                  internalSerialization.DeserializeStringFunc(m.setViewScale),
	}
}

func (m *ConfigurationResultModel) GetActive() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResultModel, *bool](m, activeKey)
}
func (m *ConfigurationResultModel) setActive(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, activeKey, val)
}
func (m *ConfigurationResultModel) GetActiveString() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResultModel, *string](m, activeStringKey)
}
func (m *ConfigurationResultModel) setActiveString(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, activeStringKey, val)
}
func (m *ConfigurationResultModel) GetAdvancedCalendarViewPortal() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResultModel, *bool](m, advancedCalendarViewPortalKey)
}
func (m *ConfigurationResultModel) setAdvancedCalendarViewPortal(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, advancedCalendarViewPortalKey, val)
}
func (m *ConfigurationResultModel) GetAutoAcceptance() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResultModel, *bool](m, autoAcceptanceKey)
}
func (m *ConfigurationResultModel) setAutoAcceptance(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, autoAcceptanceKey, val)
}
func (m *ConfigurationResultModel) GetLocaleLanguage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResultModel, *string](m, localeLanguageKey)
}
func (m *ConfigurationResultModel) setLocaleLanguage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, localeLanguageKey, val)
}
func (m *ConfigurationResultModel) GetServiceConfig() (ServiceConfig, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResultModel, ServiceConfig](m, serviceConfigKey)
}
func (m *ConfigurationResultModel) setServiceConfig(val ServiceConfig) error {
	return store.DefaultBackedModelMutatorFunc(m, serviceConfigKey, val)
}
func (m *ConfigurationResultModel) GetTaskTable() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResultModel, *string](m, taskTableKey)
}
func (m *ConfigurationResultModel) setTaskTable(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskTableKey, val)
}
func (m *ConfigurationResultModel) GetTranslations() (any, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResultModel, any](m, translationsKey)
}
func (m *ConfigurationResultModel) setTranslations(val any) error {
	return store.DefaultBackedModelMutatorFunc(m, translationsKey, val)
}
func (m *ConfigurationResultModel) GetUserDateFormatOptions() (UserDateFormatOptions, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResultModel, UserDateFormatOptions](m, userDateFormatOptionsKey)
}
func (m *ConfigurationResultModel) setUserDateFormatOptions(val UserDateFormatOptions) error {
	return store.DefaultBackedModelMutatorFunc(m, userDateFormatOptionsKey, val)
}
func (m *ConfigurationResultModel) GetUseRR() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResultModel, *bool](m, useRRKey)
}
func (m *ConfigurationResultModel) setUseRR(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, useRRKey, val)
}
func (m *ConfigurationResultModel) GetUserTimeFormat() (UserTimeFormat, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResultModel, UserTimeFormat](m, userTimeFormatKey)
}
func (m *ConfigurationResultModel) setUserTimeFormat(val UserTimeFormat) error {
	return store.DefaultBackedModelMutatorFunc(m, userTimeFormatKey, val)
}
func (m *ConfigurationResultModel) GetUserTimeFormatOptions() (UserTimeFormatOptions, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResultModel, UserTimeFormatOptions](m, userTimeFormatOptionsKey)
}
func (m *ConfigurationResultModel) setUserTimeFormatOptions(val UserTimeFormatOptions) error {
	return store.DefaultBackedModelMutatorFunc(m, userTimeFormatOptionsKey, val)
}
func (m *ConfigurationResultModel) GetViewScale() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResultModel, *string](m, viewScaleKey)
}
func (m *ConfigurationResultModel) setViewScale(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, viewScaleKey, val)
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
		internalSerialization.SerializeStringFunc(dayKey, m.GetDay),
		internalSerialization.SerializeStringFunc(monthKey, m.GetMonth),
		internalSerialization.SerializeStringFunc(weekKey, m.GetWeek),
		internalSerialization.SerializeStringFunc(weekdayKey, m.GetWeekday),
	)
}

func (m *UserDateFormatOptionsModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		dayKey:     internalSerialization.DeserializeStringFunc(m.setDay),
		monthKey:   internalSerialization.DeserializeStringFunc(m.setMonth),
		weekKey:    internalSerialization.DeserializeStringFunc(m.setWeek),
		weekdayKey: internalSerialization.DeserializeStringFunc(m.setWeekday),
	}
}

func (m *UserDateFormatOptionsModel) GetDay() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserDateFormatOptionsModel, *string](m, dayKey)
}
func (m *UserDateFormatOptionsModel) setDay(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, dayKey, val)
}
func (m *UserDateFormatOptionsModel) GetMonth() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserDateFormatOptionsModel, *string](m, monthKey)
}
func (m *UserDateFormatOptionsModel) setMonth(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, monthKey, val)
}
func (m *UserDateFormatOptionsModel) GetWeek() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserDateFormatOptionsModel, *string](m, weekKey)
}
func (m *UserDateFormatOptionsModel) setWeek(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, weekKey, val)
}
func (m *UserDateFormatOptionsModel) GetWeekday() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserDateFormatOptionsModel, *string](m, weekdayKey)
}
func (m *UserDateFormatOptionsModel) setWeekday(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, weekdayKey, val)
}
