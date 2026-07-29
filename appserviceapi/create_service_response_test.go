package appserviceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateCreateServiceResponseFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateCreateServiceResponseFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestFindServiceResultModel_GettersSetters(t *testing.T) {
	model := NewFindServiceResult()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"SysID", func(v any) error { return model.setSysID(v.(*string)) }, func() (any, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"Name", func(v any) error { return model.setName(v.(*string)) }, func() (any, error) { return model.GetName() }, internal.ToPointer("Email_East")},
		{"Number", func(v any) error { return model.setNumber(v.(*string)) }, func() (any, error) { return model.GetNumber() }, internal.ToPointer("SNSVC0001018")},
		{"Environment", func(v any) error { return model.setEnvironment(v.(*string)) }, func() (any, error) { return model.GetEnvironment() }, internal.ToPointer("Production")},
		{"Version", func(v any) error { return model.setVersion(v.(*string)) }, func() (any, error) { return model.GetVersion() }, internal.ToPointer("1.0")},
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

func TestFindServiceResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *FindServiceResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewFindServiceResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *FindServiceResult {
				m := NewFindServiceResult()
				_ = m.setSysID(internal.ToPointer("sys-id"))
				_ = m.setName(internal.ToPointer("Email_East"))
				_ = m.setNumber(internal.ToPointer("SNSVC0001018"))
				_ = m.setEnvironment(internal.ToPointer("Production"))
				_ = m.setVersion(internal.ToPointer("1.0"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *FindServiceResult {
				m := NewFindServiceResult()
				_ = m.setSysID(internal.ToPointer("sys-id"))
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

func TestFindServiceResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewFindServiceResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{sysIDKey, nameKey, numberKey, environmentKey, versionKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 5)
}

func TestCreateFindServiceResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateFindServiceResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
