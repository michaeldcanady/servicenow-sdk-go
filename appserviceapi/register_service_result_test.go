package appserviceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterServiceResultModel_GettersSetters(t *testing.T) {
	model := NewRegisterServiceResult()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"SysID", func(v any) error { return model.SetSysID(v.(*string)) }, func() (any, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"Number", func(v any) error { return model.SetNumber(v.(*string)) }, func() (any, error) { return model.GetNumber() }, internal.ToPointer("SNSVC0001019")},
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

func TestRegisterServiceResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *RegisterServiceResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewRegisterServiceResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *RegisterServiceResult {
				m := NewRegisterServiceResult()
				_ = m.SetSysID(internal.ToPointer("sys-id"))
				_ = m.SetNumber(internal.ToPointer("SNSVC0001019"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *RegisterServiceResult {
				m := NewRegisterServiceResult()
				_ = m.SetSysID(internal.ToPointer("sys-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", sysIDKey, mock.Anything).Return(errWrite)
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

func TestRegisterServiceResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewRegisterServiceResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{sysIDKey, numberKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 2)
}

func TestCreateRegisterServiceResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateRegisterServiceResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
