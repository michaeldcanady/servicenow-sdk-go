package appointmentbookingapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppointmentRequestModel_GettersSetters(t *testing.T) {
	model := NewAppointmentRequest()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"ActualEndDate", func(v interface{}) error { return model.setActualEndDate(v.(*string)) }, func() (interface{}, error) { return model.GetActualEndDate() }, ptr("val")},
		{"ActualStartDate", func(v interface{}) error { return model.setActualStartDate(v.(*string)) }, func() (interface{}, error) { return model.GetActualStartDate() }, ptr("val")},
		{"CatalogId", func(v interface{}) error { return model.setCatalogId(v.(*string)) }, func() (interface{}, error) { return model.GetCatalogId() }, ptr("val")},
		{"EndDateUTC", func(v interface{}) error { return model.setEndDateUTC(v.(*string)) }, func() (interface{}, error) { return model.GetEndDateUTC() }, ptr("val")},
		{"Location", func(v interface{}) error { return model.setLocation(v.(*string)) }, func() (interface{}, error) { return model.GetLocation() }, ptr("val")},
		{"OpenedFor", func(v interface{}) error { return model.setOpenedFor(v.(*string)) }, func() (interface{}, error) { return model.GetOpenedFor() }, ptr("val")},
		{"Reschedule", func(v interface{}) error { return model.setReschedule(v.(*bool)) }, func() (interface{}, error) { return model.GetReschedule() }, ptr(true)},
		{"ServiceConfigRule", func(v interface{}) error { return model.setServiceConfigRule(v.(*string)) }, func() (interface{}, error) { return model.GetServiceConfigRule() }, ptr("val")},
		{"StartDateUTC", func(v interface{}) error { return model.setStartDateUTC(v.(*string)) }, func() (interface{}, error) { return model.GetStartDateUTC() }, ptr("val")},
		{"TaskId", func(v interface{}) error { return model.setTaskId(v.(*string)) }, func() (interface{}, error) { return model.GetTaskId() }, ptr("val")},
		{"TaskTable", func(v interface{}) error { return model.setTaskTable(v.(*string)) }, func() (interface{}, error) { return model.GetTaskTable() }, ptr("val")},
		{"Timezone", func(v interface{}) error { return model.setTimezone(v.(*string)) }, func() (interface{}, error) { return model.GetTimezone() }, ptr("val")},
		{"ValidateRequest", func(v interface{}) error { return model.setValidateRequest(v.(*bool)) }, func() (interface{}, error) { return model.GetValidateRequest() }, ptr(true)},
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

func TestCreateAppointmentRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateAppointmentRequestFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}
