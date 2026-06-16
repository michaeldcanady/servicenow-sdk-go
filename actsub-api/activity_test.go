package actsubapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestNewActivity(t *testing.T) {
	instance := NewActivity()
	assert.NotNil(t, instance)
}

func TestCreateActivityFromDiscriminatorValue(t *testing.T) {
	instance, err := CreateActivityFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, instance)
}

func TestActivity_Serialize(t *testing.T) {
	instance := NewActivity()
	writer := &mocking.MockSerializationWriter{}

	err := instance.Serialize(writer)
	assert.NoError(t, err)
}

func TestActivity_GetFieldDeserializers(t *testing.T) {
	instance := NewActivity()
	deserializers := instance.GetFieldDeserializers()
	assert.NotNil(t, deserializers)
}

func TestActivity_GettersSetters(t *testing.T) {
	instance := NewActivity()
	val := "test-value"

	tests := []struct {
		name   string
		setter func(*string) error
		getter func() (*string, error)
	}{
		{"ActivityTypeID", instance.SetActivityTypeID, instance.GetActivityTypeID},
		{"SourceTableName", instance.SetSourceTableName, instance.GetSourceTableName},
		{"SubObjectTableName", instance.SetSubObjectTableName, instance.GetSubObjectTableName},
		{"SubObjectSysID", instance.SetSubObjectSysID, instance.GetSubObjectSysID},
		{"Title", instance.SetTitle, instance.GetTitle},
		{"SysIDKey", instance.SetSysIDKey, instance.GetSysIDKey},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(&val)
			assert.NoError(t, err)
			got, err := tt.getter()
			assert.NoError(t, err)
			assert.Equal(t, &val, got)
		})
	}

	fields := []*Field{NewField()}
	err := instance.SetContentFields(fields)
	assert.NoError(t, err)
	resFields, err := instance.GetContentFields()
	assert.NoError(t, err)
	assert.Equal(t, fields, resFields)

	err = instance.SetSubheaderFields(fields)
	assert.NoError(t, err)
	resFields, err = instance.GetSubheaderFields()
	assert.NoError(t, err)
	assert.Equal(t, fields, resFields)
}
