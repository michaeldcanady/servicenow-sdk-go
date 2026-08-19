package cdmeditorapi

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

func TestNodeResultModel_GettersSetters(t *testing.T) {
	model := NewNodeResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"SysID", func(v interface{}) error { return model.setSysID(v.(*string)) }, func() (interface{}, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, internal.ToPointer("node-name")},
		{"Type", func(v interface{}) error { return model.setType(v.(*string)) }, func() (interface{}, error) { return model.GetType() }, internal.ToPointer("folder")},
		{"Value", func(v interface{}) error { return model.setValue(v.(*string)) }, func() (interface{}, error) { return model.GetValue() }, internal.ToPointer("node-value")},
		{"Parent", func(v interface{}) error { return model.setParent(v.(*string)) }, func() (interface{}, error) { return model.GetParent() }, internal.ToPointer("parent-id")},
		{"CdmID", func(v interface{}) error { return model.setCdmID(v.(*string)) }, func() (interface{}, error) { return model.GetCdmID() }, internal.ToPointer("cdm-id")},
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

func TestValidationResultModel_GettersSetters(t *testing.T) {
	model := NewValidationResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"Status", func(v interface{}) error { return model.setStatus(v.(*string)) }, func() (interface{}, error) { return model.GetStatus() }, internal.ToPointer("success")},
		{"Errors", func(v interface{}) error { return model.setErrors(v) }, func() (interface{}, error) { return model.GetErrors() }, "no errors"},
		{"Warnings", func(v interface{}) error { return model.setWarnings(v) }, func() (interface{}, error) { return model.GetWarnings() }, "no warnings"},
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

func TestNodeCreateRequestModel_GettersSetters(t *testing.T) {
	model := NewNodeCreateRequest()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, internal.ToPointer("new-node")},
		{"Type", func(v interface{}) error { return model.setType(v.(*string)) }, func() (interface{}, error) { return model.GetType() }, internal.ToPointer("file")},
		{"ParentID", func(v interface{}) error { return model.setParentID(v.(*string)) }, func() (interface{}, error) { return model.GetParentID() }, internal.ToPointer("parent-id")},
		{"CdmID", func(v interface{}) error { return model.setCdmID(v.(*string)) }, func() (interface{}, error) { return model.GetCdmID() }, internal.ToPointer("cdm-id")},
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

func TestNodeUpdateRequestModel_GettersSetters(t *testing.T) {
	model := NewNodeUpdateRequest()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, internal.ToPointer("updated-name")},
		{"Value", func(v interface{}) error { return model.setValue(v.(*string)) }, func() (interface{}, error) { return model.GetValue() }, internal.ToPointer("updated-value")},
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

func TestCreateNodeResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateNodeResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateValidationResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateValidationResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateNodeCreateRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateNodeCreateRequestFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateNodeUpdateRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateNodeUpdateRequestFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

// ---------------------------------------------------------------------------
// MessageResult
// ---------------------------------------------------------------------------

func TestNewMessageResult(t *testing.T) {
	msg := NewMessageResult(internal.ToPointer("deleted"))
	require.NotNil(t, msg)
	assert.Equal(t, internal.ToPointer("deleted"), msg.Message)
}

func TestMessageResultModel_Serialize(t *testing.T) {
	msg := NewMessageResult(internal.ToPointer("deleted"))
	writer := mocking.NewMockSerializationWriter()
	err := msg.Serialize(writer)
	require.NoError(t, err)
	writer.AssertNotCalled(t, "WriteStringValue", mock.Anything, mock.Anything)
}

func TestMessageResultModel_GetFieldDeserializers(t *testing.T) {
	msg := NewMessageResult(internal.ToPointer("deleted"))
	assert.Nil(t, msg.GetFieldDeserializers())
}

func TestCreateMessageResultFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(n *mocking.MockParseNode)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path - returns MessageResult with string value",
			setupMock: func(n *mocking.MockParseNode) {
				n.On("GetStringValue").Return(internal.ToPointer("done"), nil)
			},
		},
		{
			name: "deserialization error propagates",
			setupMock: func(n *mocking.MockParseNode) {
				n.On("GetStringValue").Return((*string)(nil), errWrite)
			},
			wantErr: errWrite,
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := mocking.NewMockParseNode()
			tt.setupMock(node)

			parsable, err := CreateMessageResultFromDiscriminatorValue(node)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
			if tt.wantNil {
				assert.Nil(t, parsable)
				return
			}
			assert.NotNil(t, parsable)
		})
	}
}

// ---------------------------------------------------------------------------
// Serialize / GetFieldDeserializers - node/validation models
// ---------------------------------------------------------------------------

func TestNodeResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *NodeResultModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewNodeResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *NodeResultModel {
				m := NewNodeResult()
				_ = m.setSysID(internal.ToPointer("sys-id"))
				_ = m.setName(internal.ToPointer("node-name"))
				_ = m.setType(internal.ToPointer("folder"))
				_ = m.setValue(internal.ToPointer("node-value"))
				_ = m.setParent(internal.ToPointer("parent-id"))
				_ = m.setCdmID(internal.ToPointer("cdm-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *NodeResultModel {
				m := NewNodeResult()
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

func TestNodeResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewNodeResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{sysIDKey, nameKey, typeKey, valueKey, parentKey, cdmIDKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 6)
}

func TestValidationResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ValidationResultModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nil errors/warnings (SerializeAnyFunc has no nil-guard)",
			model: NewValidationResult(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteAnyValue", errorsKey, mock.Anything).Return(nil)
				w.On("WriteAnyValue", warningsKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "happy path - writes all fields",
			model: func() *ValidationResultModel {
				m := NewValidationResult()
				_ = m.setStatus(internal.ToPointer("success"))
				_ = m.setErrors("no errors")
				_ = m.setWarnings("no warnings")
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", statusKey, mock.Anything).Return(nil)
				w.On("WriteAnyValue", errorsKey, mock.Anything).Return(nil)
				w.On("WriteAnyValue", warningsKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ValidationResultModel {
				m := NewValidationResult()
				_ = m.setStatus(internal.ToPointer("failure"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", statusKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "errors write error propagates",
			model: func() *ValidationResultModel {
				m := NewValidationResult()
				_ = m.setStatus(internal.ToPointer("failure"))
				_ = m.setErrors("boom")
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", statusKey, mock.Anything).Return(nil)
				w.On("WriteAnyValue", errorsKey, mock.Anything).Return(errWrite)
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

func TestValidationResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewValidationResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{statusKey, errorsKey, warningsKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 3)
}

func TestNodeCreateRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *NodeCreateRequestModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewNodeCreateRequest(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *NodeCreateRequestModel {
				m := NewNodeCreateRequest()
				_ = m.setName(internal.ToPointer("new-node"))
				_ = m.setType(internal.ToPointer("file"))
				_ = m.setParentID(internal.ToPointer("parent-id"))
				_ = m.setCdmID(internal.ToPointer("cdm-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *NodeCreateRequestModel {
				m := NewNodeCreateRequest()
				_ = m.setName(internal.ToPointer("new-node"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", nameKey, mock.Anything).Return(errWrite)
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

func TestNodeCreateRequestModel_GetFieldDeserializers(t *testing.T) {
	model := NewNodeCreateRequest()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{nameKey, typeKey, parentIDKey, cdmIDKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 4)
}

func TestNodeUpdateRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *NodeUpdateRequestModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewNodeUpdateRequest(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *NodeUpdateRequestModel {
				m := NewNodeUpdateRequest()
				_ = m.setName(internal.ToPointer("updated-name"))
				_ = m.setValue(internal.ToPointer("updated-value"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *NodeUpdateRequestModel {
				m := NewNodeUpdateRequest()
				_ = m.setName(internal.ToPointer("updated-name"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", nameKey, mock.Anything).Return(errWrite)
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

func TestNodeUpdateRequestModel_GetFieldDeserializers(t *testing.T) {
	model := NewNodeUpdateRequest()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{nameKey, valueKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 2)
}

// ---------------------------------------------------------------------------
// Response discriminator factories
// ---------------------------------------------------------------------------

func TestResponses_CreateFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name    string
		factory func() (any, error)
	}{
		{"NodesResponse", func() (any, error) { return CreateNodesResponseFromDiscriminatorValue(nil) }},
		{"NodeResponse", func() (any, error) { return CreateNodeResponseFromDiscriminatorValue(nil) }},
		{"ValidationResponse", func() (any, error) { return CreateValidationResponseFromDiscriminatorValue(nil) }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsable, err := tt.factory()
			require.NoError(t, err)
			assert.NotNil(t, parsable)
		})
	}
}
