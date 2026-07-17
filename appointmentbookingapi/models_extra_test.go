package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
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
		{"Data", func(v interface{}) error { return model.setData(v.(*string)) }, func() (interface{}, error) { return model.GetData() }, internal.ToPointer("val")},
		{"Message", func(v interface{}) error { return model.setMessage(v.(*string)) }, func() (interface{}, error) { return model.GetMessage() }, internal.ToPointer("val")},
		{"Reason", func(v interface{}) error { return model.setReason(v.(*string)) }, func() (interface{}, error) { return model.GetReason() }, internal.ToPointer("val")},
		{"Success", func(v interface{}) error { return model.setSuccess(v.(*bool)) }, func() (interface{}, error) { return model.GetSuccess() }, internal.ToPointer(true)},
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
		{"CatalogId", func(v interface{}) error { return model.setCatalogId(v.(*string)) }, func() (interface{}, error) { return model.GetCatalogId() }, internal.ToPointer("val")},
		{"OtherInputs", func(v interface{}) error { return model.setOtherInputs(v) }, func() (interface{}, error) { return model.GetOtherInputs() }, "val"},
		{"TaskId", func(v interface{}) error { return model.setTaskId(v.(*string)) }, func() (interface{}, error) { return model.GetTaskId() }, internal.ToPointer("val")},
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
		{"DedicatedCapacity", func(v interface{}) error { return model.setDedicatedCapacity(v.(*bool)) }, func() (interface{}, error) { return model.GetDedicatedCapacity() }, internal.ToPointer(true)},
		{"FutureMaxBookableDays", func(v interface{}) error { return model.setFutureMaxBookableDays(v.(*string)) }, func() (interface{}, error) { return model.GetFutureMaxBookableDays() }, internal.ToPointer("val")},
		{"RuleId", func(v interface{}) error { return model.setRuleId(v.(*string)) }, func() (interface{}, error) { return model.GetRuleId() }, internal.ToPointer("val")},
		{"RuleName", func(v interface{}) error { return model.setRuleName(v.(*string)) }, func() (interface{}, error) { return model.GetRuleName() }, internal.ToPointer("val")},
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
		{"HasMore", func(v interface{}) error { return model.setHasMore(v.(*bool)) }, func() (interface{}, error) { return model.GetHasMore() }, internal.ToPointer(true)},
		{"NextAvailableSlot", func(v interface{}) error { return model.setNextAvailableSlot(v) }, func() (interface{}, error) { return model.GetNextAvailableSlot() }, "val"},
		{"NoApptAvailable", func(v interface{}) error { return model.setNoApptAvailable(v.(*bool)) }, func() (interface{}, error) { return model.GetNoApptAvailable() }, internal.ToPointer(true)},
		{"Success", func(v interface{}) error { return model.setSuccess(v.(*bool)) }, func() (interface{}, error) { return model.GetSuccess() }, internal.ToPointer(true)},
		{"TimeZone", func(v interface{}) error { return model.setTimeZone(v.(*string)) }, func() (interface{}, error) { return model.GetTimeZone() }, internal.ToPointer("val")},
		{"TimeZoneDisplayValue", func(v interface{}) error { return model.setTimeZoneDisplayValue(v.(*string)) }, func() (interface{}, error) { return model.GetTimeZoneDisplayValue() }, internal.ToPointer("val")},
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

func TestAvailabilitySlotModel_GetAdditionalData_Empty(t *testing.T) {
	model := NewAvailabilitySlot()

	got := model.GetAdditionalData()

	assert.Equal(t, map[string]interface{}{}, got)
}

func TestAvailabilitySlotModel_GetSetAdditionalData(t *testing.T) {
	model := NewAvailabilitySlot()
	want := map[string]interface{}{"start": "09:00", "end": "10:00"}

	model.SetAdditionalData(want)
	got := model.GetAdditionalData()

	assert.Equal(t, want, got)
}

func TestAvailabilitySlotModel_GetAdditionalData_DoesNotWriteBackOnEmptyRead(t *testing.T) {
	model := NewAvailabilitySlot()

	_ = model.GetAdditionalData()
	val, err := model.GetBackingStore().Get(additionalDataKey)

	assert.NoError(t, err)
	assert.Nil(t, val)
}
