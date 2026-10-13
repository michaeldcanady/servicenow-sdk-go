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

// AvailabilityResponse represents the availability response.
type AvailabilityResponse = core.ServiceNowItemResponse[*AvailabilityResultModel]

// CreateAvailabilityResponseFromDiscriminatorValue is a factory for creating an AvailabilityResponse.
func CreateAvailabilityResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*AvailabilityResultModel](CreateAvailabilityResultFromDiscriminatorValue), nil
}

// AvailabilityResult represents the result object in availability response.
type AvailabilityResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetAvailability() ([]AvailabilitySlot, error)
	SetAvailability([]AvailabilitySlot) error
	GetHasMore() (*bool, error)
	SetHasMore(*bool) error
	GetNextAvailableSlot() (AvailabilitySlot, error)
	SetNextAvailableSlot(AvailabilitySlot) error
	GetNoApptAvailable() (*bool, error)
	SetNoApptAvailable(*bool) error
	GetSuccess() (*bool, error)
	SetSuccess(*bool) error
	GetTimeZone() (*string, error)
	SetTimeZone(*string) error
	GetTimeZoneDisplayValue() (*string, error)
	SetTimeZoneDisplayValue(*string) error
}

// AvailabilityResultModel represents the availability result model.
type AvailabilityResultModel struct {
	core.BaseModel
}

// NewAvailabilityResult creates a new instance of AvailabilityResultModel.
func NewAvailabilityResult() *AvailabilityResultModel {
	return &AvailabilityResultModel{BaseModel: *core.NewBaseModel()}
}

// CreateAvailabilityResultFromDiscriminatorValue creates a new AvailabilityResult from a ParseNode.
func CreateAvailabilityResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAvailabilityResult(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *AvailabilityResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeCollectionOfObjectValuesFunc[AvailabilitySlot](availabilityKey, m.GetAvailability),
		internalSerialization.SerializeBoolFunc(hasMoreKey, m.GetHasMore),
		internalSerialization.SerializeObjectValueFunc[AvailabilitySlot](nextAvailableSlotKey, m.GetNextAvailableSlot),
		internalSerialization.SerializeBoolFunc(noApptAvailableKey, m.GetNoApptAvailable),
		internalSerialization.SerializeBoolFunc(successKey, m.GetSuccess),
		internalSerialization.SerializeStringFunc(timeZoneKey, m.GetTimeZone),
		internalSerialization.SerializeStringFunc(timeZoneDisplayValueKey, m.GetTimeZoneDisplayValue),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *AvailabilityResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		availabilityKey:         internalSerialization.DeserializeCollectionOfObjectValuesFunc[AvailabilitySlot](CreateAvailabilitySlotFromDiscriminatorValue, m.SetAvailability),
		hasMoreKey:              internalSerialization.DeserializeBoolFunc(m.SetHasMore),
		nextAvailableSlotKey:    internalSerialization.DeserializeObjectValueFunc[AvailabilitySlot](CreateAvailabilitySlotFromDiscriminatorValue, m.SetNextAvailableSlot),
		noApptAvailableKey:      internalSerialization.DeserializeBoolFunc(m.SetNoApptAvailable),
		successKey:              internalSerialization.DeserializeBoolFunc(m.SetSuccess),
		timeZoneKey:             internalSerialization.DeserializeStringFunc(m.SetTimeZone),
		timeZoneDisplayValueKey: internalSerialization.DeserializeStringFunc(m.SetTimeZoneDisplayValue),
	}
}

// GetAvailability returns the availability value.
func (m *AvailabilityResultModel) GetAvailability() ([]AvailabilitySlot, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityResultModel, []AvailabilitySlot](m, availabilityKey)
}

// SetAvailability sets the availability value.
func (m *AvailabilityResultModel) SetAvailability(val []AvailabilitySlot) error {
	return store.DefaultBackedModelMutatorFunc(m, availabilityKey, val)
}

// GetHasMore returns the has more value.
func (m *AvailabilityResultModel) GetHasMore() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityResultModel, *bool](m, hasMoreKey)
}

// SetHasMore sets the has more value.
func (m *AvailabilityResultModel) SetHasMore(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, hasMoreKey, val)
}

// GetNextAvailableSlot returns the next available slot value.
func (m *AvailabilityResultModel) GetNextAvailableSlot() (AvailabilitySlot, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityResultModel, AvailabilitySlot](m, nextAvailableSlotKey)
}

// SetNextAvailableSlot sets the next available slot value.
func (m *AvailabilityResultModel) SetNextAvailableSlot(val AvailabilitySlot) error {
	return store.DefaultBackedModelMutatorFunc(m, nextAvailableSlotKey, val)
}

// GetNoApptAvailable returns the no appt available value.
func (m *AvailabilityResultModel) GetNoApptAvailable() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityResultModel, *bool](m, noApptAvailableKey)
}

// SetNoApptAvailable sets the no appt available value.
func (m *AvailabilityResultModel) SetNoApptAvailable(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, noApptAvailableKey, val)
}

// GetSuccess returns the success value.
func (m *AvailabilityResultModel) GetSuccess() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityResultModel, *bool](m, successKey)
}

// SetSuccess sets the success value.
func (m *AvailabilityResultModel) SetSuccess(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, successKey, val)
}

// GetTimeZone returns the time zone value.
func (m *AvailabilityResultModel) GetTimeZone() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityResultModel, *string](m, timeZoneKey)
}

// SetTimeZone sets the time zone value.
func (m *AvailabilityResultModel) SetTimeZone(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, timeZoneKey, val)
}

// GetTimeZoneDisplayValue returns the time zone display value value.
func (m *AvailabilityResultModel) GetTimeZoneDisplayValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityResultModel, *string](m, timeZoneDisplayValueKey)
}

// SetTimeZoneDisplayValue sets the time zone display value value.
func (m *AvailabilityResultModel) SetTimeZoneDisplayValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, timeZoneDisplayValueKey, val)
}
