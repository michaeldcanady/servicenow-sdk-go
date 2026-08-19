package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRPVariableModel_GettersSetters(t *testing.T) {
	model := NewRPVariable()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"DisplayName", func(v any) error { return model.SetDisplayName(v.(*string)) }, func() (any, error) { return model.GetDisplayName() }, internal.ToPointer("Contact")},
		{"Label", func(v any) error { return model.SetLabel(v.(*string)) }, func() (any, error) { return model.GetLabel() }, internal.ToPointer("Contact label")},
		{"Name", func(v any) error { return model.SetName(v.(*string)) }, func() (any, error) { return model.GetName() }, internal.ToPointer("contact")},
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

func TestCreateRPVariableFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateRPVariableFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestRPVariableModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *RPVariableModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewRPVariable(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *RPVariableModel {
				m := NewRPVariable()
				_ = m.SetDisplayName(internal.ToPointer("Contact"))
				_ = m.SetLabel(internal.ToPointer("Contact label"))
				_ = m.SetName(internal.ToPointer("contact"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *RPVariableModel {
				m := NewRPVariable()
				_ = m.SetDisplayName(internal.ToPointer("Contact"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", displayNameKey, mock.Anything).Return(errWrite)
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

func TestRPVariableModel_GetFieldDeserializers(t *testing.T) {
	model := NewRPVariable()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{displayNameKey, labelKey, nameKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 3)
}
