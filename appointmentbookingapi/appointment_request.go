// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// AppointmentRequest represents the appointment request.
type AppointmentRequest interface {
	serialization.Parsable
	kiotaStore.BackedModel

	// TODO: date, required
	GetActualEndDate() (*string, error)
	SetActualEndDate(*string) error
	// TODO: date, required
	GetActualStartDate() (*string, error)
	SetActualStartDate(*string) error
	GetCatalogID() (*string, error)
	SetCatalogID(*string) error
	// TODO: date, required
	GetEndDateUTC() (*string, error)
	SetEndDateUTC(*string) error
	// required
	GetLocation() (*string, error)
	SetLocation(*string) error
	// required
	GetOpenedFor() (*string, error)
	SetOpenedFor(*string) error

	// required
	GetReschedule() (*bool, error)
	SetReschedule(*bool) error
	GetServiceConfigRule() (*string, error)
	SetServiceConfigRule(*string) error
	// TODO: date, required
	GetStartDateUTC() (*string, error)
	SetStartDateUTC(*string) error
	GetTaskID() (*string, error)
	SetTaskID(*string) error
	// TODO: required
	GetTaskTable() (*string, error)
	SetTaskTable(*string) error
	// TODO: required
	GetTimezone() (*string, error)
	SetTimezone(*string) error
	GetValidateRequest() (*bool, error)
	SetValidateRequest(*bool) error
}

// AppointmentRequestModel represents the appointment request model.
type AppointmentRequestModel struct {
	core.BaseModel
}

// NewAppointmentRequest creates a new instance of AppointmentRequestModel.
func NewAppointmentRequest() *AppointmentRequestModel {
	return &AppointmentRequestModel{BaseModel: *core.NewBaseModel()}
}

// CreateAppointmentRequestFromDiscriminatorValue creates a new AppointmentRequest from a ParseNode.
func CreateAppointmentRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAppointmentRequest(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *AppointmentRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(actualEndDateKey, m.GetActualEndDate),
		internalSerialization.SerializeStringFunc(actualStartDateKey, m.GetActualStartDate),
		internalSerialization.SerializeStringFunc(catalogIDKey, m.GetCatalogID),
		internalSerialization.SerializeStringFunc(endDateUTCKey, m.GetEndDateUTC),
		internalSerialization.SerializeStringFunc(locationKey, m.GetLocation),
		internalSerialization.SerializeStringFunc(openedForKey, m.GetOpenedFor),
		internalSerialization.SerializeBoolFunc(rescheduleKey, m.GetReschedule),
		internalSerialization.SerializeStringFunc(serviceConfigRuleKey, m.GetServiceConfigRule),
		internalSerialization.SerializeStringFunc(startDateUTCKey, m.GetStartDateUTC),
		internalSerialization.SerializeStringFunc(taskIDKey, m.GetTaskID),
		internalSerialization.SerializeStringFunc(taskTableKey, m.GetTaskTable),
		internalSerialization.SerializeStringFunc(timezoneKey, m.GetTimezone),
		internalSerialization.SerializeBoolFunc(validateRequestKey, m.GetValidateRequest),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *AppointmentRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		actualEndDateKey:     internalSerialization.DeserializeStringFunc(m.SetActualEndDate),
		actualStartDateKey:   internalSerialization.DeserializeStringFunc(m.SetActualStartDate),
		catalogIDKey:         internalSerialization.DeserializeStringFunc(m.SetCatalogID),
		endDateUTCKey:        internalSerialization.DeserializeStringFunc(m.SetEndDateUTC),
		locationKey:          internalSerialization.DeserializeStringFunc(m.SetLocation),
		openedForKey:         internalSerialization.DeserializeStringFunc(m.SetOpenedFor),
		rescheduleKey:        internalSerialization.DeserializeBoolFunc(m.SetReschedule),
		serviceConfigRuleKey: internalSerialization.DeserializeStringFunc(m.SetServiceConfigRule),
		startDateUTCKey:      internalSerialization.DeserializeStringFunc(m.SetStartDateUTC),
		taskIDKey:            internalSerialization.DeserializeStringFunc(m.SetTaskID),
		taskTableKey:         internalSerialization.DeserializeStringFunc(m.SetTaskTable),
		timezoneKey:          internalSerialization.DeserializeStringFunc(m.SetTimezone),
		validateRequestKey:   internalSerialization.DeserializeBoolFunc(m.SetValidateRequest),
	}
}

// GetActualEndDate returns the actual end date value.
func (m *AppointmentRequestModel) GetActualEndDate() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, actualEndDateKey)
}

// SetActualEndDate sets the actual end date value.
func (m *AppointmentRequestModel) SetActualEndDate(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, actualEndDateKey, val)
}

// GetActualStartDate returns the actual start date value.
func (m *AppointmentRequestModel) GetActualStartDate() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, actualStartDateKey)
}

// SetActualStartDate sets the actual start date value.
func (m *AppointmentRequestModel) SetActualStartDate(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, actualStartDateKey, val)
}

// GetCatalogID returns the catalog id value.
func (m *AppointmentRequestModel) GetCatalogID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, catalogIDKey)
}

// SetCatalogID sets the catalog id value.
func (m *AppointmentRequestModel) SetCatalogID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, catalogIDKey, val)
}

// GetEndDateUTC returns the end date utc value.
func (m *AppointmentRequestModel) GetEndDateUTC() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, endDateUTCKey)
}

// SetEndDateUTC sets the end date utc value.
func (m *AppointmentRequestModel) SetEndDateUTC(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, endDateUTCKey, val)
}

// GetLocation returns the location value.
func (m *AppointmentRequestModel) GetLocation() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, locationKey)
}

// SetLocation sets the location value.
func (m *AppointmentRequestModel) SetLocation(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, locationKey, val)
}

// GetOpenedFor returns the opened for value.
func (m *AppointmentRequestModel) GetOpenedFor() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, openedForKey)
}

// SetOpenedFor sets the opened for value.
func (m *AppointmentRequestModel) SetOpenedFor(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, openedForKey, val)
}

// GetReschedule returns the reschedule value.
func (m *AppointmentRequestModel) GetReschedule() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *bool](m, rescheduleKey)
}

// SetReschedule sets the reschedule value.
func (m *AppointmentRequestModel) SetReschedule(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, rescheduleKey, val)
}

// GetServiceConfigRule returns the service config rule value.
func (m *AppointmentRequestModel) GetServiceConfigRule() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, serviceConfigRuleKey)
}

// SetServiceConfigRule sets the service config rule value.
func (m *AppointmentRequestModel) SetServiceConfigRule(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, serviceConfigRuleKey, val)
}

// GetStartDateUTC returns the start date utc value.
func (m *AppointmentRequestModel) GetStartDateUTC() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, startDateUTCKey)
}

// SetStartDateUTC sets the start date utc value.
func (m *AppointmentRequestModel) SetStartDateUTC(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, startDateUTCKey, val)
}

// GetTaskID returns the task id value.
func (m *AppointmentRequestModel) GetTaskID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, taskIDKey)
}

// SetTaskID sets the task id value.
func (m *AppointmentRequestModel) SetTaskID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskIDKey, val)
}

// GetTaskTable returns the task table value.
func (m *AppointmentRequestModel) GetTaskTable() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, taskTableKey)
}

// SetTaskTable sets the task table value.
func (m *AppointmentRequestModel) SetTaskTable(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskTableKey, val)
}

// GetTimezone returns the timezone value.
func (m *AppointmentRequestModel) GetTimezone() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *string](m, timezoneKey)
}

// SetTimezone sets the timezone value.
func (m *AppointmentRequestModel) SetTimezone(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, timezoneKey, val)
}

// GetValidateRequest returns the validate request value.
func (m *AppointmentRequestModel) GetValidateRequest() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentRequestModel, *bool](m, validateRequestKey)
}

// SetValidateRequest sets the validate request value.
func (m *AppointmentRequestModel) SetValidateRequest(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, validateRequestKey, val)
}
