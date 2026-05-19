package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	activeKey                      = "active"
	activeStringKey                = "active_string"
	advancedCalendarViewPortalKey  = "advanced_calendar_view_portal"
	autoAcceptanceKey              = "auto_acceptance"
	localeLanguageKey              = "locale_language"
	serviceConfigKey               = "service_config"
	taskTableKey                   = "task_table"
	translationsKey                = "translations"
	userDateFormatOptionsKey       = "userDateFormatOptions"
	useRRKey                       = "useRR"
	userTimeFormatKey              = "userTimeFormat"
	userTimeFormatOptionsKey       = "userTimeFormatOptions"
	viewScaleKey                   = "view_scale"
)

// ConfigurationResponse represents the configuration response.
type ConfigurationResponse = newInternal.ServiceNowItemResponse[*ConfigurationResultModel]

// CreateConfigurationResponseFromDiscriminatorValue is a factory for creating a ConfigurationResponse.
func CreateConfigurationResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*ConfigurationResultModel](CreateConfigurationResultFromDiscriminatorValue), nil
}

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
	GetServiceConfig() (any, error)
	setServiceConfig(any) error
	GetTaskTable() (*string, error)
	setTaskTable(*string) error
	GetTranslations() (any, error)
	setTranslations(any) error
	GetUserDateFormatOptions() (any, error)
	setUserDateFormatOptions(any) error
	GetUseRR() (*bool, error)
	setUseRR(*bool) error
	GetUserTimeFormat() (any, error)
	setUserTimeFormat(any) error
	GetUserTimeFormatOptions() (any, error)
	setUserTimeFormatOptions(any) error
	GetViewScale() (*string, error)
	setViewScale(*string) error
}

// ConfigurationResultModel implementation of ConfigurationResult
type ConfigurationResultModel struct {
	newInternal.BaseModel
}

func NewConfigurationResult() *ConfigurationResultModel {
	return &ConfigurationResultModel{
		BaseModel: *newInternal.NewBaseModel(),
	}
}

func (m *ConfigurationResultModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeBoolFunc(activeKey)(m.GetActive),
		internalSerialization.SerializeStringFunc(activeStringKey)(m.GetActiveString),
		internalSerialization.SerializeBoolFunc(advancedCalendarViewPortalKey)(m.GetAdvancedCalendarViewPortal),
		internalSerialization.SerializeBoolFunc(autoAcceptanceKey)(m.GetAutoAcceptance),
		internalSerialization.SerializeStringFunc(localeLanguageKey)(m.GetLocaleLanguage),
		internalSerialization.SerializeAnyFunc(serviceConfigKey)(m.GetServiceConfig),
		internalSerialization.SerializeStringFunc(taskTableKey)(m.GetTaskTable),
		internalSerialization.SerializeAnyFunc(translationsKey)(m.GetTranslations),
		internalSerialization.SerializeAnyFunc(userDateFormatOptionsKey)(m.GetUserDateFormatOptions),
		internalSerialization.SerializeBoolFunc(useRRKey)(m.GetUseRR),
		internalSerialization.SerializeAnyFunc(userTimeFormatKey)(m.GetUserTimeFormat),
		internalSerialization.SerializeAnyFunc(userTimeFormatOptionsKey)(m.GetUserTimeFormatOptions),
		internalSerialization.SerializeStringFunc(viewScaleKey)(m.GetViewScale),
	)
}

func (m *ConfigurationResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		activeKey:                      internalSerialization.DeserializeBoolFunc()(m.setActive),
		activeStringKey:                internalSerialization.DeserializeStringFunc()(m.setActiveString),
		advancedCalendarViewPortalKey:  internalSerialization.DeserializeBoolFunc()(m.setAdvancedCalendarViewPortal),
		autoAcceptanceKey:              internalSerialization.DeserializeBoolFunc()(m.setAutoAcceptance),
		localeLanguageKey:              internalSerialization.DeserializeStringFunc()(m.setLocaleLanguage),
		serviceConfigKey:               internalSerialization.DeserializeAnyFunc()(m.setServiceConfig),
		taskTableKey:                   internalSerialization.DeserializeStringFunc()(m.setTaskTable),
		translationsKey:                internalSerialization.DeserializeAnyFunc()(m.setTranslations),
		userDateFormatOptionsKey:       internalSerialization.DeserializeAnyFunc()(m.setUserDateFormatOptions),
		useRRKey:                       internalSerialization.DeserializeBoolFunc()(m.setUseRR),
		userTimeFormatKey:              internalSerialization.DeserializeAnyFunc()(m.setUserTimeFormat),
		userTimeFormatOptionsKey:       internalSerialization.DeserializeAnyFunc()(m.setUserTimeFormatOptions),
		viewScaleKey:                   internalSerialization.DeserializeStringFunc()(m.setViewScale),
	}
}

func (m *ConfigurationResultModel) GetActive() (*bool, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), activeKey) }
func (m *ConfigurationResultModel) setActive(val *bool) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), activeKey, val) }
func (m *ConfigurationResultModel) GetActiveString() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), activeStringKey) }
func (m *ConfigurationResultModel) setActiveString(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), activeStringKey, val) }
func (m *ConfigurationResultModel) GetAdvancedCalendarViewPortal() (*bool, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), advancedCalendarViewPortalKey) }
func (m *ConfigurationResultModel) setAdvancedCalendarViewPortal(val *bool) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), advancedCalendarViewPortalKey, val) }
func (m *ConfigurationResultModel) GetAutoAcceptance() (*bool, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), autoAcceptanceKey) }
func (m *ConfigurationResultModel) setAutoAcceptance(val *bool) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), autoAcceptanceKey, val) }
func (m *ConfigurationResultModel) GetLocaleLanguage() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), localeLanguageKey) }
func (m *ConfigurationResultModel) setLocaleLanguage(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), localeLanguageKey, val) }
func (m *ConfigurationResultModel) GetServiceConfig() (any, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, any](m.GetBackingStore(), serviceConfigKey) }
func (m *ConfigurationResultModel) setServiceConfig(val any) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), serviceConfigKey, val) }
func (m *ConfigurationResultModel) GetTaskTable() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), taskTableKey) }
func (m *ConfigurationResultModel) setTaskTable(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), taskTableKey, val) }
func (m *ConfigurationResultModel) GetTranslations() (any, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, any](m.GetBackingStore(), translationsKey) }
func (m *ConfigurationResultModel) setTranslations(val any) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), translationsKey, val) }
func (m *ConfigurationResultModel) GetUserDateFormatOptions() (any, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, any](m.GetBackingStore(), userDateFormatOptionsKey) }
func (m *ConfigurationResultModel) setUserDateFormatOptions(val any) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), userDateFormatOptionsKey, val) }
func (m *ConfigurationResultModel) GetUseRR() (*bool, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), useRRKey) }
func (m *ConfigurationResultModel) setUseRR(val *bool) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), useRRKey, val) }
func (m *ConfigurationResultModel) GetUserTimeFormat() (any, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, any](m.GetBackingStore(), userTimeFormatKey) }
func (m *ConfigurationResultModel) setUserTimeFormat(val any) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), userTimeFormatKey, val) }
func (m *ConfigurationResultModel) GetUserTimeFormatOptions() (any, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, any](m.GetBackingStore(), userTimeFormatOptionsKey) }
func (m *ConfigurationResultModel) setUserTimeFormatOptions(val any) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), userTimeFormatOptionsKey, val) }
func (m *ConfigurationResultModel) GetViewScale() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), viewScaleKey) }
func (m *ConfigurationResultModel) setViewScale(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), viewScaleKey, val) }

func CreateConfigurationResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewConfigurationResult(), nil
}
