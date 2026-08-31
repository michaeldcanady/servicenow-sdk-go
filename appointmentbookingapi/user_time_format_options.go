// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi // nolint:dupl // shares field-count shape with RpVariableModel by coincidence, not copy-paste; distinct API concept, not worth sacrificing named accessors for

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// UserTimeFormatOptions represents userTimeFormatOptions nested object.
type UserTimeFormatOptions interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetHour() (*int8, error)
	SetHour(*int8) error
	GetHourCycle() (*string, error)
	SetHourCycle(*string) error
	GetMinute() (*int8, error)
	SetMinute(*int8) error
}

// UserTimeFormatOptionsModel represents the user time format options model.
type UserTimeFormatOptionsModel struct {
	core.BaseModel
}

// NewUserTimeFormatOptions creates a new instance of UserTimeFormatOptionsModel.
func NewUserTimeFormatOptions() *UserTimeFormatOptionsModel {
	return &UserTimeFormatOptionsModel{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateUserTimeFormatOptionsFromDiscriminatorValue creates a new UserTimeFormatOptions from a ParseNode.
func CreateUserTimeFormatOptionsFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUserTimeFormatOptions(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *UserTimeFormatOptionsModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeInt8Func(hourKey, m.GetHour),
		internalSerialization.SerializeStringFunc(hourCycleKey, m.GetHourCycle),
		internalSerialization.SerializeInt8Func(minuteKey, m.GetMinute),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *UserTimeFormatOptionsModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		hourKey:      internalSerialization.DeserializeInt8Func(m.SetHour),
		hourCycleKey: internalSerialization.DeserializeStringFunc(m.SetHourCycle),
		minuteKey:    internalSerialization.DeserializeInt8Func(m.SetMinute),
	}
}

// GetHour returns the hour value.
func (m *UserTimeFormatOptionsModel) GetHour() (*int8, error) {
	return store.DefaultBackedModelAccessorFunc[*UserTimeFormatOptionsModel, *int8](m, hourKey)
}

// SetHour sets the hour value.
func (m *UserTimeFormatOptionsModel) SetHour(val *int8) error {
	return store.DefaultBackedModelMutatorFunc(m, hourKey, val)
}

// GetHourCycle returns the hour cycle value.
func (m *UserTimeFormatOptionsModel) GetHourCycle() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserTimeFormatOptionsModel, *string](m, hourCycleKey)
}

// SetHourCycle sets the hour cycle value.
func (m *UserTimeFormatOptionsModel) SetHourCycle(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, hourCycleKey, val)
}

// GetMinute returns the minute value.
func (m *UserTimeFormatOptionsModel) GetMinute() (*int8, error) {
	return store.DefaultBackedModelAccessorFunc[*UserTimeFormatOptionsModel, *int8](m, minuteKey)
}

// SetMinute sets the minute value.
func (m *UserTimeFormatOptionsModel) SetMinute(val *int8) error {
	return store.DefaultBackedModelMutatorFunc(m, minuteKey, val)
}
