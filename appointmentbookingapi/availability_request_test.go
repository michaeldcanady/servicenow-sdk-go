package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAvailabilityRequestModel_GettersSetters(t *testing.T) {
	model := NewAvailabilityRequest()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"CatalogId", func(v interface{}) error { return model.SetCatalogID(v.(*string)) }, func() (interface{}, error) { return model.GetCatalogID() }, internal.ToPointer("val")},
		{"EndDate", func(v interface{}) error { return model.SetEndDate(v.(*string)) }, func() (interface{}, error) { return model.GetEndDate() }, internal.ToPointer("val")},
		{"FetchDaysSlot", func(v interface{}) error { return model.SetFetchDaysSlot(v.(*bool)) }, func() (interface{}, error) { return model.GetFetchDaysSlot() }, internal.ToPointer(true)},
		{"FullDay", func(v interface{}) error { return model.SetFullDay(v.(*bool)) }, func() (interface{}, error) { return model.GetFullDay() }, internal.ToPointer(true)},
		{"GetNextAvailableSlot", func(v interface{}) error { return model.SetGetNextAvailableSlot(v.(*bool)) }, func() (interface{}, error) { return model.GetGetNextAvailableSlot() }, internal.ToPointer(true)},
		{"Limit", func(v interface{}) error { return model.SetLimit(v.(*int32)) }, func() (interface{}, error) { return model.GetLimit() }, internal.ToPointer(int32(10))},
		{"Location", func(v interface{}) error { return model.SetLocation(v.(*string)) }, func() (interface{}, error) { return model.GetLocation() }, internal.ToPointer("val")},
		{"OpenedFor", func(v interface{}) error { return model.SetOpenedFor(v.(*string)) }, func() (interface{}, error) { return model.GetOpenedFor() }, internal.ToPointer("val")},
		{"OtherInputs", func(v interface{}) error { return model.SetOtherInputs(v) }, func() (interface{}, error) { return model.GetOtherInputs() }, "val"},
		{"ServiceConfigRule", func(v interface{}) error { return model.SetServiceConfigRule(v.(*string)) }, func() (interface{}, error) { return model.GetServiceConfigRule() }, internal.ToPointer("val")},
		{"StartDate", func(v interface{}) error { return model.SetStartDate(v.(*string)) }, func() (interface{}, error) { return model.GetStartDate() }, internal.ToPointer("val")},
		{"TaskId", func(v interface{}) error { return model.SetTaskID(v.(*string)) }, func() (interface{}, error) { return model.GetTaskID() }, internal.ToPointer("val")},
		{"TaskTable", func(v interface{}) error { return model.SetTaskTable(v.(*string)) }, func() (interface{}, error) { return model.GetTaskTable() }, internal.ToPointer("val")},
		{"UseReadReplica", func(v interface{}) error { return model.SetUseReadReplica(v.(*bool)) }, func() (interface{}, error) { return model.GetUseReadReplica() }, internal.ToPointer(true)},
		{"View", func(v interface{}) error { return model.SetView(v.(*string)) }, func() (interface{}, error) { return model.GetView() }, internal.ToPointer("val")},
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

func TestCreateAvailabilityRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateAvailabilityRequestFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
