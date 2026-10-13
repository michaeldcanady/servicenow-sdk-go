// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package activitysubscriptionsapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewActivity(t *testing.T) {
	instance := NewActivity()
	assert.NotNil(t, instance)
}

func TestCreateActivityFromDiscriminatorValue(t *testing.T) {
	instance, err := CreateActivityFromDiscriminatorValue(nil)
	require.NoError(t, err)
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
			require.NoError(t, err)
			got, err := tt.getter()
			require.NoError(t, err)
			assert.Equal(t, &val, got)
		})
	}

	fields := []*Field{NewField()}
	err := instance.SetContentFields(fields)
	require.NoError(t, err)
	resFields, err := instance.GetContentFields()
	require.NoError(t, err)
	assert.Equal(t, fields, resFields)

	err = instance.SetSubheaderFields(fields)
	require.NoError(t, err)
	resFields, err = instance.GetSubheaderFields()
	require.NoError(t, err)
	assert.Equal(t, fields, resFields)
}
