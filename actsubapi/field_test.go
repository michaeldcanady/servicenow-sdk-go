package actsubapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestNewField(t *testing.T) {
	instance := NewField()
	assert.NotNil(t, instance)
}

func TestCreateFieldFromDiscriminatorValue(t *testing.T) {
	instance, err := CreateFieldFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, instance)
}

func TestField_Serialize(t *testing.T) {
	instance := NewField()
	writer := &mocking.MockSerializationWriter{}

	err := instance.Serialize(writer)
	assert.NoError(t, err)
}

func TestField_GetFieldDeserializers(t *testing.T) {
	instance := NewField()
	deserializers := instance.GetFieldDeserializers()
	assert.NotNil(t, deserializers)
}

func TestField_GettersSetters(t *testing.T) {
	instance := NewField()
	val := "test-value"

	tests := []struct {
		name   string
		setter func(*string) error
		getter func() (*string, error)
	}{
		{"DeepLinkToSubObject", instance.SetDeepLinkToSubObject, instance.GetDeepLinkToSubObject},
		{"DisplayAsTimeAgo", instance.SetDisplayAsTimeAgo, instance.GetDisplayAsTimeAgo},
		{"Label", instance.SetLabel, instance.GetLabel},
		{"ShowLabel", instance.SetShowLabel, instance.GetShowLabel},
		{"Type", instance.SetType, instance.GetType},
		{"Value", instance.SetValue, instance.GetValue},
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
}
