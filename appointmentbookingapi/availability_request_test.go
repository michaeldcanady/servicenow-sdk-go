package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		{"View", func(v interface{}) error { return model.SetView(v.(*View)) }, func() (interface{}, error) { return model.GetView() }, internal.ToPointer(ViewPortal)},
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

func TestAvailabilityRequestModel_GetFieldDeserializers(t *testing.T) {
	model := NewAvailabilityRequest()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{
		catalogIDKey, endDateKey, fetchDaysSlotKey, fullDayKey, getNextAvailableSlotKey,
		limitKey, locationKey, openedForKey, otherInputsKey, serviceConfigRuleKey,
		startDateKey, taskIDKey, taskTableKey, useReadReplicaKey, viewKey,
	} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 15)
}

func TestAvailabilityRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *AvailabilityRequestModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			// SerializeAnyFunc has no nil-skip, so an otherwise-empty model still emits
			// the otherInputs key with a nil value.
			name:  "empty model still writes otherInputs",
			model: NewAvailabilityRequest(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteAnyValue", otherInputsKey, nil).Return(nil)
			},
		},
		{
			name: "happy path - writes all fields",
			model: func() *AvailabilityRequestModel {
				m := NewAvailabilityRequest()
				_ = m.SetCatalogID(internal.ToPointer("catalog-id"))
				_ = m.SetEndDate(internal.ToPointer("2026-08-01"))
				_ = m.SetFetchDaysSlot(internal.ToPointer(true))
				_ = m.SetFullDay(internal.ToPointer(false))
				_ = m.SetGetNextAvailableSlot(internal.ToPointer(true))
				_ = m.SetLimit(internal.ToPointer(int32(10)))
				_ = m.SetLocation(internal.ToPointer("location"))
				_ = m.SetOpenedFor(internal.ToPointer("user"))
				_ = m.SetOtherInputs(map[string]any{"key": "value"})
				_ = m.SetServiceConfigRule(internal.ToPointer("rule"))
				_ = m.SetStartDate(internal.ToPointer("2026-07-29"))
				_ = m.SetTaskID(internal.ToPointer("task-id"))
				_ = m.SetTaskTable(internal.ToPointer("task"))
				_ = m.SetUseReadReplica(internal.ToPointer(true))
				_ = m.SetView(internal.ToPointer(ViewPortal))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteInt32Value", limitKey, mock.Anything).Return(nil)
				w.On("WriteAnyValue", otherInputsKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *AvailabilityRequestModel {
				m := NewAvailabilityRequest()
				_ = m.SetCatalogID(internal.ToPointer("catalog-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", catalogIDKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			if test.setupMock != nil {
				test.setupMock(writer)
			}

			err := test.model.Serialize(writer)

			if test.wantErr != nil {
				require.ErrorIs(t, err, test.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}
