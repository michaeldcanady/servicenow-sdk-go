package documentsapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// errWrite is a stand-in for an error a serialization.SerializationWriter can
// return from a Write* call.
var errWrite = errors.New("write error")

func TestDocumentModel_GettersSetters(t *testing.T) {
	model := NewDocument()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"SysID", func(v any) error { return model.SetSysID(v.(*string)) }, func() (any, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"Name", func(v any) error { return model.SetName(v.(*string)) }, func() (any, error) { return model.GetName() }, internal.ToPointer("name")},
		{"Type", func(v any) error { return model.SetType(v.(*string)) }, func() (any, error) { return model.GetType() }, internal.ToPointer("pdf")},
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

func TestCreateDocumentFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateDocumentFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestDocumentModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *DocumentModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "empty model writes nothing",
			model: NewDocument(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *DocumentModel {
				m := NewDocument()
				_ = m.SetSysID(internal.ToPointer("sys-id"))
				_ = m.SetName(internal.ToPointer("name"))
				_ = m.SetType(internal.ToPointer("pdf"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *DocumentModel {
				m := NewDocument()
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

func TestDocumentModel_GetFieldDeserializers(t *testing.T) {
	model := NewDocument()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{sysIDKey, nameKey, typeKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 3)
}
