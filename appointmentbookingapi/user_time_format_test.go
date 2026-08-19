package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUserTimeFormatModel_GettersSetters(t *testing.T) {
	model := NewUserTimeFormat()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Type", func(v any) error { return model.SetType(v.(*TimeFormat)) }, func() (any, error) { return model.GetType() }, internal.ToPointer(TimeFormat12Hr)},
		{"Value", func(v any) error { return model.SetValue(v.(*string)) }, func() (any, error) { return model.GetValue() }, internal.ToPointer("h:mm a")},
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

func TestCreateUserTimeFormatFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateUserTimeFormatFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestUserTimeFormatModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *UserTimeFormatModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewUserTimeFormat(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *UserTimeFormatModel {
				m := NewUserTimeFormat()
				_ = m.SetType(internal.ToPointer(TimeFormat12Hr))
				_ = m.SetValue(internal.ToPointer("h:mm a"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *UserTimeFormatModel {
				m := NewUserTimeFormat()
				_ = m.SetType(internal.ToPointer(TimeFormat12Hr))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", typeKey, mock.Anything).Return(errWrite)
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

func TestUserTimeFormatModel_GetFieldDeserializers(t *testing.T) {
	model := NewUserTimeFormat()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{typeKey, valueKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 2)
}
