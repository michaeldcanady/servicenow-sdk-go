package appointmentbookingapi

import (
	"testing"

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
		{"CatalogId", func(v interface{}) error { return model.setCatalogId(v.(*string)) }, func() (interface{}, error) { return model.GetCatalogId() }, ptr("val")},
		{"EndDate", func(v interface{}) error { return model.setEndDate(v.(*string)) }, func() (interface{}, error) { return model.GetEndDate() }, ptr("val")},
		{"FetchDaysSlot", func(v interface{}) error { return model.setFetchDaysSlot(v.(*bool)) }, func() (interface{}, error) { return model.GetFetchDaysSlot() }, ptr(true)},
		{"FullDay", func(v interface{}) error { return model.setFullDay(v.(*bool)) }, func() (interface{}, error) { return model.GetFullDay() }, ptr(true)},
		{"GetNextAvailableSlot", func(v interface{}) error { return model.setGetNextAvailableSlot(v.(*bool)) }, func() (interface{}, error) { return model.GetGetNextAvailableSlot() }, ptr(true)},
		{"Limit", func(v interface{}) error { return model.setLimit(v.(*int32)) }, func() (interface{}, error) { return model.GetLimit() }, ptr(int32(10))},
		{"Location", func(v interface{}) error { return model.setLocation(v.(*string)) }, func() (interface{}, error) { return model.GetLocation() }, ptr("val")},
		{"OpenedFor", func(v interface{}) error { return model.setOpenedFor(v.(*string)) }, func() (interface{}, error) { return model.GetOpenedFor() }, ptr("val")},
		{"OtherInputs", func(v interface{}) error { return model.setOtherInputs(v) }, func() (interface{}, error) { return model.GetOtherInputs() }, "val"},
		{"ServiceConfigRule", func(v interface{}) error { return model.setServiceConfigRule(v.(*string)) }, func() (interface{}, error) { return model.GetServiceConfigRule() }, ptr("val")},
		{"StartDate", func(v interface{}) error { return model.setStartDate(v.(*string)) }, func() (interface{}, error) { return model.GetStartDate() }, ptr("val")},
		{"TaskId", func(v interface{}) error { return model.setTaskId(v.(*string)) }, func() (interface{}, error) { return model.GetTaskId() }, ptr("val")},
		{"TaskTable", func(v interface{}) error { return model.setTaskTable(v.(*string)) }, func() (interface{}, error) { return model.GetTaskTable() }, ptr("val")},
		{"UseReadReplica", func(v interface{}) error { return model.setUseReadReplica(v.(*bool)) }, func() (interface{}, error) { return model.GetUseReadReplica() }, ptr(true)},
		{"View", func(v interface{}) error { return model.setView(v.(*string)) }, func() (interface{}, error) { return model.GetView() }, ptr("val")},
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
