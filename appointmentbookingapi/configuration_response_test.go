package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// ConfigurationResponse
// ---------------------------------------------------------------------------

func TestCreateConfigurationResponseFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateConfigurationResponseFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

// ---------------------------------------------------------------------------
// ConfigurationResult
// ---------------------------------------------------------------------------

func TestConfigurationResult_GettersSetters(t *testing.T) {
	model := NewConfigurationResult()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Active", func(v any) error { return model.SetActive(v.(*bool)) }, func() (any, error) { return model.GetActive() }, internal.ToPointer(true)},
		{"ActiveString", func(v any) error { return model.SetActiveString(v.(*string)) }, func() (any, error) { return model.GetActiveString() }, internal.ToPointer("true")},
		{"AdvancedCalendarViewPortal", func(v any) error { return model.SetAdvancedCalendarViewPortal(v.(*bool)) }, func() (any, error) { return model.GetAdvancedCalendarViewPortal() }, internal.ToPointer(true)},
		{"AutoAcceptance", func(v any) error { return model.SetAutoAcceptance(v.(*bool)) }, func() (any, error) { return model.GetAutoAcceptance() }, internal.ToPointer(false)},
		{"LocaleLanguage", func(v any) error { return model.SetLocaleLanguage(v.(*string)) }, func() (any, error) { return model.GetLocaleLanguage() }, internal.ToPointer("en")},
		{"ServiceConfig", func(v any) error { return model.SetServiceConfig(v.(ServiceConfig)) }, func() (any, error) { return model.GetServiceConfig() }, ServiceConfig(NewServiceConfig())},
		{"TaskTable", func(v any) error { return model.SetTaskTable(v.(*string)) }, func() (any, error) { return model.GetTaskTable() }, internal.ToPointer("task")},
		{"Translations", func(v any) error { return model.SetTranslations(v) }, func() (any, error) { return model.GetTranslations() }, "translated"},
		{"UserDateFormatOptions", func(v any) error { return model.SetUserDateFormatOptions(v.(UserDateFormatOptions)) }, func() (any, error) { return model.GetUserDateFormatOptions() }, UserDateFormatOptions(NewUserDateFormatOptions())},
		{"UseRR", func(v any) error { return model.SetUseRR(v.(*bool)) }, func() (any, error) { return model.GetUseRR() }, internal.ToPointer(true)},
		{"UserTimeFormat", func(v any) error { return model.SetUserTimeFormat(v.(UserTimeFormat)) }, func() (any, error) { return model.GetUserTimeFormat() }, UserTimeFormat(NewUserTimeFormat())},
		{"UserTimeFormatOptions", func(v any) error { return model.SetUserTimeFormatOptions(v.(UserTimeFormatOptions)) }, func() (any, error) { return model.GetUserTimeFormatOptions() }, UserTimeFormatOptions(NewUserTimeFormatOptions())},
		{"ViewScale", func(v any) error { return model.SetViewScale(v.(*string)) }, func() (any, error) { return model.GetViewScale() }, internal.ToPointer("day")},
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

func TestCreateConfigurationResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateConfigurationResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestConfigurationResult_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ConfigurationResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes only the any-typed field unconditionally",
			model: NewConfigurationResult(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteAnyValue", translationsKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "happy path - writes all fields including nested objects",
			model: func() *ConfigurationResult {
				m := NewConfigurationResult()
				_ = m.SetActive(internal.ToPointer(true))
				_ = m.SetActiveString(internal.ToPointer("true"))
				_ = m.SetAdvancedCalendarViewPortal(internal.ToPointer(true))
				_ = m.SetAutoAcceptance(internal.ToPointer(false))
				_ = m.SetLocaleLanguage(internal.ToPointer("en"))
				_ = m.SetServiceConfig(NewServiceConfig())
				_ = m.SetTaskTable(internal.ToPointer("task"))
				_ = m.SetTranslations("translated")
				_ = m.SetUserDateFormatOptions(NewUserDateFormatOptions())
				_ = m.SetUseRR(internal.ToPointer(true))
				_ = m.SetUserTimeFormat(NewUserTimeFormat())
				_ = m.SetUserTimeFormatOptions(NewUserTimeFormatOptions())
				_ = m.SetViewScale(internal.ToPointer("day"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				w.On("WriteAnyValue", translationsKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ConfigurationResult {
				m := NewConfigurationResult()
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
			model: func() *ConfigurationResult {
				m := NewConfigurationResult()
				_ = m.SetServiceConfig(NewServiceConfig())
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteObjectValue", serviceConfigKey, mock.Anything, mock.Anything).Return(errWrite)
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

func TestConfigurationResult_GetFieldDeserializers(t *testing.T) {
	model := NewConfigurationResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{
		activeKey, activeStringKey, advancedCalendarViewPortalKey, autoAcceptanceKey,
		localeLanguageKey, serviceConfigKey, taskTableKey, translationsKey,
		userDateFormatOptionsKey, useRRKey, userTimeFormatKey, userTimeFormatOptionsKey,
		viewScaleKey,
	} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 13)
}

// ---------------------------------------------------------------------------
// UserDateFormatOptionsModel
// ---------------------------------------------------------------------------

func TestUserDateFormatOptionsModel_GettersSetters(t *testing.T) {
	model := NewUserDateFormatOptions()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Day", func(v any) error { return model.SetDay(v.(*string)) }, func() (any, error) { return model.GetDay() }, internal.ToPointer("2-digit")},
		{"Month", func(v any) error { return model.SetMonth(v.(*string)) }, func() (any, error) { return model.GetMonth() }, internal.ToPointer("short")},
		{"Week", func(v any) error { return model.SetWeek(v.(*string)) }, func() (any, error) { return model.GetWeek() }, internal.ToPointer("1")},
		{"Weekday", func(v any) error { return model.SetWeekday(v.(*string)) }, func() (any, error) { return model.GetWeekday() }, internal.ToPointer("long")},
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

func TestCreateUserDateFormatOptionsFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateUserDateFormatOptionsFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestUserDateFormatOptionsModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *UserDateFormatOptionsModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewUserDateFormatOptions(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *UserDateFormatOptionsModel {
				m := NewUserDateFormatOptions()
				_ = m.SetDay(internal.ToPointer("2-digit"))
				_ = m.SetMonth(internal.ToPointer("short"))
				_ = m.SetWeek(internal.ToPointer("1"))
				_ = m.SetWeekday(internal.ToPointer("long"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *UserDateFormatOptionsModel {
				m := NewUserDateFormatOptions()
				_ = m.SetDay(internal.ToPointer("2-digit"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", dayKey, mock.Anything).Return(errWrite)
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

func TestUserDateFormatOptionsModel_GetFieldDeserializers(t *testing.T) {
	model := NewUserDateFormatOptions()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{dayKey, monthKey, weekKey, weekdayKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 4)
}
