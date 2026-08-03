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

// GetActive returns the active value.
func (m *ConfigurationResult) GetActive() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *bool](m, activeKey)
}

// SetActive sets the active value.
func (m *ConfigurationResult) SetActive(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, activeKey, val)
}

// GetActiveString returns the active string value.
func (m *ConfigurationResult) GetActiveString() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *string](m, activeStringKey)
}

// SetActiveString sets the active string value.
func (m *ConfigurationResult) SetActiveString(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, activeStringKey, val)
}

// GetAdvancedCalendarViewPortal returns the advanced calendar view portal value.
func (m *ConfigurationResult) GetAdvancedCalendarViewPortal() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *bool](m, advancedCalendarViewPortalKey)
}

// SetAdvancedCalendarViewPortal sets the advanced calendar view portal value.
func (m *ConfigurationResult) SetAdvancedCalendarViewPortal(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, advancedCalendarViewPortalKey, val)
}

// GetAutoAcceptance returns the auto acceptance value.
func (m *ConfigurationResult) GetAutoAcceptance() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *bool](m, autoAcceptanceKey)
}

// SetAutoAcceptance sets the auto acceptance value.
func (m *ConfigurationResult) SetAutoAcceptance(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, autoAcceptanceKey, val)
}

// GetLocaleLanguage returns the locale language value.
func (m *ConfigurationResult) GetLocaleLanguage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *string](m, localeLanguageKey)
}

// SetLocaleLanguage sets the locale language value.
func (m *ConfigurationResult) SetLocaleLanguage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, localeLanguageKey, val)
}

// GetServiceConfig returns the service config value.
func (m *ConfigurationResult) GetServiceConfig() (ServiceConfig, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, ServiceConfig](m, serviceConfigKey)
}

// SetServiceConfig sets the service config value.
func (m *ConfigurationResult) SetServiceConfig(val ServiceConfig) error {
	return store.DefaultBackedModelMutatorFunc(m, serviceConfigKey, val)
}

// GetTaskTable returns the task table value.
func (m *ConfigurationResult) GetTaskTable() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *string](m, taskTableKey)
}

// SetTaskTable sets the task table value.
func (m *ConfigurationResult) SetTaskTable(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskTableKey, val)
}

// GetTranslations returns the translations value.
func (m *ConfigurationResult) GetTranslations() (any, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, any](m, translationsKey)
}

// SetTranslations sets the translations value.
func (m *ConfigurationResult) SetTranslations(val any) error {
	return store.DefaultBackedModelMutatorFunc(m, translationsKey, val)
}

// GetUserDateFormatOptions returns the user date format options value.
func (m *ConfigurationResult) GetUserDateFormatOptions() (UserDateFormatOptions, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, UserDateFormatOptions](m, userDateFormatOptionsKey)
}

// SetUserDateFormatOptions sets the user date format options value.
func (m *ConfigurationResult) SetUserDateFormatOptions(val UserDateFormatOptions) error {
	return store.DefaultBackedModelMutatorFunc(m, userDateFormatOptionsKey, val)
}

// GetUseRR returns the use rr value.
func (m *ConfigurationResult) GetUseRR() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *bool](m, useRRKey)
}

// SetUseRR sets the use rr value.
func (m *ConfigurationResult) SetUseRR(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, useRRKey, val)
}

// GetUserTimeFormat returns the user time format value.
func (m *ConfigurationResult) GetUserTimeFormat() (UserTimeFormat, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, UserTimeFormat](m, userTimeFormatKey)
}

// SetUserTimeFormat sets the user time format value.
func (m *ConfigurationResult) SetUserTimeFormat(val UserTimeFormat) error {
	return store.DefaultBackedModelMutatorFunc(m, userTimeFormatKey, val)
}

// GetUserTimeFormatOptions returns the user time format options value.
func (m *ConfigurationResult) GetUserTimeFormatOptions() (UserTimeFormatOptions, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, UserTimeFormatOptions](m, userTimeFormatOptionsKey)
}

// SetUserTimeFormatOptions sets the user time format options value.
func (m *ConfigurationResult) SetUserTimeFormatOptions(val UserTimeFormatOptions) error {
	return store.DefaultBackedModelMutatorFunc(m, userTimeFormatOptionsKey, val)
}

// GetViewScale returns the view scale value.
func (m *ConfigurationResult) GetViewScale() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *string](m, viewScaleKey)
}

// SetViewScale sets the view scale value.
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
		internalSerialization.SerializeStringFunc(monthKey, m.GetMonth),
		internalSerialization.SerializeStringFunc(weekKey, m.GetWeek),
		internalSerialization.SerializeStringFunc(weekdayKey, m.GetWeekday),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *UserDateFormatOptionsModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		dayKey:     internalSerialization.DeserializeStringFunc(m.SetDay),
		monthKey:   internalSerialization.DeserializeStringFunc(m.SetMonth),
		weekKey:    internalSerialization.DeserializeStringFunc(m.SetWeek),
		weekdayKey: internalSerialization.DeserializeStringFunc(m.SetWeekday),
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
func (m *UserDateFormatOptionsModel) GetMonth() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserDateFormatOptionsModel, *string](m, monthKey)
}

// SetMonth sets the month value.
func (m *UserDateFormatOptionsModel) SetMonth(val *string) error {
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
func (m *UserDateFormatOptionsModel) GetWeekday() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserDateFormatOptionsModel, *string](m, weekdayKey)
}

// SetWeekday sets the weekday value.
func (m *UserDateFormatOptionsModel) SetWeekday(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, weekdayKey, val)
}
