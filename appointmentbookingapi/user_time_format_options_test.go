package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUserTimeFormatOptionsModel_GettersSetters(t *testing.T) {
	model := NewUserTimeFormatOptions()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Hour", func(v any) error { return model.SetHour(v.(*string)) }, func() (any, error) { return model.GetHour() }, internal.ToPointer("2-digit")},
		{"HourCycle", func(v any) error { return model.SetHourCycle(v.(*string)) }, func() (any, error) { return model.GetHourCycle() }, internal.ToPointer("h12")},
		{"Minute", func(v any) error { return model.SetMinute(v.(*string)) }, func() (any, error) { return model.GetMinute() }, internal.ToPointer("2-digit")},
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

func TestCreateUserTimeFormatOptionsFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateUserTimeFormatOptionsFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestUserTimeFormatOptionsModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *UserTimeFormatOptionsModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewUserTimeFormatOptions(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *UserTimeFormatOptionsModel {
				m := NewUserTimeFormatOptions()
				_ = m.SetHour(internal.ToPointer("2-digit"))
				_ = m.SetHourCycle(internal.ToPointer("h12"))
				_ = m.SetMinute(internal.ToPointer("2-digit"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *UserTimeFormatOptionsModel {
				m := NewUserTimeFormatOptions()
				_ = m.SetHour(internal.ToPointer("2-digit"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", hourKey, mock.Anything).Return(errWrite)
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

func TestUserTimeFormatOptionsModel_GetFieldDeserializers(t *testing.T) {
	model := NewUserTimeFormatOptions()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{hourKey, hourCycleKey, minuteKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 3)
}
