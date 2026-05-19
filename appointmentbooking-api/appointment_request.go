package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)


type AppointmentRequest interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetActualEndDate() (*string, error)
	setActualEndDate(*string) error
	GetActualStartDate() (*string, error)
	setActualStartDate(*string) error
	GetCatalogId() (*string, error)
	setCatalogId(*string) error
	GetEndDateUTC() (*string, error)
	setEndDateUTC(*string) error
	GetLocation() (*string, error)
	setLocation(*string) error
	GetOpenedFor() (*string, error)
	setOpenedFor(*string) error
	GetReschedule() (*bool, error)
	setReschedule(*bool) error
	GetServiceConfigRule() (*string, error)
	setServiceConfigRule(*string) error
	GetStartDateUTC() (*string, error)
	setStartDateUTC(*string) error
	GetTaskId() (*string, error)
	setTaskId(*string) error
	GetTaskTable() (*string, error)
	setTaskTable(*string) error
	GetTimezone() (*string, error)
	setTimezone(*string) error
	GetValidateRequest() (*bool, error)
	setValidateRequest(*bool) error
}

type AppointmentRequestModel struct {
	newInternal.BaseModel
}

func NewAppointmentRequest() *AppointmentRequestModel {
	return &AppointmentRequestModel{BaseModel: *newInternal.NewBaseModel()}
}

func (m *AppointmentRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(actualEndDateKey)(m.GetActualEndDate),
		internalSerialization.SerializeStringFunc(actualStartDateKey)(m.GetActualStartDate),
		internalSerialization.SerializeStringFunc(catalogIdKey)(m.GetCatalogId),
		internalSerialization.SerializeStringFunc(endDateUTCKey)(m.GetEndDateUTC),
		internalSerialization.SerializeStringFunc(locationKey)(m.GetLocation),
		internalSerialization.SerializeStringFunc(openedForKey)(m.GetOpenedFor),
		internalSerialization.SerializeBoolFunc(rescheduleKey)(m.GetReschedule),
		internalSerialization.SerializeStringFunc(serviceConfigRuleKey)(m.GetServiceConfigRule),
		internalSerialization.SerializeStringFunc(startDateUTCKey)(m.GetStartDateUTC),
		internalSerialization.SerializeStringFunc(taskIdKey)(m.GetTaskId),
		internalSerialization.SerializeStringFunc(taskTableKey)(m.GetTaskTable),
		internalSerialization.SerializeStringFunc(timezoneKey)(m.GetTimezone),
		internalSerialization.SerializeBoolFunc(validateRequestKey)(m.GetValidateRequest),
	)
}

func (m *AppointmentRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		actualEndDateKey:     internalSerialization.DeserializeStringFunc()(m.setActualEndDate),
		actualStartDateKey:   internalSerialization.DeserializeStringFunc()(m.setActualStartDate),
		catalogIdKey:         internalSerialization.DeserializeStringFunc()(m.setCatalogId),
		endDateUTCKey:        internalSerialization.DeserializeStringFunc()(m.setEndDateUTC),
		locationKey:          internalSerialization.DeserializeStringFunc()(m.setLocation),
		openedForKey:         internalSerialization.DeserializeStringFunc()(m.setOpenedFor),
		rescheduleKey:        internalSerialization.DeserializeBoolFunc()(m.setReschedule),
		serviceConfigRuleKey: internalSerialization.DeserializeStringFunc()(m.setServiceConfigRule),
		startDateUTCKey:      internalSerialization.DeserializeStringFunc()(m.setStartDateUTC),
		taskIdKey:            internalSerialization.DeserializeStringFunc()(m.setTaskId),
		taskTableKey:         internalSerialization.DeserializeStringFunc()(m.setTaskTable),
		timezoneKey:          internalSerialization.DeserializeStringFunc()(m.setTimezone),
		validateRequestKey:   internalSerialization.DeserializeBoolFunc()(m.setValidateRequest),
	}
}

func (m *AppointmentRequestModel) GetActualEndDate() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), actualEndDateKey) }
func (m *AppointmentRequestModel) setActualEndDate(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), actualEndDateKey, val) }
func (m *AppointmentRequestModel) GetActualStartDate() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), actualStartDateKey) }
func (m *AppointmentRequestModel) setActualStartDate(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), actualStartDateKey, val) }
func (m *AppointmentRequestModel) GetCatalogId() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), catalogIdKey) }
func (m *AppointmentRequestModel) setCatalogId(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), catalogIdKey, val) }
func (m *AppointmentRequestModel) GetEndDateUTC() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), endDateUTCKey) }
func (m *AppointmentRequestModel) setEndDateUTC(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), endDateUTCKey, val) }
func (m *AppointmentRequestModel) GetLocation() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), locationKey) }
func (m *AppointmentRequestModel) setLocation(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), locationKey, val) }
func (m *AppointmentRequestModel) GetOpenedFor() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), openedForKey) }
func (m *AppointmentRequestModel) setOpenedFor(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), openedForKey, val) }
func (m *AppointmentRequestModel) GetReschedule() (*bool, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), rescheduleKey) }
func (m *AppointmentRequestModel) setReschedule(val *bool) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), rescheduleKey, val) }
func (m *AppointmentRequestModel) GetServiceConfigRule() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), serviceConfigRuleKey) }
func (m *AppointmentRequestModel) setServiceConfigRule(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), serviceConfigRuleKey, val) }
func (m *AppointmentRequestModel) GetStartDateUTC() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), startDateUTCKey) }
func (m *AppointmentRequestModel) setStartDateUTC(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), startDateUTCKey, val) }
func (m *AppointmentRequestModel) GetTaskId() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), taskIdKey) }
func (m *AppointmentRequestModel) setTaskId(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), taskIdKey, val) }
func (m *AppointmentRequestModel) GetTaskTable() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), taskTableKey) }
func (m *AppointmentRequestModel) setTaskTable(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), taskTableKey, val) }
func (m *AppointmentRequestModel) GetTimezone() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), timezoneKey) }
func (m *AppointmentRequestModel) setTimezone(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), timezoneKey, val) }
func (m *AppointmentRequestModel) GetValidateRequest() (*bool, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), validateRequestKey) }
func (m *AppointmentRequestModel) setValidateRequest(val *bool) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), validateRequestKey, val) }

func CreateAppointmentRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAppointmentRequest(), nil
}
