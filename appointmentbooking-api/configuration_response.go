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

type ConfigurationResponse = newInternal.ServiceNowItemResponse[*ConfigurationResultModel]

func CreateConfigurationResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*ConfigurationResultModel](CreateConfigurationResultFromDiscriminatorValue), nil
}

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

func (m *ConfigurationResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		activeKey:                      internalSerialization.DeserializeBoolFunc()(m.setActive),
		activeStringKey:                internalSerialization.DeserializeStringFunc()(m.setActiveString),
		advancedCalendarViewPortalKey:  internalSerialization.DeserializeBoolFunc()(m.setAdvancedCalendarViewPortal),
		autoAcceptanceKey:              internalSerialization.DeserializeBoolFunc()(m.setAutoAcceptance),
		localeLanguageKey:              internalSerialization.DeserializeStringFunc()(m.setLocaleLanguage),
		serviceConfigKey:               internalSerialization.DeserializeObjectValueFunc[ServiceConfig](CreateServiceConfigFromDiscriminatorValue)(m.setServiceConfig),
		taskTableKey:                   internalSerialization.DeserializeStringFunc()(m.setTaskTable),
		translationsKey:                internalSerialization.DeserializeAnyFunc()(m.setTranslations),
		userDateFormatOptionsKey:       internalSerialization.DeserializeObjectValueFunc[UserDateFormatOptions](CreateUserDateFormatOptionsFromDiscriminatorValue)(m.setUserDateFormatOptions),
		useRRKey:                       internalSerialization.DeserializeBoolFunc()(m.setUseRR),
		userTimeFormatKey:              internalSerialization.DeserializeObjectValueFunc[UserTimeFormat](CreateUserTimeFormatFromDiscriminatorValue)(m.setUserTimeFormat),
		userTimeFormatOptionsKey:       internalSerialization.DeserializeObjectValueFunc[UserTimeFormatOptions](CreateUserTimeFormatOptionsFromDiscriminatorValue)(m.setUserTimeFormatOptions),
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
func (m *ConfigurationResultModel) GetServiceConfig() (ServiceConfig, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, ServiceConfig](m.GetBackingStore(), serviceConfigKey) }
func (m *ConfigurationResultModel) setServiceConfig(val ServiceConfig) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), serviceConfigKey, val) }
func (m *ConfigurationResultModel) GetTaskTable() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), taskTableKey) }
func (m *ConfigurationResultModel) setTaskTable(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), taskTableKey, val) }
func (m *ConfigurationResultModel) GetTranslations() (any, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, any](m.GetBackingStore(), translationsKey) }
func (m *ConfigurationResultModel) setTranslations(val any) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), translationsKey, val) }
func (m *ConfigurationResultModel) GetUserDateFormatOptions() (UserDateFormatOptions, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, UserDateFormatOptions](m.GetBackingStore(), userDateFormatOptionsKey) }
func (m *ConfigurationResultModel) setUserDateFormatOptions(val UserDateFormatOptions) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), userDateFormatOptionsKey, val) }
func (m *ConfigurationResultModel) GetUseRR() (*bool, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), useRRKey) }
func (m *ConfigurationResultModel) setUseRR(val *bool) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), useRRKey, val) }
func (m *ConfigurationResultModel) GetUserTimeFormat() (UserTimeFormat, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, UserTimeFormat](m.GetBackingStore(), userTimeFormatKey) }
func (m *ConfigurationResultModel) setUserTimeFormat(val UserTimeFormat) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), userTimeFormatKey, val) }
func (m *ConfigurationResultModel) GetUserTimeFormatOptions() (UserTimeFormatOptions, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, UserTimeFormatOptions](m.GetBackingStore(), userTimeFormatOptionsKey) }
func (m *ConfigurationResultModel) setUserTimeFormatOptions(val UserTimeFormatOptions) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), userTimeFormatOptionsKey, val) }
func (m *ConfigurationResultModel) GetViewScale() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), viewScaleKey) }
func (m *ConfigurationResultModel) setViewScale(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), viewScaleKey, val) }

func CreateConfigurationResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewConfigurationResult(), nil
}

type ServiceConfig interface { serialization.Parsable; kiotaStore.BackedModel }
type ServiceConfigModel struct { newInternal.BaseModel }
func NewServiceConfig() *ServiceConfigModel { return &ServiceConfigModel{BaseModel: *newInternal.NewBaseModel()} }
func (m *ServiceConfigModel) Serialize(writer serialization.SerializationWriter) error { val, _ := m.GetBackingStore().Get("additionalData"); if val != nil { return writer.WriteAdditionalData(val.(map[string]interface{})) }; return nil }
func (m *ServiceConfigModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error { return map[string]func(serialization.ParseNode) error{"*": func(n serialization.ParseNode) error { val, _ := n.GetRawValue(); m.GetBackingStore().Set("additionalData", val); return nil }} }
func CreateServiceConfigFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) { return NewServiceConfig(), nil }

type UserDateFormatOptions interface { serialization.Parsable; kiotaStore.BackedModel }
type UserDateFormatOptionsModel struct { newInternal.BaseModel }
func NewUserDateFormatOptions() *UserDateFormatOptionsModel { return &UserDateFormatOptionsModel{BaseModel: *newInternal.NewBaseModel()} }
func (m *UserDateFormatOptionsModel) Serialize(writer serialization.SerializationWriter) error { val, _ := m.GetBackingStore().Get("additionalData"); if val != nil { return writer.WriteAdditionalData(val.(map[string]interface{})) }; return nil }
func (m *UserDateFormatOptionsModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error { return map[string]func(serialization.ParseNode) error{"*": func(n serialization.ParseNode) error { val, _ := n.GetRawValue(); m.GetBackingStore().Set("additionalData", val); return nil }} }
func CreateUserDateFormatOptionsFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) { return NewUserDateFormatOptions(), nil }

type UserTimeFormat interface { serialization.Parsable; kiotaStore.BackedModel }
type UserTimeFormatModel struct { newInternal.BaseModel }
func NewUserTimeFormat() *UserTimeFormatModel { return &UserTimeFormatModel{BaseModel: *newInternal.NewBaseModel()} }
func (m *UserTimeFormatModel) Serialize(writer serialization.SerializationWriter) error { val, _ := m.GetBackingStore().Get("additionalData"); if val != nil { return writer.WriteAdditionalData(val.(map[string]interface{})) }; return nil }
func (m *UserTimeFormatModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error { return map[string]func(serialization.ParseNode) error{"*": func(n serialization.ParseNode) error { val, _ := n.GetRawValue(); m.GetBackingStore().Set("additionalData", val); return nil }} }
func CreateUserTimeFormatFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) { return NewUserTimeFormat(), nil }

type UserTimeFormatOptions interface { serialization.Parsable; kiotaStore.BackedModel }
type UserTimeFormatOptionsModel struct { newInternal.BaseModel }
func NewUserTimeFormatOptions() *UserTimeFormatOptionsModel { return &UserTimeFormatOptionsModel{BaseModel: *newInternal.NewBaseModel()} }
func (m *UserTimeFormatOptionsModel) Serialize(writer serialization.SerializationWriter) error { val, _ := m.GetBackingStore().Get("additionalData"); if val != nil { return writer.WriteAdditionalData(val.(map[string]interface{})) }; return nil }
func (m *UserTimeFormatOptionsModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error { return map[string]func(serialization.ParseNode) error{"*": func(n serialization.ParseNode) error { val, _ := n.GetRawValue(); m.GetBackingStore().Set("additionalData", val); return nil }} }
func CreateUserTimeFormatOptionsFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) { return NewUserTimeFormatOptions(), nil }
