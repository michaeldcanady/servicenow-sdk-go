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
	SetActualEndDate(*string) error
	GetActualStartDate() (*string, error)
	SetActualStartDate(*string) error
	GetCatalogId() (*string, error)
	SetCatalogId(*string) error
	GetEndDateUTC() (*string, error)
	SetEndDateUTC(*string) error
	GetLocation() (*string, error)
	SetLocation(*string) error
	GetOpenedFor() (*string, error)
	SetOpenedFor(*string) error
	GetReschedule() (*bool, error)
	SetReschedule(*bool) error
	GetServiceConfigRule() (*string, error)
	SetServiceConfigRule(*string) error
	GetStartDateUTC() (*string, error)
	SetStartDateUTC(*string) error
	GetTaskId() (*string, error)
	SetTaskId(*string) error
	GetTaskTable() (*string, error)
	SetTaskTable(*string) error
	GetTimezone() (*string, error)
	SetTimezone(*string) error
	GetValidateRequest() (*bool, error)
	SetValidateRequest(*bool) error
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
		internalSerialization.SerializeStringFunc(actualEndDateKey, m.GetActualEndDate),
		internalSerialization.SerializeStringFunc(actualStartDateKey, m.GetActualStartDate),
		internalSerialization.SerializeStringFunc(catalogIDKey, m.GetCatalogId),
		internalSerialization.SerializeStringFunc(endDateUTCKey, m.GetEndDateUTC),
		internalSerialization.SerializeStringFunc(locationKey, m.GetLocation),
		internalSerialization.SerializeStringFunc(openedForKey, m.GetOpenedFor),
		internalSerialization.SerializeBoolFunc(rescheduleKey, m.GetReschedule),
		internalSerialization.SerializeStringFunc(serviceConfigRuleKey, m.GetServiceConfigRule),
		internalSerialization.SerializeStringFunc(startDateUTCKey, m.GetStartDateUTC),
		internalSerialization.SerializeStringFunc(taskIDKey, m.GetTaskId),
		internalSerialization.SerializeStringFunc(taskTableKey, m.GetTaskTable),
		internalSerialization.SerializeStringFunc(timezoneKey, m.GetTimezone),
		internalSerialization.SerializeBoolFunc(validateRequestKey, m.GetValidateRequest),
	)
}

func (m *AppointmentRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		actualEndDateKey:     internalSerialization.DeserializeStringFunc(m.SetActualEndDate),
		actualStartDateKey:   internalSerialization.DeserializeStringFunc(m.SetActualStartDate),
		catalogIDKey:         internalSerialization.DeserializeStringFunc(m.SetCatalogId),
		endDateUTCKey:        internalSerialization.DeserializeStringFunc(m.SetEndDateUTC),
		locationKey:          internalSerialization.DeserializeStringFunc(m.SetLocation),
		openedForKey:         internalSerialization.DeserializeStringFunc(m.SetOpenedFor),
		rescheduleKey:        internalSerialization.DeserializeBoolFunc(m.SetReschedule),
		serviceConfigRuleKey: internalSerialization.DeserializeStringFunc(m.SetServiceConfigRule),
		startDateUTCKey:      internalSerialization.DeserializeStringFunc(m.SetStartDateUTC),
		taskIDKey:            internalSerialization.DeserializeStringFunc(m.SetTaskId),
		taskTableKey:         internalSerialization.DeserializeStringFunc(m.SetTaskTable),
		timezoneKey:          internalSerialization.DeserializeStringFunc(m.SetTimezone),
		validateRequestKey:   internalSerialization.DeserializeBoolFunc(m.SetValidateRequest),
	}
}

func (m *AppointmentRequestModel) GetActualEndDate() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, actualEndDateKey)
}
func (m *AppointmentRequestModel) SetActualEndDate(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, actualEndDateKey, val)
}
func (m *AppointmentRequestModel) GetActualStartDate() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, actualStartDateKey)
}
func (m *AppointmentRequestModel) SetActualStartDate(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, actualStartDateKey, val)
}
func (m *AppointmentRequestModel) GetCatalogId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, catalogIDKey)
}
func (m *AppointmentRequestModel) SetCatalogId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, catalogIDKey, val)
}
func (m *AppointmentRequestModel) GetEndDateUTC() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, endDateUTCKey)
}
func (m *AppointmentRequestModel) SetEndDateUTC(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, endDateUTCKey, val)
}
func (m *AppointmentRequestModel) GetLocation() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, locationKey)
}
func (m *AppointmentRequestModel) SetLocation(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, locationKey, val)
}
func (m *AppointmentRequestModel) GetOpenedFor() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, openedForKey)
}
func (m *AppointmentRequestModel) SetOpenedFor(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, openedForKey, val)
}
func (m *AppointmentRequestModel) GetReschedule() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *bool](m, rescheduleKey)
}
func (m *AppointmentRequestModel) SetReschedule(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, rescheduleKey, val)
}
func (m *AppointmentRequestModel) GetServiceConfigRule() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, serviceConfigRuleKey)
}
func (m *AppointmentRequestModel) SetServiceConfigRule(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, serviceConfigRuleKey, val)
}
func (m *AppointmentRequestModel) GetStartDateUTC() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, startDateUTCKey)
}
func (m *AppointmentRequestModel) SetStartDateUTC(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, startDateUTCKey, val)
}
func (m *AppointmentRequestModel) GetTaskId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, taskIDKey)
}
func (m *AppointmentRequestModel) SetTaskId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskIDKey, val)
}
func (m *AppointmentRequestModel) GetTaskTable() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, taskTableKey)
}
func (m *AppointmentRequestModel) SetTaskTable(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskTableKey, val)
}
func (m *AppointmentRequestModel) GetTimezone() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, timezoneKey)
}
func (m *AppointmentRequestModel) SetTimezone(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, timezoneKey, val)
}
func (m *AppointmentRequestModel) GetValidateRequest() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *bool](m, validateRequestKey)
}
func (m *AppointmentRequestModel) SetValidateRequest(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, validateRequestKey, val)
}
