// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi

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

func TestUploadStatusOutputModel_GettersSetters(t *testing.T) {
	model := NewUploadStatusOutput()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"SysID", func(v interface{}) error { return model.setSysID(v.(*string)) }, func() (interface{}, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"Number", func(v interface{}) error { return model.setNumber(v.(*string)) }, func() (interface{}, error) { return model.GetNumber() }, internal.ToPointer("123")},
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

func TestUploadStatusResultModel_GettersSetters(t *testing.T) {
	model := NewUploadStatusResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"Type", func(v interface{}) error { return model.setType(v.(*string)) }, func() (interface{}, error) { return model.GetType() }, internal.ToPointer("upload")},
		{"State", func(v interface{}) error { return model.setState(v.(*string)) }, func() (interface{}, error) { return model.GetState() }, internal.ToPointer("completed")},
		{"Output", func(v interface{}) error { return model.setOutput(v.(*UploadStatusOutput)) }, func() (interface{}, error) { return model.GetOutput() }, NewUploadStatusOutput()},
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

func TestExportResultModel_GettersSetters(t *testing.T) {
	model := NewExportResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"SysID", func(v interface{}) error { return model.setSysID(v.(*string)) }, func() (interface{}, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, internal.ToPointer("export-name")},
		{"State", func(v interface{}) error { return model.setState(v.(*string)) }, func() (interface{}, error) { return model.GetState() }, internal.ToPointer("active")},
		{"Status", func(v interface{}) error { return model.setStatus(v.(*string)) }, func() (interface{}, error) { return model.GetStatus() }, internal.ToPointer("success")},
		{"Message", func(v interface{}) error { return model.setMessage(v.(*string)) }, func() (interface{}, error) { return model.GetMessage() }, internal.ToPointer("export successful")},
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

func TestExportStatusResultModel_GettersSetters(t *testing.T) {
	model := NewExportStatusResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"State", func(v interface{}) error { return model.setState(v.(*string)) }, func() (interface{}, error) { return model.GetState() }, internal.ToPointer("completed")},
		{"Status", func(v interface{}) error { return model.setStatus(v.(*string)) }, func() (interface{}, error) { return model.GetStatus() }, internal.ToPointer("success")},
		{"Message", func(v interface{}) error { return model.setMessage(v.(*string)) }, func() (interface{}, error) { return model.GetMessage() }, internal.ToPointer("status message")},
		{"Progress", func(v interface{}) error { return model.setProgress(v.(*string)) }, func() (interface{}, error) { return model.GetProgress() }, internal.ToPointer("100")},
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

func TestCreateUploadStatusResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateUploadStatusResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateExportResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateExportResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateComponentUploadRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateComponentUploadRequestFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateUploadStatusOutputFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateUploadStatusOutputFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestSharedLibraryComponentApplicationModel_GettersSetters(t *testing.T) {
	model := NewSharedLibraryComponentApplication()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"SysID", func(v interface{}) error { return model.setSysID(v.(*string)) }, func() (interface{}, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, internal.ToPointer("app-name")},
		{"Version", func(v interface{}) error { return model.setVersion(v.(*string)) }, func() (interface{}, error) { return model.GetVersion() }, internal.ToPointer("1.0.0")},
		{"Description", func(v interface{}) error { return model.setDescription(v.(*string)) }, func() (interface{}, error) { return model.GetDescription() }, internal.ToPointer("app-desc")},
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, internal.ToPointer("actual-app-name")},
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

func TestComponentUploadRequestModel_GettersSetters(t *testing.T) {
	model := NewComponentUploadRequest()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, internal.ToPointer("app-name")},
		{"ComponentName", func(v interface{}) error { return model.setComponentName(v.(*string)) }, func() (interface{}, error) { return model.GetComponentName() }, internal.ToPointer("comp-name")},
		{"Data", func(v interface{}) error { return model.setData(v.(*string)) }, func() (interface{}, error) { return model.GetData() }, internal.ToPointer("base64-data")},
		{"Format", func(v interface{}) error { return model.setFormat(v.(*string)) }, func() (interface{}, error) { return model.GetFormat() }, internal.ToPointer("json")},
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

func TestCreateSharedLibraryComponentApplicationFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateSharedLibraryComponentApplicationFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestComponentVarsUploadRequestModel_GettersSetters(t *testing.T) {
	model := NewComponentVarsUploadRequest()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, internal.ToPointer("app-name")},
		{"ComponentName", func(v interface{}) error { return model.setComponentName(v.(*string)) }, func() (interface{}, error) { return model.GetComponentName() }, internal.ToPointer("comp-name")},
		{"Vars", func(v interface{}) error { return model.setVars(v) }, func() (interface{}, error) { return model.GetVars() }, "vars"},
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

func TestCollectionUploadRequestModel_GettersSetters(t *testing.T) {
	model := NewCollectionUploadRequest()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, internal.ToPointer("app-name")},
		{"CollectionName", func(v interface{}) error { return model.setCollectionName(v.(*string)) }, func() (interface{}, error) { return model.GetCollectionName() }, internal.ToPointer("coll-name")},
		{"Data", func(v interface{}) error { return model.setData(v.(*string)) }, func() (interface{}, error) { return model.GetData() }, internal.ToPointer("base64-data")},
		{"Format", func(v interface{}) error { return model.setFormat(v.(*string)) }, func() (interface{}, error) { return model.GetFormat() }, internal.ToPointer("json")},
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

func TestSharedComponentUpdateRequestModel_GettersSetters(t *testing.T) {
	model := NewSharedComponentUpdateRequest()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, internal.ToPointer("app-name")},
		{"SharedComponentName", func(v interface{}) error { return model.setSharedComponentName(v.(*string)) }, func() (interface{}, error) { return model.GetSharedComponentName() }, internal.ToPointer("shared-name")},
		{"Data", func(v interface{}) error { return model.setData(v.(*string)) }, func() (interface{}, error) { return model.GetData() }, internal.ToPointer("new-data")},
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

func TestCreateExportStatusResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateExportStatusResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateComponentVarsUploadRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateComponentVarsUploadRequestFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateCollectionUploadRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateCollectionUploadRequestFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateSharedComponentUpdateRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateSharedComponentUpdateRequestFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

// ---------------------------------------------------------------------------
// Response discriminator factories
// ---------------------------------------------------------------------------

func TestResponses_CreateFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name    string
		factory func() (any, error)
	}{
		{"UploadStatusResponse", func() (any, error) { return CreateUploadStatusResponseFromDiscriminatorValue(nil) }},
		{"ExportsResponse", func() (any, error) { return CreateExportsResponseFromDiscriminatorValue(nil) }},
		{"ExportStatusResponse", func() (any, error) { return CreateExportStatusResponseFromDiscriminatorValue(nil) }},
		{"SharedLibrariesComponentsApplicationsResponse", func() (any, error) {
			return CreateSharedLibrariesComponentsApplicationsResponseFromDiscriminatorValue(nil)
		}},
		{"DeployableUpdateResponse", func() (any, error) { return CreateDeployableUpdateResponseFromDiscriminatorValue(nil) }},
		{"SharedComponentUpdateResponse", func() (any, error) { return CreateSharedComponentUpdateResponseFromDiscriminatorValue(nil) }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsable, err := tt.factory()
			require.NoError(t, err)
			assert.NotNil(t, parsable)
		})
	}
}

// ---------------------------------------------------------------------------
// Media
// ---------------------------------------------------------------------------

func TestNewMedia(t *testing.T) {
	media := NewMedia("application/octet-stream", []byte("data"))
	require.NotNil(t, media)
	assert.Equal(t, "application/octet-stream", media.GetContentType())
	assert.Equal(t, []byte("data"), media.GetData())
}

func TestMedia_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		media     *Media
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "happy path - writes byte array",
			media: NewMedia("application/octet-stream", []byte("data")),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteByteArrayValue", "", []byte("data")).Return(nil)
			},
		},
		{
			name:  "write error propagates",
			media: NewMedia("application/octet-stream", []byte("data")),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteByteArrayValue", "", []byte("data")).Return(errWrite)
			},
			wantErr: errWrite,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			tt.setupMock(writer)

			err := tt.media.Serialize(writer)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMedia_GetFieldDeserializers(t *testing.T) {
	media := NewMedia("application/octet-stream", []byte("data"))
	assert.Nil(t, media.GetFieldDeserializers())
}

// ---------------------------------------------------------------------------
// Serialize / GetFieldDeserializers - upload/export/shared-component models
// ---------------------------------------------------------------------------

func TestUploadStatusOutputModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *UploadStatusOutput
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewUploadStatusOutput(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *UploadStatusOutput {
				m := NewUploadStatusOutput()
				_ = m.setSysID(internal.ToPointer("sys-id"))
				_ = m.setNumber(internal.ToPointer("123"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *UploadStatusOutput {
				m := NewUploadStatusOutput()
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

func TestUploadStatusOutputModel_GetFieldDeserializers(t *testing.T) {
	model := NewUploadStatusOutput()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{sysIDKey, numberKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 2)
}

func TestUploadStatusResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *UploadStatusResultModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewUploadStatusResult(),
		},
		{
			name: "happy path - writes all fields including nested output",
			model: func() *UploadStatusResultModel {
				m := NewUploadStatusResult()
				_ = m.setType(internal.ToPointer("upload"))
				_ = m.setState(internal.ToPointer("completed"))
				_ = m.setOutput(NewUploadStatusOutput())
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteObjectValue", outputKey, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *UploadStatusResultModel {
				m := NewUploadStatusResult()
				_ = m.setType(internal.ToPointer("upload"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", typeKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "nested object write error propagates",
			model: func() *UploadStatusResultModel {
				m := NewUploadStatusResult()
				_ = m.setOutput(NewUploadStatusOutput())
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteObjectValue", outputKey, mock.Anything, mock.Anything).Return(errWrite)
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

func TestUploadStatusResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewUploadStatusResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{typeKey, stateKey, outputKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 3)
}

func TestExportResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ExportResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewExportResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *ExportResult {
				m := NewExportResult()
				_ = m.setSysID(internal.ToPointer("sys-id"))
				_ = m.setName(internal.ToPointer("export-name"))
				_ = m.setState(internal.ToPointer("active"))
				_ = m.setStatus(internal.ToPointer("success"))
				_ = m.setMessage(internal.ToPointer("ok"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ExportResult {
				m := NewExportResult()
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

func TestExportResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewExportResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{sysIDKey, nameKey, stateKey, statusKey, messageKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 5)
}

func TestExportStatusResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ExportStatusResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewExportStatusResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *ExportStatusResult {
				m := NewExportStatusResult()
				_ = m.setState(internal.ToPointer("completed"))
				_ = m.setStatus(internal.ToPointer("success"))
				_ = m.setMessage(internal.ToPointer("done"))
				_ = m.setProgress(internal.ToPointer("100"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ExportStatusResult {
				m := NewExportStatusResult()
				_ = m.setState(internal.ToPointer("completed"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", stateKey, mock.Anything).Return(errWrite)
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

func TestExportStatusResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewExportStatusResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{stateKey, statusKey, messageKey, progressKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 4)
}

func TestSharedLibraryComponentApplicationModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *SharedLibraryComponentApplication
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewSharedLibraryComponentApplication(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *SharedLibraryComponentApplication {
				m := NewSharedLibraryComponentApplication()
				_ = m.setSysID(internal.ToPointer("sys-id"))
				_ = m.setName(internal.ToPointer("app-name"))
				_ = m.setVersion(internal.ToPointer("1.0.0"))
				_ = m.setDescription(internal.ToPointer("desc"))
				_ = m.setAppName(internal.ToPointer("actual-app-name"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *SharedLibraryComponentApplication {
				m := NewSharedLibraryComponentApplication()
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

func TestSharedLibraryComponentApplicationModel_GetFieldDeserializers(t *testing.T) {
	model := NewSharedLibraryComponentApplication()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{sysIDKey, nameKey, versionKey, descriptionKey, appNameKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 5)
}

func TestCollectionUploadRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *CollectionUploadRequest
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewCollectionUploadRequest(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *CollectionUploadRequest {
				m := NewCollectionUploadRequest()
				_ = m.setAppName(internal.ToPointer("app-name"))
				_ = m.setCollectionName(internal.ToPointer("coll-name"))
				_ = m.setData(internal.ToPointer("base64-data"))
				_ = m.setFormat(internal.ToPointer("json"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *CollectionUploadRequest {
				m := NewCollectionUploadRequest()
				_ = m.setAppName(internal.ToPointer("app-name"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", appNameKey, mock.Anything).Return(errWrite)
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

func TestCollectionUploadRequestModel_GetFieldDeserializers(t *testing.T) {
	model := NewCollectionUploadRequest()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{appNameKey, collectionNameKey, dataKey, formatKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 4)
}

func TestComponentUploadRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ComponentUploadRequest
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewComponentUploadRequest(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *ComponentUploadRequest {
				m := NewComponentUploadRequest()
				_ = m.setAppName(internal.ToPointer("app-name"))
				_ = m.setComponentName(internal.ToPointer("comp-name"))
				_ = m.setData(internal.ToPointer("base64-data"))
				_ = m.setFormat(internal.ToPointer("json"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ComponentUploadRequest {
				m := NewComponentUploadRequest()
				_ = m.setAppName(internal.ToPointer("app-name"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", appNameKey, mock.Anything).Return(errWrite)
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

func TestComponentUploadRequestModel_GetFieldDeserializers(t *testing.T) {
	model := NewComponentUploadRequest()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{appNameKey, componentNameKey, dataKey, formatKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 4)
}

func TestComponentVarsUploadRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ComponentVarsUploadRequest
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nil vars (SerializeAnyFunc has no nil-guard)",
			model: NewComponentVarsUploadRequest(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteAnyValue", varsKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "happy path - writes all fields",
			model: func() *ComponentVarsUploadRequest {
				m := NewComponentVarsUploadRequest()
				_ = m.setAppName(internal.ToPointer("app-name"))
				_ = m.setComponentName(internal.ToPointer("comp-name"))
				_ = m.setVars("vars")
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteAnyValue", varsKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ComponentVarsUploadRequest {
				m := NewComponentVarsUploadRequest()
				_ = m.setAppName(internal.ToPointer("app-name"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", appNameKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "vars write error propagates",
			model: func() *ComponentVarsUploadRequest {
				m := NewComponentVarsUploadRequest()
				_ = m.setVars("vars")
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteAnyValue", varsKey, mock.Anything).Return(errWrite)
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

func TestComponentVarsUploadRequestModel_GetFieldDeserializers(t *testing.T) {
	model := NewComponentVarsUploadRequest()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{appNameKey, componentNameKey, varsKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 3)
}

func TestSharedComponentUpdateRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *SharedComponentUpdateRequest
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewSharedComponentUpdateRequest(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *SharedComponentUpdateRequest {
				m := NewSharedComponentUpdateRequest()
				_ = m.setAppName(internal.ToPointer("app-name"))
				_ = m.setSharedComponentName(internal.ToPointer("shared-name"))
				_ = m.setData(internal.ToPointer("new-data"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *SharedComponentUpdateRequest {
				m := NewSharedComponentUpdateRequest()
				_ = m.setAppName(internal.ToPointer("app-name"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", appNameKey, mock.Anything).Return(errWrite)
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

func TestSharedComponentUpdateRequestModel_GetFieldDeserializers(t *testing.T) {
	model := NewSharedComponentUpdateRequest()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{appNameKey, sharedComponentNameKey, dataKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 3)
}
