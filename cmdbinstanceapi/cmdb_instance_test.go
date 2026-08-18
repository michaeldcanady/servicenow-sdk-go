package cmdbinstanceapi

import (
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// errWrite is a stand-in for an error a serialization.SerializationWriter can
// return from a Write* call.
var errWrite = errors.New("write error")

func TestNewCmdbInstance(t *testing.T) {
	model := NewCmdbInstance()
	require.NotNil(t, model)
}

func TestCmdbInstanceModel_GettersSetters(t *testing.T) {
	model := NewCmdbInstance()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"SysID", func(v any) error { return model.SetSysID(v.(*string)) }, func() (any, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"Name", func(v any) error { return model.SetName(v.(*string)) }, func() (any, error) { return model.GetName() }, internal.ToPointer("server01")},
		{"ClassName", func(v any) error { return model.SetClassName(v.(*string)) }, func() (any, error) { return model.GetClassName() }, internal.ToPointer("cmdb_ci_linux_server")},
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

func TestCmdbInstanceModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *CmdbInstanceModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			// Unlike appserviceapi's models, CmdbInstanceModel.Serialize has no
			// conversion.IsNil(m) guard, so a nil receiver surfaces ErrNilModel
			// from the first accessor call instead of a no-op nil return.
			name:    "nil model surfaces ErrNilModel",
			model:   nil,
			wantErr: snerrors.ErrNilModel,
		},
		{
			name:  "empty model writes nothing",
			model: NewCmdbInstance(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *CmdbInstanceModel {
				m := NewCmdbInstance()
				_ = m.SetSysID(internal.ToPointer("sys-id"))
				_ = m.SetName(internal.ToPointer("server01"))
				_ = m.SetClassName(internal.ToPointer("cmdb_ci_linux_server"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *CmdbInstanceModel {
				m := NewCmdbInstance()
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

func TestCmdbInstanceModel_GetFieldDeserializers(t *testing.T) {
	model := NewCmdbInstance()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{sysIDKey, nameKey, classNameKey, attributesKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 4)
}

func TestCmdbInstanceModel_AttributesPromoteTopLevelFields(t *testing.T) {
	model := NewCmdbInstance()
	nested := NewCmdbInstance()
	require.NoError(t, nested.SetSysID(internal.ToPointer("from-attrs")))
	require.NoError(t, nested.SetName(internal.ToPointer("name-from-attrs")))

	node := mocking.NewMockParseNode()
	node.On("GetObjectValue", mock.Anything).Return(nested, nil)

	require.NoError(t, model.GetFieldDeserializers()[attributesKey](node))

	sysID, err := model.GetSysID()
	require.NoError(t, err)
	require.NotNil(t, sysID)
	assert.Equal(t, "from-attrs", *sysID)

	name, err := model.GetName()
	require.NoError(t, err)
	require.NotNil(t, name)
	assert.Equal(t, "name-from-attrs", *name)

	attrs, err := model.GetAttributes()
	require.NoError(t, err)
	assert.Equal(t, nested, attrs)
}

func TestCmdbInstanceModel_AttributesDoNotOverwriteTopLevelFields(t *testing.T) {
	model := NewCmdbInstance()
	require.NoError(t, model.SetSysID(internal.ToPointer("top-level")))

	nested := NewCmdbInstance()
	require.NoError(t, nested.SetSysID(internal.ToPointer("from-attrs")))

	node := mocking.NewMockParseNode()
	node.On("GetObjectValue", mock.Anything).Return(nested, nil)

	require.NoError(t, model.GetFieldDeserializers()[attributesKey](node))

	sysID, err := model.GetSysID()
	require.NoError(t, err)
	require.NotNil(t, sysID)
	assert.Equal(t, "top-level", *sysID)
}

func TestCreateCmdbInstanceFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateCmdbInstanceFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
