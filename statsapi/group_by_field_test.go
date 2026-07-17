package statsapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewGroupByField(t *testing.T) {
	field := NewGroupByField()
	assert.NotNil(t, field)
}

func TestCreateGroupByFieldFromDiscriminatorValue(t *testing.T) {
	field, err := CreateGroupByFieldFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, field)
}

func TestGroupByField_GettersAndSetters(t *testing.T) {
	tests := []struct {
		name   string
		setter func(*GroupByFieldModel, *string) error
		getter func(*GroupByFieldModel) (*string, error)
	}{
		{
			name:   "Field",
			setter: func(m *GroupByFieldModel, v *string) error { return m.setField(v) },
			getter: func(m *GroupByFieldModel) (*string, error) { return m.GetField() },
		},
		{
			name:   "Value",
			setter: func(m *GroupByFieldModel, v *string) error { return m.setValue(v) },
			getter: func(m *GroupByFieldModel) (*string, error) { return m.GetValue() },
		},
		{
			name:   "DisplayValue",
			setter: func(m *GroupByFieldModel, v *string) error { return m.setDisplayValue(v) },
			getter: func(m *GroupByFieldModel) (*string, error) { return m.GetDisplayValue() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewGroupByField()
			val := internal.ToPointer("test-value")

			err := tt.setter(m, val)
			assert.NoError(t, err)

			res, err := tt.getter(m)
			assert.NoError(t, err)
			assert.Equal(t, val, res)
		})
	}
}

func TestGroupByField_Serialize(t *testing.T) {
	writer := mocking.NewMockSerializationWriter()
	writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)

	field := NewGroupByField()
	_ = field.setField(internal.ToPointer("priority"))
	_ = field.setValue(internal.ToPointer("1"))

	err := field.Serialize(writer)
	assert.NoError(t, err)

	var nilField *GroupByFieldModel
	err = nilField.Serialize(writer)
	assert.NoError(t, err)
}

func TestGroupByField_GetFieldDeserializers(t *testing.T) {
	field := NewGroupByField()
	deser := field.GetFieldDeserializers()
	assert.NotNil(t, deser[groupByFieldFieldKey])
	assert.NotNil(t, deser[groupByFieldValueKey])
	assert.NotNil(t, deser[groupByFieldDisplayValueKey])
}
