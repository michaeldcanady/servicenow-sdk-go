package cdmeditorapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestNodeResultModel_GettersSetters(t *testing.T) {
	model := NewNodeResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"SysId", func(v interface{}) error { return model.setSysId(v.(*string)) }, func() (interface{}, error) { return model.GetSysId() }, internal.ToPointer("sys-id")},
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, internal.ToPointer("node-name")},
		{"Type", func(v interface{}) error { return model.setType(v.(*string)) }, func() (interface{}, error) { return model.GetType() }, internal.ToPointer("folder")},
		{"Value", func(v interface{}) error { return model.setValue(v.(*string)) }, func() (interface{}, error) { return model.GetValue() }, internal.ToPointer("node-value")},
		{"Parent", func(v interface{}) error { return model.setParent(v.(*string)) }, func() (interface{}, error) { return model.GetParent() }, internal.ToPointer("parent-id")},
		{"CdmId", func(v interface{}) error { return model.setCdmId(v.(*string)) }, func() (interface{}, error) { return model.GetCdmId() }, internal.ToPointer("cdm-id")},
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
		{"Status", func(v interface{}) error { return model.setStatus(v.(*string)) }, func() (interface{}, error) { return model.GetStatus() }, internal.ToPointer("success")},
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
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, internal.ToPointer("new-node")},
		{"Type", func(v interface{}) error { return model.setType(v.(*string)) }, func() (interface{}, error) { return model.GetType() }, internal.ToPointer("file")},
		{"ParentId", func(v interface{}) error { return model.setParentId(v.(*string)) }, func() (interface{}, error) { return model.GetParentId() }, internal.ToPointer("parent-id")},
		{"CdmId", func(v interface{}) error { return model.setCdmId(v.(*string)) }, func() (interface{}, error) { return model.GetCdmId() }, internal.ToPointer("cdm-id")},
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
		{"Name", func(v interface{}) error { return model.setName(v.(*string)) }, func() (interface{}, error) { return model.GetName() }, internal.ToPointer("updated-name")},
		{"Value", func(v interface{}) error { return model.setValue(v.(*string)) }, func() (interface{}, error) { return model.GetValue() }, internal.ToPointer("updated-value")},
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
