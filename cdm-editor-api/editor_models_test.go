package cdmeditorapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptr[T any](v T) *T {
	return &v
}

func TestNodeResultModel_GettersSetters(t *testing.T) {
	model := NewNodeResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"SysId", func(v interface{}) error { return model.setSysId(v.(*string)) }, func() (interface{}, error) { return model.GetSysId() }, ptr("sys-id")},
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, ptr("node-name")},
		{"Type", func(v interface{}) error { return model.setType(v.(*string)) }, func() (interface{}, error) { return model.GetType() }, ptr("folder")},
		{"Value", func(v interface{}) error { return model.setValue(v.(*string)) }, func() (interface{}, error) { return model.GetValue() }, ptr("node-value")},
		{"Parent", func(v interface{}) error { return model.setParent(v.(*string)) }, func() (interface{}, error) { return model.GetParent() }, ptr("parent-id")},
		{"CdmId", func(v interface{}) error { return model.setCdmId(v.(*string)) }, func() (interface{}, error) { return model.GetCdmId() }, ptr("cdm-id")},
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

func TestValidationResultModel_GettersSetters(t *testing.T) {
	model := NewValidationResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"Status", func(v interface{}) error { return model.setStatus(v.(*string)) }, func() (interface{}, error) { return model.GetStatus() }, ptr("success")},
		{"Errors", func(v interface{}) error { return model.setErrors(v) }, func() (interface{}, error) { return model.GetErrors() }, "no errors"},
		{"Warnings", func(v interface{}) error { return model.setWarnings(v) }, func() (interface{}, error) { return model.GetWarnings() }, "no warnings"},
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

func TestNodeCreateRequestModel_GettersSetters(t *testing.T) {
	model := NewNodeCreateRequest()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, ptr("new-node")},
		{"Type", func(v interface{}) error { return model.setType(v.(*string)) }, func() (interface{}, error) { return model.GetType() }, ptr("file")},
		{"ParentId", func(v interface{}) error { return model.setParentId(v.(*string)) }, func() (interface{}, error) { return model.GetParentId() }, ptr("parent-id")},
		{"CdmId", func(v interface{}) error { return model.setCdmId(v.(*string)) }, func() (interface{}, error) { return model.GetCdmId() }, ptr("cdm-id")},
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

func TestNodeUpdateRequestModel_GettersSetters(t *testing.T) {
	model := NewNodeUpdateRequest()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, ptr("updated-name")},
		{"Value", func(v interface{}) error { return model.setValue(v.(*string)) }, func() (interface{}, error) { return model.GetValue() }, ptr("updated-value")},
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

func TestCreateNodeResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateNodeResultFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateValidationResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateValidationResultFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateNodeCreateRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateNodeCreateRequestFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateNodeUpdateRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateNodeUpdateRequestFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}
