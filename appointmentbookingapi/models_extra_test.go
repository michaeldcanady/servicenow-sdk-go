// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppointmentResultModel_GettersSetters(t *testing.T) {
	model := NewAppointmentResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"Data", func(v interface{}) error { return model.SetData(v.(*string)) }, func() (interface{}, error) { return model.GetData() }, internal.ToPointer("val")},
		{"Message", func(v interface{}) error { return model.SetMessage(v.(*string)) }, func() (interface{}, error) { return model.GetMessage() }, internal.ToPointer("val")},
		{"Reason", func(v interface{}) error { return model.SetReason(v.(*string)) }, func() (interface{}, error) { return model.GetReason() }, internal.ToPointer("val")},
		{"Success", func(v interface{}) error { return model.SetSuccess(v.(*bool)) }, func() (interface{}, error) { return model.GetSuccess() }, internal.ToPointer(true)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			require.NoError(t, err)
			got, err := tt.getter()
			require.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestCreateAppointmentResponseFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateAppointmentResponseFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestExecuteRuleConditionsRequestModel_GettersSetters(t *testing.T) {
	model := NewExecuteRuleConditionsRequest()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"CatalogId", func(v interface{}) error { return model.SetCatalogID(v.(*string)) }, func() (interface{}, error) { return model.GetCatalogID() }, internal.ToPointer("val")},
		{"OtherInputs", func(v interface{}) error { return model.SetOtherInputs(v) }, func() (interface{}, error) { return model.GetOtherInputs() }, "val"},
		{"TaskId", func(v interface{}) error { return model.SetTaskID(v.(*string)) }, func() (interface{}, error) { return model.GetTaskID() }, internal.ToPointer("val")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			require.NoError(t, err)
			got, err := tt.getter()
			require.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestExecuteRuleConditionsResultModel_GettersSetters(t *testing.T) {
	model := NewExecuteRuleConditionsResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"DedicatedCapacity", func(v interface{}) error { return model.SetDedicatedCapacity(v.(*bool)) }, func() (interface{}, error) { return model.GetDedicatedCapacity() }, internal.ToPointer(true)},
		{"FutureMaxBookableDays", func(v interface{}) error { return model.SetFutureMaxBookableDays(v.(*string)) }, func() (interface{}, error) { return model.GetFutureMaxBookableDays() }, internal.ToPointer("val")},
		{"RuleId", func(v interface{}) error { return model.SetRuleID(v.(*string)) }, func() (interface{}, error) { return model.GetRuleID() }, internal.ToPointer("val")},
		{"RuleName", func(v interface{}) error { return model.SetRuleName(v.(*string)) }, func() (interface{}, error) { return model.GetRuleName() }, internal.ToPointer("val")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			require.NoError(t, err)
			got, err := tt.getter()
			require.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestAvailabilityResultModel_GettersSetters(t *testing.T) {
	model := NewAvailabilityResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"Availability", func(v interface{}) error { return model.SetAvailability(v.([]AvailabilitySlot)) }, func() (interface{}, error) { return model.GetAvailability() }, []AvailabilitySlot{NewAvailabilitySlot()}},
		{"HasMore", func(v interface{}) error { return model.SetHasMore(v.(*bool)) }, func() (interface{}, error) { return model.GetHasMore() }, internal.ToPointer(true)},
		{"NextAvailableSlot", func(v interface{}) error { return model.SetNextAvailableSlot(v.(AvailabilitySlot)) }, func() (interface{}, error) { return model.GetNextAvailableSlot() }, AvailabilitySlot(NewAvailabilitySlot())},
		{"NoApptAvailable", func(v interface{}) error { return model.SetNoApptAvailable(v.(*bool)) }, func() (interface{}, error) { return model.GetNoApptAvailable() }, internal.ToPointer(true)},
		{"Success", func(v interface{}) error { return model.SetSuccess(v.(*bool)) }, func() (interface{}, error) { return model.GetSuccess() }, internal.ToPointer(true)},
		{"TimeZone", func(v interface{}) error { return model.SetTimeZone(v.(*string)) }, func() (interface{}, error) { return model.GetTimeZone() }, internal.ToPointer("val")},
		{"TimeZoneDisplayValue", func(v interface{}) error { return model.SetTimeZoneDisplayValue(v.(*string)) }, func() (interface{}, error) { return model.GetTimeZoneDisplayValue() }, internal.ToPointer("val")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			require.NoError(t, err)
			got, err := tt.getter()
			require.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestCreateAvailabilityResponseFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateAvailabilityResponseFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateAvailabilitySlotFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateAvailabilitySlotFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestAvailabilitySlotModel_GettersSetters(t *testing.T) {
	model := NewAvailabilitySlot()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Available", func(v any) error { return model.SetAvailable(v.(*bool)) }, func() (any, error) { return model.GetAvailable() }, internal.ToPointer(true)},
		{"EndDate", func(v any) error { return model.SetEndDate(v.(*string)) }, func() (any, error) { return model.GetEndDate() }, internal.ToPointer("2026-07-29 10:00:00")},
		{"EndDateDisplay", func(v any) error { return model.SetEndDateDisplay(v.(*string)) }, func() (any, error) { return model.GetEndDateDisplay() }, internal.ToPointer("10:00 AM")},
		{"EndDateUTC", func(v any) error { return model.SetEndDateUTC(v.(*string)) }, func() (any, error) { return model.GetEndDateUTC() }, internal.ToPointer("2026-07-29 17:00:00")},
		{"StartDate", func(v any) error { return model.SetStartDate(v.(*string)) }, func() (any, error) { return model.GetStartDate() }, internal.ToPointer("2026-07-29 09:00:00")},
		{"StartDateDisplay", func(v any) error { return model.SetStartDateDisplay(v.(*string)) }, func() (any, error) { return model.GetStartDateDisplay() }, internal.ToPointer("9:00 AM")},
		{"StartDateUTC", func(v any) error { return model.SetStartDateUTC(v.(*string)) }, func() (any, error) { return model.GetStartDateUTC() }, internal.ToPointer("2026-07-29 16:00:00")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			require.NoError(t, err)
			got, err := tt.getter()
			require.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestAvailabilitySlotModel_GettersOnNilModel(t *testing.T) {
	var model *AvailabilitySlotModel

	got, err := model.GetAvailable()

	require.Error(t, err)
	assert.Nil(t, got)
}
