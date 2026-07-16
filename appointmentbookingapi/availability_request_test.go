package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestAvailabilityRequestModel_GettersSetters(t *testing.T) {
	model := NewAvailabilityRequest()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"CatalogId", func(v interface{}) error { return model.setCatalogId(v.(*string)) }, func() (interface{}, error) { return model.GetCatalogId() }, internal.ToPointer("val")},
		{"EndDate", func(v interface{}) error { return model.setEndDate(v.(*string)) }, func() (interface{}, error) { return model.GetEndDate() }, internal.ToPointer("val")},
		{"FetchDaysSlot", func(v interface{}) error { return model.setFetchDaysSlot(v.(*bool)) }, func() (interface{}, error) { return model.GetFetchDaysSlot() }, internal.ToPointer(true)},
		{"FullDay", func(v interface{}) error { return model.setFullDay(v.(*bool)) }, func() (interface{}, error) { return model.GetFullDay() }, internal.ToPointer(true)},
		{"GetNextAvailableSlot", func(v interface{}) error { return model.setGetNextAvailableSlot(v.(*bool)) }, func() (interface{}, error) { return model.GetGetNextAvailableSlot() }, internal.ToPointer(true)},
		{"Limit", func(v interface{}) error { return model.setLimit(v.(*int32)) }, func() (interface{}, error) { return model.GetLimit() }, internal.ToPointer(int32(10))},
		{"Location", func(v interface{}) error { return model.setLocation(v.(*string)) }, func() (interface{}, error) { return model.GetLocation() }, internal.ToPointer("val")},
		{"OpenedFor", func(v interface{}) error { return model.setOpenedFor(v.(*string)) }, func() (interface{}, error) { return model.GetOpenedFor() }, internal.ToPointer("val")},
		{"OtherInputs", func(v interface{}) error { return model.setOtherInputs(v) }, func() (interface{}, error) { return model.GetOtherInputs() }, "val"},
		{"ServiceConfigRule", func(v interface{}) error { return model.setServiceConfigRule(v.(*string)) }, func() (interface{}, error) { return model.GetServiceConfigRule() }, internal.ToPointer("val")},
		{"StartDate", func(v interface{}) error { return model.setStartDate(v.(*string)) }, func() (interface{}, error) { return model.GetStartDate() }, internal.ToPointer("val")},
		{"TaskId", func(v interface{}) error { return model.setTaskId(v.(*string)) }, func() (interface{}, error) { return model.GetTaskId() }, internal.ToPointer("val")},
		{"TaskTable", func(v interface{}) error { return model.setTaskTable(v.(*string)) }, func() (interface{}, error) { return model.GetTaskTable() }, internal.ToPointer("val")},
		{"UseReadReplica", func(v interface{}) error { return model.setUseReadReplica(v.(*bool)) }, func() (interface{}, error) { return model.GetUseReadReplica() }, internal.ToPointer(true)},
		{"View", func(v interface{}) error { return model.setView(v.(*string)) }, func() (interface{}, error) { return model.GetView() }, internal.ToPointer("val")},
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

func TestCreateAvailabilityRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateAvailabilityRequestFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}
