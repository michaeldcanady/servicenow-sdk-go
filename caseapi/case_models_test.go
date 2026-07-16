package caseapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestReferenceModel_GettersSetters(t *testing.T) {
	model := NewReference()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"Link", func(v interface{}) error { return model.setLink(v.(*string)) }, func() (interface{}, error) { return model.GetLink() }, internal.ToPointer("https://example.com")},
		{"Value", func(v interface{}) error { return model.setValue(v.(*string)) }, func() (interface{}, error) { return model.GetValue() }, internal.ToPointer("test-value")},
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

func TestCaseResultModel_GettersSetters(t *testing.T) {
	model := NewCaseResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"SysId", func(v interface{}) error { return model.setSysId(v.(*string)) }, func() (interface{}, error) { return model.GetSysId() }, internal.ToPointer("sys-id")},
		{"Number", func(v interface{}) error { return model.setNumber(v.(*string)) }, func() (interface{}, error) { return model.GetNumber() }, internal.ToPointer("CASE001")},
		{"ShortDescription", func(v interface{}) error { return model.setShortDescription(v.(*string)) }, func() (interface{}, error) { return model.GetShortDescription() }, internal.ToPointer("Short desc")},
		{"Description", func(v interface{}) error { return model.setDescription(v.(*string)) }, func() (interface{}, error) { return model.GetDescription() }, internal.ToPointer("Full desc")},
		{"State", func(v interface{}) error { return model.setState(v.(*string)) }, func() (interface{}, error) { return model.GetState() }, internal.ToPointer("10")},
		{"Priority", func(v interface{}) error { return model.setPriority(v.(*string)) }, func() (interface{}, error) { return model.GetPriority() }, internal.ToPointer("1")},
		{"Category", func(v interface{}) error { return model.setCategory(v.(*string)) }, func() (interface{}, error) { return model.GetCategory() }, internal.ToPointer("inquiry")},
		{"AssignmentGroup", func(v interface{}) error { return model.setAssignmentGroup(v.(Reference)) }, func() (interface{}, error) { return model.GetAssignmentGroup() }, NewReference()},
		{"AssignedTo", func(v interface{}) error { return model.setAssignedTo(v.(Reference)) }, func() (interface{}, error) { return model.GetAssignedTo() }, NewReference()},
		{"Contact", func(v interface{}) error { return model.setContact(v.(Reference)) }, func() (interface{}, error) { return model.GetContact() }, NewReference()},
		{"Account", func(v interface{}) error { return model.setAccount(v.(Reference)) }, func() (interface{}, error) { return model.GetAccount() }, NewReference()},
		{"SysCreatedOn", func(v interface{}) error { return model.setSysCreatedOn(v.(*string)) }, func() (interface{}, error) { return model.GetSysCreatedOn() }, internal.ToPointer("2023-01-01 12:00:00")},
		{"SysUpdatedOn", func(v interface{}) error { return model.setSysUpdatedOn(v.(*string)) }, func() (interface{}, error) { return model.GetSysUpdatedOn() }, internal.ToPointer("2023-01-01 13:00:00")},
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

func TestActivitiesResultModel_GettersSetters(t *testing.T) {
	model := NewActivitiesResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"SysId", func(v interface{}) error { return model.setSysId(v.(*string)) }, func() (interface{}, error) { return model.GetSysId() }, internal.ToPointer("sys-id")},
		{"Type", func(v interface{}) error { return model.setType(v.(*string)) }, func() (interface{}, error) { return model.GetType() }, internal.ToPointer("work_notes")},
		{"Value", func(v interface{}) error { return model.setValue(v.(*string)) }, func() (interface{}, error) { return model.GetValue() }, internal.ToPointer("test value")},
		{"User", func(v interface{}) error { return model.setUser(v.(*string)) }, func() (interface{}, error) { return model.GetUser() }, internal.ToPointer("admin")},
		{"SysCreatedOn", func(v interface{}) error { return model.setSysCreatedOn(v.(*string)) }, func() (interface{}, error) { return model.GetSysCreatedOn() }, internal.ToPointer("2023-01-01 12:00:00")},
		{"FieldName", func(v interface{}) error { return model.setFieldName(v.(*string)) }, func() (interface{}, error) { return model.GetFieldName() }, internal.ToPointer("work_notes")},
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

func TestFieldValuesResultModel_GettersSetters(t *testing.T) {
	model := NewFieldValuesResult()

	tests := []struct {
		name   string
		setter func(val interface{}) error
		getter func() (interface{}, error)
		value  interface{}
	}{
		{"Label", func(v interface{}) error { return model.setLabel(v.(*string)) }, func() (interface{}, error) { return model.GetLabel() }, internal.ToPointer("label")},
		{"Value", func(v interface{}) error { return model.setValue(v.(*string)) }, func() (interface{}, error) { return model.GetValue() }, internal.ToPointer("value")},
		{"Sequence", func(v interface{}) error { return model.setSequence(v.(*int32)) }, func() (interface{}, error) { return model.GetSequence() }, internal.ToPointer(int32(1))},
		{"DependentValue", func(v interface{}) error { return model.setDependentValue(v.(*string)) }, func() (interface{}, error) { return model.GetDependentValue() }, internal.ToPointer("dep")},
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

func TestCreateReferenceFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateReferenceFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateCaseResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateCaseResultFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateActivitiesResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateActivitiesResultFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateFieldValuesResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateFieldValuesResultFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, parsable)
}
