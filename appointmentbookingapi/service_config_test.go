package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestServiceConfigModel_GettersSetters(t *testing.T) {
	model := NewServiceConfig()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Active", func(v any) error { return model.SetActive(v.(*bool)) }, func() (any, error) { return model.GetActive() }, internal.ToPointer(true)},
		{"ActiveString", func(v any) error { return model.SetActiveString(v.(*string)) }, func() (any, error) { return model.GetActiveString() }, internal.ToPointer("true")},
		{"AppointmentBookingConfig", func(v any) error { return model.SetAppointmentBookingConfig(v.(*string)) }, func() (any, error) { return model.GetAppointmentBookingConfig() }, internal.ToPointer("cfg")},
		{"AppointmentDuration", func(v any) error { return model.SetAppointmentDuration(v.(*string)) }, func() (any, error) { return model.GetAppointmentDuration() }, internal.ToPointer("30")},
		{"AppointmentsPerBookableSlot", func(v any) error { return model.SetAppointmentsPerBookableSlot(v.(*string)) }, func() (any, error) { return model.GetAppointmentsPerBookableSlot() }, internal.ToPointer("1")},
		{"BookableDays", func(v any) error { return model.SetBookableDays(v.(*string)) }, func() (any, error) { return model.GetBookableDays() }, internal.ToPointer("MON,TUE")},
		{"CancelByTime", func(v any) error { return model.SetCancelByTime(v.(*string)) }, func() (any, error) { return model.GetCancelByTime() }, internal.ToPointer("24")},
		{"DefaultTimezone", func(v any) error { return model.SetDefaultTimezone(v.(*DefaultTimeZone)) }, func() (any, error) { return model.GetDefaultTimezone() }, internal.ToPointer(DefaultTimeZoneUser)},
		{"EnableAdvancedConfig", func(v any) error { return model.SetEnableAdvancedConfig(v.(*bool)) }, func() (any, error) { return model.GetEnableAdvancedConfig() }, internal.ToPointer(false)},
		{"FieldMapping", func(v any) error { return model.SetFieldMapping(v.(FieldMapping)) }, func() (any, error) { return model.GetFieldMapping() }, FieldMapping(NewFieldMapping())},
		{"FutureBookableMaxDays", func(v any) error { return model.SetFutureBookableMaxDays(v.(*string)) }, func() (any, error) { return model.GetFutureBookableMaxDays() }, internal.ToPointer("30")},
		{"LeadTime", func(v any) error { return model.SetLeadTime(v.(*string)) }, func() (any, error) { return model.GetLeadTime() }, internal.ToPointer("1")},
		{"Mandatory", func(v any) error { return model.SetMandatory(v.(*string)) }, func() (any, error) { return model.GetMandatory() }, internal.ToPointer("true")},
		{"UseSlotEndTimeAs", func(v any) error { return model.SetUseSlotEndTimeAs(v.(*string)) }, func() (any, error) { return model.GetUseSlotEndTimeAs() }, internal.ToPointer("end")},
		{"WorkDuration", func(v any) error { return model.SetWorkDuration(v.(*string)) }, func() (any, error) { return model.GetWorkDuration() }, internal.ToPointer("480")},
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

func TestCreateServiceConfigFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateServiceConfigFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestServiceConfigModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ServiceConfigModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewServiceConfig(),
		},
		{
			name: "happy path - writes all fields including nested field mapping",
			model: func() *ServiceConfigModel {
				m := NewServiceConfig()
				_ = m.SetActive(internal.ToPointer(true))
				_ = m.SetActiveString(internal.ToPointer("true"))
				_ = m.SetAppointmentBookingConfig(internal.ToPointer("cfg"))
				_ = m.SetAppointmentDuration(internal.ToPointer("30"))
				_ = m.SetAppointmentsPerBookableSlot(internal.ToPointer("1"))
				_ = m.SetBookableDays(internal.ToPointer("MON"))
				_ = m.SetCancelByTime(internal.ToPointer("24"))
				_ = m.SetDefaultTimezone(internal.ToPointer(DefaultTimeZoneUser))
				_ = m.SetEnableAdvancedConfig(internal.ToPointer(false))
				_ = m.SetFieldMapping(NewFieldMapping())
				_ = m.SetFutureBookableMaxDays(internal.ToPointer("30"))
				_ = m.SetLeadTime(internal.ToPointer("1"))
				_ = m.SetMandatory(internal.ToPointer("true"))
				_ = m.SetUseSlotEndTimeAs(internal.ToPointer("end"))
				_ = m.SetWorkDuration(internal.ToPointer("480"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ServiceConfigModel {
				m := NewServiceConfig()
				_ = m.SetActive(internal.ToPointer(true))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteBoolValue", activeKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "nested object write error propagates",
			model: func() *ServiceConfigModel {
				m := NewServiceConfig()
				_ = m.SetFieldMapping(NewFieldMapping())
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteObjectValue", fieldMappingKey, mock.Anything, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			if tt.setupMock != nil {
				tt.setupMock(writer)
			}

			err := tt.model.Serialize(writer)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestServiceConfigModel_GetFieldDeserializers(t *testing.T) {
	model := NewServiceConfig()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{
		activeKey, activeStringKey, appointmentBookingConfigKey, appointmentDurationKey,
		appointmentsPerBookableSlotKey, bookableDaysKey, cancelByTimeKey, defaultTimezoneKey,
		enableAdvancedConfigKey, fieldMappingKey, futureBookableMaxDaysKey, leadTimeKey,
		mandatoryKey, useSlotEndTimeAsKey, workDurationKey,
	} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 15)
}
