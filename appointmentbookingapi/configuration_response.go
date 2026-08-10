package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
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
		internalSerialization.SerializeEnumFunc(viewScaleKey, m.GetViewScale),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ConfigurationResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		activeKey:                     internalSerialization.DeserializeBoolFunc(m.SetActive),                     // V
		activeStringKey:               internalSerialization.DeserializeStringFunc(m.SetActiveString),             // V
		advancedCalendarViewPortalKey: internalSerialization.DeserializeBoolFunc(m.SetAdvancedCalendarViewPortal), // V
		autoAcceptanceKey:             internalSerialization.DeserializeBoolFunc(m.SetAutoAcceptance),             // V
		// TODO:  ISO 639.1 language code
		localeLanguageKey: internalSerialization.DeserializeStringFunc(m.SetLocaleLanguage),                                                // V
		serviceConfigKey:  internalSerialization.DeserializeObjectValueFunc(CreateServiceConfigFromDiscriminatorValue, m.SetServiceConfig), // V
		taskTableKey:      internalSerialization.DeserializeStringFunc(m.SetTaskTable),                                                     // V
		// TODO: "object," arbitrary map key-value pairs
		translationsKey:          internalSerialization.DeserializeAnyFunc(m.SetTranslations),                                                                     // V
		userDateFormatOptionsKey: internalSerialization.DeserializeObjectValueFunc(CreateUserDateFormatOptionsFromDiscriminatorValue, m.SetUserDateFormatOptions), // V
		useRRKey:                 internalSerialization.DeserializeBoolFunc(m.SetUseRR),                                                                           // V
		userTimeFormatKey:        internalSerialization.DeserializeObjectValueFunc(CreateUserTimeFormatFromDiscriminatorValue, m.SetUserTimeFormat),               // V
		userTimeFormatOptionsKey: internalSerialization.DeserializeObjectValueFunc(CreateUserTimeFormatOptionsFromDiscriminatorValue, m.SetUserTimeFormatOptions), // V
		viewScaleKey:             internalSerialization.DeserializeEnumFunc(ParseViewScale, m.SetViewScale),                                                       // V
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
func (m *ConfigurationResult) GetViewScale() (*ViewScale, error) {
	return store.DefaultBackedModelAccessorFunc[*ConfigurationResult, *ViewScale](m, viewScaleKey)
}

// SetViewScale sets the view scale value.
func (m *ConfigurationResult) SetViewScale(val *ViewScale) error {
	return store.DefaultBackedModelMutatorFunc(m, viewScaleKey, val)
}
