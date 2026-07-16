package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
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
	core.BaseModel
}

func NewAppointmentRequest() *AppointmentRequestModel {
	return &AppointmentRequestModel{BaseModel: *core.NewBaseModel()}
}

func CreateAppointmentRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAppointmentRequest(), nil
}

func (m *AppointmentRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(actualEndDateKey)(m.GetActualEndDate),
		internalSerialization.SerializeStringFunc(actualStartDateKey)(m.GetActualStartDate),
		internalSerialization.SerializeStringFunc(catalogIDKey)(m.GetCatalogId),
		internalSerialization.SerializeStringFunc(endDateUTCKey)(m.GetEndDateUTC),
		internalSerialization.SerializeStringFunc(locationKey)(m.GetLocation),
		internalSerialization.SerializeStringFunc(openedForKey)(m.GetOpenedFor),
		internalSerialization.SerializeBoolFunc(rescheduleKey)(m.GetReschedule),
		internalSerialization.SerializeStringFunc(serviceConfigRuleKey)(m.GetServiceConfigRule),
		internalSerialization.SerializeStringFunc(startDateUTCKey)(m.GetStartDateUTC),
		internalSerialization.SerializeStringFunc(taskIDKey)(m.GetTaskId),
		internalSerialization.SerializeStringFunc(taskTableKey)(m.GetTaskTable),
		internalSerialization.SerializeStringFunc(timezoneKey)(m.GetTimezone),
		internalSerialization.SerializeBoolFunc(validateRequestKey)(m.GetValidateRequest),
	)
}

func (m *AppointmentRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		actualEndDateKey:     internalSerialization.DeserializeStringFunc()(m.setActualEndDate),
		actualStartDateKey:   internalSerialization.DeserializeStringFunc()(m.setActualStartDate),
		catalogIDKey:         internalSerialization.DeserializeStringFunc()(m.setCatalogId),
		endDateUTCKey:        internalSerialization.DeserializeStringFunc()(m.setEndDateUTC),
		locationKey:          internalSerialization.DeserializeStringFunc()(m.setLocation),
		openedForKey:         internalSerialization.DeserializeStringFunc()(m.setOpenedFor),
		rescheduleKey:        internalSerialization.DeserializeBoolFunc()(m.setReschedule),
		serviceConfigRuleKey: internalSerialization.DeserializeStringFunc()(m.setServiceConfigRule),
		startDateUTCKey:      internalSerialization.DeserializeStringFunc()(m.setStartDateUTC),
		taskIDKey:            internalSerialization.DeserializeStringFunc()(m.setTaskId),
		taskTableKey:         internalSerialization.DeserializeStringFunc()(m.setTaskTable),
		timezoneKey:          internalSerialization.DeserializeStringFunc()(m.setTimezone),
		validateRequestKey:   internalSerialization.DeserializeBoolFunc()(m.setValidateRequest),
	}
}

func (m *AppointmentRequestModel) GetActualEndDate() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, actualEndDateKey)
}
func (m *AppointmentRequestModel) setActualEndDate(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, actualEndDateKey, val)
}
func (m *AppointmentRequestModel) GetActualStartDate() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, actualStartDateKey)
}
func (m *AppointmentRequestModel) setActualStartDate(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, actualStartDateKey, val)
}
func (m *AppointmentRequestModel) GetCatalogId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, catalogIDKey)
}
func (m *AppointmentRequestModel) setCatalogId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, catalogIDKey, val)
}
func (m *AppointmentRequestModel) GetEndDateUTC() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, endDateUTCKey)
}
func (m *AppointmentRequestModel) setEndDateUTC(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, endDateUTCKey, val)
}
func (m *AppointmentRequestModel) GetLocation() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, locationKey)
}
func (m *AppointmentRequestModel) setLocation(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, locationKey, val)
}
func (m *AppointmentRequestModel) GetOpenedFor() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, openedForKey)
}
func (m *AppointmentRequestModel) setOpenedFor(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, openedForKey, val)
}
func (m *AppointmentRequestModel) GetReschedule() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *bool](m, rescheduleKey)
}
func (m *AppointmentRequestModel) setReschedule(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, rescheduleKey, val)
}
func (m *AppointmentRequestModel) GetServiceConfigRule() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, serviceConfigRuleKey)
}
func (m *AppointmentRequestModel) setServiceConfigRule(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, serviceConfigRuleKey, val)
}
func (m *AppointmentRequestModel) GetStartDateUTC() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, startDateUTCKey)
}
func (m *AppointmentRequestModel) setStartDateUTC(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, startDateUTCKey, val)
}
func (m *AppointmentRequestModel) GetTaskId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, taskIDKey)
}
func (m *AppointmentRequestModel) setTaskId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskIDKey, val)
}
func (m *AppointmentRequestModel) GetTaskTable() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, taskTableKey)
}
func (m *AppointmentRequestModel) setTaskTable(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskTableKey, val)
}
func (m *AppointmentRequestModel) GetTimezone() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, timezoneKey)
}
func (m *AppointmentRequestModel) setTimezone(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, timezoneKey, val)
}
func (m *AppointmentRequestModel) GetValidateRequest() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *bool](m, validateRequestKey)
}
func (m *AppointmentRequestModel) setValidateRequest(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, validateRequestKey, val)
}
