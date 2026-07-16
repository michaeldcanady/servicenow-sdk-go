package cdmapplicationsapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestUploadStatusOutputModel_GettersSetters(t *testing.T) {
	model := NewUploadStatusOutput()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"SysId", func(v interface{}) error { return model.setSysId(v.(*string)) }, func() (interface{}, error) { return model.GetSysId() }, internal.ToPointer("sys-id")},
		{"Number", func(v interface{}) error { return model.setNumber(v.(*string)) }, func() (interface{}, error) { return model.GetNumber() }, internal.ToPointer("123")},
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
		{"Type", func(v interface{}) error { return model.setType(v.(*string)) }, func() (interface{}, error) { return model.GetType() }, internal.ToPointer("upload")},
		{"State", func(v interface{}) error { return model.setState(v.(*string)) }, func() (interface{}, error) { return model.GetState() }, internal.ToPointer("completed")},
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
		{"SysId", func(v interface{}) error { return model.setSysId(v.(*string)) }, func() (interface{}, error) { return model.GetSysId() }, internal.ToPointer("sys-id")},
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, internal.ToPointer("export-name")},
		{"State", func(v interface{}) error { return model.setState(v.(*string)) }, func() (interface{}, error) { return model.GetState() }, internal.ToPointer("active")},
		{"Status", func(v interface{}) error { return model.setStatus(v.(*string)) }, func() (interface{}, error) { return model.GetStatus() }, internal.ToPointer("success")},
		{"Message", func(v interface{}) error { return model.setMessage(v.(*string)) }, func() (interface{}, error) { return model.GetMessage() }, internal.ToPointer("export successful")},
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
		{"State", func(v interface{}) error { return model.setState(v.(*string)) }, func() (interface{}, error) { return model.GetState() }, internal.ToPointer("completed")},
		{"Status", func(v interface{}) error { return model.setStatus(v.(*string)) }, func() (interface{}, error) { return model.GetStatus() }, internal.ToPointer("success")},
		{"Message", func(v interface{}) error { return model.setMessage(v.(*string)) }, func() (interface{}, error) { return model.GetMessage() }, internal.ToPointer("status message")},
		{"Progress", func(v interface{}) error { return model.setProgress(v.(*string)) }, func() (interface{}, error) { return model.GetProgress() }, internal.ToPointer("100")},
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
		{"SysId", func(v interface{}) error { return model.setSysId(v.(*string)) }, func() (interface{}, error) { return model.GetSysId() }, internal.ToPointer("sys-id")},
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, internal.ToPointer("app-name")},
		{"Version", func(v interface{}) error { return model.setVersion(v.(*string)) }, func() (interface{}, error) { return model.GetVersion() }, internal.ToPointer("1.0.0")},
		{"Description", func(v interface{}) error { return model.setDescription(v.(*string)) }, func() (interface{}, error) { return model.GetDescription() }, internal.ToPointer("app-desc")},
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, internal.ToPointer("actual-app-name")},
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
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, internal.ToPointer("app-name")},
		{"ComponentName", func(v interface{}) error { return model.setComponentName(v.(*string)) }, func() (interface{}, error) { return model.GetComponentName() }, internal.ToPointer("comp-name")},
		{"Data", func(v interface{}) error { return model.setData(v.(*string)) }, func() (interface{}, error) { return model.GetData() }, internal.ToPointer("base64-data")},
		{"Format", func(v interface{}) error { return model.setFormat(v.(*string)) }, func() (interface{}, error) { return model.GetFormat() }, internal.ToPointer("json")},
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
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, internal.ToPointer("app-name")},
		{"ComponentName", func(v interface{}) error { return model.setComponentName(v.(*string)) }, func() (interface{}, error) { return model.GetComponentName() }, internal.ToPointer("comp-name")},
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
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, internal.ToPointer("app-name")},
		{"CollectionName", func(v interface{}) error { return model.setCollectionName(v.(*string)) }, func() (interface{}, error) { return model.GetCollectionName() }, internal.ToPointer("coll-name")},
		{"Data", func(v interface{}) error { return model.setData(v.(*string)) }, func() (interface{}, error) { return model.GetData() }, internal.ToPointer("base64-data")},
		{"Format", func(v interface{}) error { return model.setFormat(v.(*string)) }, func() (interface{}, error) { return model.GetFormat() }, internal.ToPointer("json")},
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
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, internal.ToPointer("app-name")},
		{"DeployableName", func(v interface{}) error { return model.setDeployableName(v.(*string)) }, func() (interface{}, error) { return model.GetDeployableName() }, internal.ToPointer("deploy-name")},
		{"Data", func(v interface{}) error { return model.setData(v.(*string)) }, func() (interface{}, error) { return model.GetData() }, internal.ToPointer("new-data")},
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
		{"AppName", func(v interface{}) error { return model.setAppName(v.(*string)) }, func() (interface{}, error) { return model.GetAppName() }, internal.ToPointer("app-name")},
		{"SharedComponentName", func(v interface{}) error { return model.setSharedComponentName(v.(*string)) }, func() (interface{}, error) { return model.GetSharedComponentName() }, internal.ToPointer("shared-name")},
		{"Data", func(v interface{}) error { return model.setData(v.(*string)) }, func() (interface{}, error) { return model.GetData() }, internal.ToPointer("new-data")},
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
