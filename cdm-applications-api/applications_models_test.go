package cdmapplicationsapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptr[T any](v T) *T {
	return &v
}

func TestUploadStatusOutputModel_GettersSetters(t *testing.T) {
	model := NewUploadStatusOutput()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"SysId", func(v interface{}) error { return model.setSysId(v.(*string)) }, func() (interface{}, error) { return model.GetSysId() }, ptr("sys-id")},
		{"Number", func(v interface{}) error { return model.setNumber(v.(*string)) }, func() (interface{}, error) { return model.GetNumber() }, ptr("123")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
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
		{"Type", func(v interface{}) error { return model.setType(v.(*string)) }, func() (interface{}, error) { return model.GetType() }, ptr("upload")},
		{"State", func(v interface{}) error { return model.setState(v.(*string)) }, func() (interface{}, error) { return model.GetState() }, ptr("completed")},
		{"Output", func(v interface{}) error { return model.setOutput(v.(*UploadStatusOutput)) }, func() (interface{}, error) { return model.GetOutput() }, NewUploadStatusOutput()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
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
		{"SysId", func(v interface{}) error { return model.setSysId(v.(*string)) }, func() (interface{}, error) { return model.GetSysId() }, ptr("sys-id")},
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, ptr("export-name")},
		{"State", func(v interface{}) error { return model.setState(v.(*string)) }, func() (interface{}, error) { return model.GetState() }, ptr("active")},
		{"Status", func(v interface{}) error { return model.setStatus(v.(*string)) }, func() (interface{}, error) { return model.GetStatus() }, ptr("success")},
		{"Message", func(v interface{}) error { return model.setMessage(v.(*string)) }, func() (interface{}, error) { return model.GetMessage() }, ptr("export successful")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
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
		{"State", func(v interface{}) error { return model.setState(v.(*string)) }, func() (interface{}, error) { return model.GetState() }, ptr("completed")},
		{"Status", func(v interface{}) error { return model.setStatus(v.(*string)) }, func() (interface{}, error) { return model.GetStatus() }, ptr("success")},
		{"Message", func(v interface{}) error { return model.setMessage(v.(*string)) }, func() (interface{}, error) { return model.GetMessage() }, ptr("status message")},
		{"Progress", func(v interface{}) error { return model.setProgress(v.(*string)) }, func() (interface{}, error) { return model.GetProgress() }, ptr("100")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestCreateUploadStatusResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateUploadStatusResultFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateExportResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateExportResultFromDiscriminatorValue(nil)
	assert.NoError(t, err)
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
		{"SysId", func(v interface{}) error { return model.setSysId(v.(*string)) }, func() (interface{}, error) { return model.GetSysId() }, ptr("sys-id")},
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, ptr("app-name")},
		{"Version", func(v interface{}) error { return model.setVersion(v.(*string)) }, func() (interface{}, error) { return model.GetVersion() }, ptr("1.0.0")},
		{"Description", func(v interface{}) error { return model.setDescription(v.(*string)) }, func() (interface{}, error) { return model.GetDescription() }, ptr("app-desc")},
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, ptr("actual-app-name")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
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
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, ptr("app-name")},
		{"ComponentName", func(v interface{}) error { return model.setComponentName(v.(*string)) }, func() (interface{}, error) { return model.GetComponentName() }, ptr("comp-name")},
		{"Data", func(v interface{}) error { return model.setData(v.(*string)) }, func() (interface{}, error) { return model.GetData() }, ptr("base64-data")},
		{"Format", func(v interface{}) error { return model.setFormat(v.(*string)) }, func() (interface{}, error) { return model.GetFormat() }, ptr("json")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestCreateSharedLibraryComponentApplicationFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateSharedLibraryComponentApplicationFromDiscriminatorValue(nil)
	assert.NoError(t, err)
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
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, ptr("app-name")},
		{"ComponentName", func(v interface{}) error { return model.setComponentName(v.(*string)) }, func() (interface{}, error) { return model.GetComponentName() }, ptr("comp-name")},
		{"Vars", func(v interface{}) error { return model.setVars(v) }, func() (interface{}, error) { return model.GetVars() }, "vars"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
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
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, ptr("app-name")},
		{"CollectionName", func(v interface{}) error { return model.setCollectionName(v.(*string)) }, func() (interface{}, error) { return model.GetCollectionName() }, ptr("coll-name")},
		{"Data", func(v interface{}) error { return model.setData(v.(*string)) }, func() (interface{}, error) { return model.GetData() }, ptr("base64-data")},
		{"Format", func(v interface{}) error { return model.setFormat(v.(*string)) }, func() (interface{}, error) { return model.GetFormat() }, ptr("json")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestDeployableUpdateRequestModel_GettersSetters(t *testing.T) {
	model := NewDeployableUpdateRequest()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, ptr("app-name")},
		{"DeployableName", func(v interface{}) error { return model.setDeployableName(v.(*string)) }, func() (interface{}, error) { return model.GetDeployableName() }, ptr("deploy-name")},
		{"Data", func(v interface{}) error { return model.setData(v.(*string)) }, func() (interface{}, error) { return model.GetData() }, ptr("new-data")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
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
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, ptr("app-name")},
		{"SharedComponentName", func(v interface{}) error { return model.setSharedComponentName(v.(*string)) }, func() (interface{}, error) { return model.GetSharedComponentName() }, ptr("shared-name")},
		{"Data", func(v interface{}) error { return model.setData(v.(*string)) }, func() (interface{}, error) { return model.GetData() }, ptr("new-data")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestCreateExportStatusResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateExportStatusResultFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateComponentVarsUploadRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateComponentVarsUploadRequestFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateCollectionUploadRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateCollectionUploadRequestFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateDeployableUpdateRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateDeployableUpdateRequestFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateSharedComponentUpdateRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateSharedComponentUpdateRequestFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}
