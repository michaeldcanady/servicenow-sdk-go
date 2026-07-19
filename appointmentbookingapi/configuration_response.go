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
type ConfigurationResponse = core.ServiceNowItemResponse[*ConfigurationResult]

// CreateConfigurationResponseFromDiscriminatorValue is a factory for creating a ConfigurationResponse.
func CreateConfigurationResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*ConfigurationResult](CreateConfigurationResultFromDiscriminatorValue), nil
}

// ConfigurationResponseModel is the implementation of ConfigurationResponse.
type ConfigurationResponseModel = core.BaseServiceNowItemResponse[*ConfigurationResult]

// ConfigurationResult represents the result object in configuration response.
type ConfigurationResult struct {
	core.BaseModel
}

// NewConfigurationResult creates a new instance of ConfigurationResultModel
func NewConfigurationResult() *ConfigurationResult {
	return &ConfigurationResult{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateConfigurationResultFromDiscriminatorValue is a factory for creating a ConfigurationResult model.
func CreateConfigurationResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewConfigurationResult(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *ConfigurationResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeBoolFunc(activeKey, m.GetActive),
		internalSerialization.SerializeStringFunc(activeStringKey, m.GetActiveString),
		internalSerialization.SerializeBoolFunc(advancedCalendarViewPortalKey, m.GetAdvancedCalendarViewPortal),
		internalSerialization.SerializeBoolFunc(autoAcceptanceKey, m.GetAutoAcceptance),
		internalSerialization.SerializeStringFunc(localeLanguageKey, m.GetLocaleLanguage),
		internalSerialization.SerializeObjectValueFunc(serviceConfigKey, m.GetServiceConfig),
		internalSerialization.SerializeStringFunc(taskTableKey, m.GetTaskTable),
		internalSerialization.SerializeAnyFunc(translationsKey, m.GetTranslations),
		internalSerialization.SerializeObjectValueFunc(userDateFormatOptionsKey, m.GetUserDateFormatOptions),
		internalSerialization.SerializeBoolFunc(useRRKey, m.GetUseRR),
		internalSerialization.SerializeObjectValueFunc(userTimeFormatKey, m.GetUserTimeFormat),
		internalSerialization.SerializeObjectValueFunc(userTimeFormatOptionsKey, m.GetUserTimeFormatOptions),
		internalSerialization.SerializeStringFunc(viewScaleKey, m.GetViewScale),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ConfigurationResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		activeKey:                     internalSerialization.DeserializeBoolFunc(m.SetActive),
		activeStringKey:               internalSerialization.DeserializeStringFunc(m.SetActiveString),
		advancedCalendarViewPortalKey: internalSerialization.DeserializeBoolFunc(m.SetAdvancedCalendarViewPortal),
		autoAcceptanceKey:             internalSerialization.DeserializeBoolFunc(m.SetAutoAcceptance),
		localeLanguageKey:             internalSerialization.DeserializeStringFunc(m.SetLocaleLanguage),
		serviceConfigKey:              internalSerialization.DeserializeObjectValueFunc(CreateServiceConfigFromDiscriminatorValue, m.SetServiceConfig),
		taskTableKey:                  internalSerialization.DeserializeStringFunc(m.SetTaskTable),
		translationsKey:               internalSerialization.DeserializeAnyFunc(m.SetTranslations),
		userDateFormatOptionsKey:      internalSerialization.DeserializeObjectValueFunc(CreateUserDateFormatOptionsFromDiscriminatorValue, m.SetUserDateFormatOptions),
		useRRKey:                      internalSerialization.DeserializeBoolFunc(m.SetUseRR),
		userTimeFormatKey:             internalSerialization.DeserializeObjectValueFunc(CreateUserTimeFormatFromDiscriminatorValue, m.SetUserTimeFormat),
		userTimeFormatOptionsKey:      internalSerialization.DeserializeObjectValueFunc(CreateUserTimeFormatOptionsFromDiscriminatorValue, m.SetUserTimeFormatOptions),
		viewScaleKey:                  internalSerialization.DeserializeStringFunc(m.SetViewScale),
	}
}

func (m *ConfigurationResult) GetActive() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *bool](m, activeKey)
}
func (m *ConfigurationResult) SetActive(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, activeKey, val)
}
func (m *ConfigurationResult) GetActiveString() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *string](m, activeStringKey)
}
func (m *ConfigurationResult) SetActiveString(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, activeStringKey, val)
}
func (m *ConfigurationResult) GetAdvancedCalendarViewPortal() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *bool](m, advancedCalendarViewPortalKey)
}
func (m *ConfigurationResult) SetAdvancedCalendarViewPortal(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, advancedCalendarViewPortalKey, val)
}
func (m *ConfigurationResult) GetAutoAcceptance() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *bool](m, autoAcceptanceKey)
}
func (m *ConfigurationResult) SetAutoAcceptance(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, autoAcceptanceKey, val)
}
func (m *ConfigurationResult) GetLocaleLanguage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *string](m, localeLanguageKey)
}
func (m *ConfigurationResult) SetLocaleLanguage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, localeLanguageKey, val)
}
func (m *ConfigurationResult) GetServiceConfig() (ServiceConfig, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, ServiceConfig](m, serviceConfigKey)
}
func (m *ConfigurationResult) SetServiceConfig(val ServiceConfig) error {
	return store.DefaultBackedModelMutatorFunc(m, serviceConfigKey, val)
}
func (m *ConfigurationResult) GetTaskTable() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *string](m, taskTableKey)
}
func (m *ConfigurationResult) SetTaskTable(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskTableKey, val)
}
func (m *ConfigurationResult) GetTranslations() (any, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, any](m, translationsKey)
}
func (m *ConfigurationResult) SetTranslations(val any) error {
	return store.DefaultBackedModelMutatorFunc(m, translationsKey, val)
}
func (m *ConfigurationResult) GetUserDateFormatOptions() (UserDateFormatOptions, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, UserDateFormatOptions](m, userDateFormatOptionsKey)
}
func (m *ConfigurationResult) SetUserDateFormatOptions(val UserDateFormatOptions) error {
	return store.DefaultBackedModelMutatorFunc(m, userDateFormatOptionsKey, val)
}
func (m *ConfigurationResult) GetUseRR() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *bool](m, useRRKey)
}
func (m *ConfigurationResult) SetUseRR(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, useRRKey, val)
}
func (m *ConfigurationResult) GetUserTimeFormat() (UserTimeFormat, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, UserTimeFormat](m, userTimeFormatKey)
}
func (m *ConfigurationResult) SetUserTimeFormat(val UserTimeFormat) error {
	return store.DefaultBackedModelMutatorFunc(m, userTimeFormatKey, val)
}
func (m *ConfigurationResult) GetUserTimeFormatOptions() (UserTimeFormatOptions, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, UserTimeFormatOptions](m, userTimeFormatOptionsKey)
}
func (m *ConfigurationResult) SetUserTimeFormatOptions(val UserTimeFormatOptions) error {
	return store.DefaultBackedModelMutatorFunc(m, userTimeFormatOptionsKey, val)
}
func (m *ConfigurationResult) GetViewScale() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *string](m, viewScaleKey)
}
func (m *ConfigurationResult) SetViewScale(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, viewScaleKey, val)
}

// UserDateFormatOptions represents userDateFormatOptions nested object.
type UserDateFormatOptions interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetDay() (*string, error)
	SetDay(*string) error
	GetMonth() (*string, error)
	SetMonth(*string) error
	GetWeek() (*string, error)
	SetWeek(*string) error
	GetWeekday() (*string, error)
	SetWeekday(*string) error
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
		dayKey:     internalSerialization.DeserializeStringFunc(m.SetDay),
		monthKey:   internalSerialization.DeserializeStringFunc(m.SetMonth),
		weekKey:    internalSerialization.DeserializeStringFunc(m.SetWeek),
		weekdayKey: internalSerialization.DeserializeStringFunc(m.SetWeekday),
	}
}

func (m *UserDateFormatOptionsModel) GetDay() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserDateFormatOptionsModel, *string](m, dayKey)
}
func (m *UserDateFormatOptionsModel) SetDay(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, dayKey, val)
}
func (m *UserDateFormatOptionsModel) GetMonth() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserDateFormatOptionsModel, *string](m, monthKey)
}
func (m *UserDateFormatOptionsModel) SetMonth(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, monthKey, val)
}
func (m *UserDateFormatOptionsModel) GetWeek() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserDateFormatOptionsModel, *string](m, weekKey)
}
func (m *UserDateFormatOptionsModel) SetWeek(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, weekKey, val)
}
func (m *UserDateFormatOptionsModel) GetWeekday() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserDateFormatOptionsModel, *string](m, weekdayKey)
}
func (m *UserDateFormatOptionsModel) SetWeekday(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, weekdayKey, val)
}
