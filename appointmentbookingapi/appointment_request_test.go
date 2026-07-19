package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
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
		{"ActualEndDate", func(v interface{}) error { return model.SetActualEndDate(v.(*string)) }, func() (interface{}, error) { return model.GetActualEndDate() }, internal.ToPointer("val")},
		{"ActualStartDate", func(v interface{}) error { return model.SetActualStartDate(v.(*string)) }, func() (interface{}, error) { return model.GetActualStartDate() }, internal.ToPointer("val")},
		{"CatalogId", func(v interface{}) error { return model.SetCatalogId(v.(*string)) }, func() (interface{}, error) { return model.GetCatalogId() }, internal.ToPointer("val")},
		{"EndDateUTC", func(v interface{}) error { return model.SetEndDateUTC(v.(*string)) }, func() (interface{}, error) { return model.GetEndDateUTC() }, internal.ToPointer("val")},
		{"Location", func(v interface{}) error { return model.SetLocation(v.(*string)) }, func() (interface{}, error) { return model.GetLocation() }, internal.ToPointer("val")},
		{"OpenedFor", func(v interface{}) error { return model.SetOpenedFor(v.(*string)) }, func() (interface{}, error) { return model.GetOpenedFor() }, internal.ToPointer("val")},
		{"Reschedule", func(v interface{}) error { return model.SetReschedule(v.(*bool)) }, func() (interface{}, error) { return model.GetReschedule() }, internal.ToPointer(true)},
		{"ServiceConfigRule", func(v interface{}) error { return model.SetServiceConfigRule(v.(*string)) }, func() (interface{}, error) { return model.GetServiceConfigRule() }, internal.ToPointer("val")},
		{"StartDateUTC", func(v interface{}) error { return model.SetStartDateUTC(v.(*string)) }, func() (interface{}, error) { return model.GetStartDateUTC() }, internal.ToPointer("val")},
		{"TaskId", func(v interface{}) error { return model.SetTaskId(v.(*string)) }, func() (interface{}, error) { return model.GetTaskId() }, internal.ToPointer("val")},
		{"TaskTable", func(v interface{}) error { return model.SetTaskTable(v.(*string)) }, func() (interface{}, error) { return model.GetTaskTable() }, internal.ToPointer("val")},
		{"Timezone", func(v interface{}) error { return model.SetTimezone(v.(*string)) }, func() (interface{}, error) { return model.GetTimezone() }, internal.ToPointer("val")},
		{"ValidateRequest", func(v interface{}) error { return model.SetValidateRequest(v.(*bool)) }, func() (interface{}, error) { return model.GetValidateRequest() }, internal.ToPointer(true)},
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
