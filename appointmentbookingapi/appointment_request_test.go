package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
		{"CatalogId", func(v interface{}) error { return model.SetCatalogID(v.(*string)) }, func() (interface{}, error) { return model.GetCatalogID() }, internal.ToPointer("val")},
		{"EndDateUTC", func(v interface{}) error { return model.SetEndDateUTC(v.(*string)) }, func() (interface{}, error) { return model.GetEndDateUTC() }, internal.ToPointer("val")},
		{"Location", func(v interface{}) error { return model.SetLocation(v.(*string)) }, func() (interface{}, error) { return model.GetLocation() }, internal.ToPointer("val")},
		{"OpenedFor", func(v interface{}) error { return model.SetOpenedFor(v.(*string)) }, func() (interface{}, error) { return model.GetOpenedFor() }, internal.ToPointer("val")},
		{"Reschedule", func(v interface{}) error { return model.SetReschedule(v.(*bool)) }, func() (interface{}, error) { return model.GetReschedule() }, internal.ToPointer(true)},
		{"ServiceConfigRule", func(v interface{}) error { return model.SetServiceConfigRule(v.(*string)) }, func() (interface{}, error) { return model.GetServiceConfigRule() }, internal.ToPointer("val")},
		{"StartDateUTC", func(v interface{}) error { return model.SetStartDateUTC(v.(*string)) }, func() (interface{}, error) { return model.GetStartDateUTC() }, internal.ToPointer("val")},
		{"TaskId", func(v interface{}) error { return model.SetTaskID(v.(*string)) }, func() (interface{}, error) { return model.GetTaskID() }, internal.ToPointer("val")},
		{"TaskTable", func(v interface{}) error { return model.SetTaskTable(v.(*string)) }, func() (interface{}, error) { return model.GetTaskTable() }, internal.ToPointer("val")},
		{"Timezone", func(v interface{}) error { return model.SetTimezone(v.(*string)) }, func() (interface{}, error) { return model.GetTimezone() }, internal.ToPointer("val")},
		{"ValidateRequest", func(v interface{}) error { return model.SetValidateRequest(v.(*bool)) }, func() (interface{}, error) { return model.GetValidateRequest() }, internal.ToPointer(true)},
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

func TestCreateAppointmentRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateAppointmentRequestFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestAppointmentRequestModel_GetFieldDeserializers(t *testing.T) {
	model := NewAppointmentRequest()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{
		actualEndDateKey, actualStartDateKey, catalogIDKey, endDateUTCKey, locationKey,
		openedForKey, rescheduleKey, serviceConfigRuleKey, startDateUTCKey, taskIDKey,
		taskTableKey, timezoneKey, validateRequestKey,
	} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 13)
}

func TestAppointmentRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *AppointmentRequestModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			// Every field here is a primitive serializer, all of which skip nil, so an
			// empty model writes nothing at all.
			name:  "empty model writes nothing",
			model: NewAppointmentRequest(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *AppointmentRequestModel {
				m := NewAppointmentRequest()
				_ = m.SetActualEndDate(internal.ToPointer("2026-08-01"))
				_ = m.SetActualStartDate(internal.ToPointer("2026-07-29"))
				_ = m.SetCatalogID(internal.ToPointer("catalog-id"))
				_ = m.SetEndDateUTC(internal.ToPointer("2026-08-01T10:00:00Z"))
				_ = m.SetLocation(internal.ToPointer("location"))
				_ = m.SetOpenedFor(internal.ToPointer("user"))
				_ = m.SetReschedule(internal.ToPointer(true))
				_ = m.SetServiceConfigRule(internal.ToPointer("rule"))
				_ = m.SetStartDateUTC(internal.ToPointer("2026-07-29T10:00:00Z"))
				_ = m.SetTaskID(internal.ToPointer("task-id"))
				_ = m.SetTaskTable(internal.ToPointer("task"))
				_ = m.SetTimezone(internal.ToPointer("UTC"))
				_ = m.SetValidateRequest(internal.ToPointer(false))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *AppointmentRequestModel {
				m := NewAppointmentRequest()
				_ = m.SetActualEndDate(internal.ToPointer("2026-08-01"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", actualEndDateKey, mock.Anything).Return(errWrite)
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
