package appointmentbookingapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppointmentResultModel_GettersSetters(t *testing.T) {
	model := NewAppointmentResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"Data", func(v interface{}) error { return model.setData(v.(*string)) }, func() (interface{}, error) { return model.GetData() }, ptr("val")},
		{"Message", func(v interface{}) error { return model.setMessage(v.(*string)) }, func() (interface{}, error) { return model.GetMessage() }, ptr("val")},
		{"Reason", func(v interface{}) error { return model.setReason(v.(*string)) }, func() (interface{}, error) { return model.GetReason() }, ptr("val")},
		{"Success", func(v interface{}) error { return model.setSuccess(v.(*bool)) }, func() (interface{}, error) { return model.GetSuccess() }, ptr(true)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestCreateAppointmentResponseFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateAppointmentResponseFromDiscriminatorValue(nil)
	assert.NoError(t, err)
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
		{"CatalogId", func(v interface{}) error { return model.setCatalogId(v.(*string)) }, func() (interface{}, error) { return model.GetCatalogId() }, ptr("val")},
		{"OtherInputs", func(v interface{}) error { return model.setOtherInputs(v) }, func() (interface{}, error) { return model.GetOtherInputs() }, "val"},
		{"TaskId", func(v interface{}) error { return model.setTaskId(v.(*string)) }, func() (interface{}, error) { return model.GetTaskId() }, ptr("val")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
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
		{"DedicatedCapacity", func(v interface{}) error { return model.setDedicatedCapacity(v.(*bool)) }, func() (interface{}, error) { return model.GetDedicatedCapacity() }, ptr(true)},
		{"FutureMaxBookableDays", func(v interface{}) error { return model.setFutureMaxBookableDays(v.(*string)) }, func() (interface{}, error) { return model.GetFutureMaxBookableDays() }, ptr("val")},
		{"RuleId", func(v interface{}) error { return model.setRuleId(v.(*string)) }, func() (interface{}, error) { return model.GetRuleId() }, ptr("val")},
		{"RuleName", func(v interface{}) error { return model.setRuleName(v.(*string)) }, func() (interface{}, error) { return model.GetRuleName() }, ptr("val")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
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
		{"Availability", func(v interface{}) error { return model.setAvailability(v.([]AvailabilitySlot)) }, func() (interface{}, error) { return model.GetAvailability() }, []AvailabilitySlot{NewAvailabilitySlot()}},
		{"HasMore", func(v interface{}) error { return model.setHasMore(v.(*bool)) }, func() (interface{}, error) { return model.GetHasMore() }, ptr(true)},
		{"NextAvailableSlot", func(v interface{}) error { return model.setNextAvailableSlot(v) }, func() (interface{}, error) { return model.GetNextAvailableSlot() }, "val"},
		{"NoApptAvailable", func(v interface{}) error { return model.setNoApptAvailable(v.(*bool)) }, func() (interface{}, error) { return model.GetNoApptAvailable() }, ptr(true)},
		{"Success", func(v interface{}) error { return model.setSuccess(v.(*bool)) }, func() (interface{}, error) { return model.GetSuccess() }, ptr(true)},
		{"TimeZone", func(v interface{}) error { return model.setTimeZone(v.(*string)) }, func() (interface{}, error) { return model.GetTimeZone() }, ptr("val")},
		{"TimeZoneDisplayValue", func(v interface{}) error { return model.setTimeZoneDisplayValue(v.(*string)) }, func() (interface{}, error) { return model.GetTimeZoneDisplayValue() }, ptr("val")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestCreateAvailabilityResponseFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateAvailabilityResponseFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateAvailabilitySlotFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateAvailabilitySlotFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}

