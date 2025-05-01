package tableapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
)

func TestNewElementValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				val := "test"

				elementValue := NewElementValue(val)

				assert.NotNil(t, elementValue)
				assert.IsType(t, &ElementValueModel{}, elementValue)
				assert.Equal(t, val, elementValue.val)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestCreateElementValueFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				parseNode := mocking.NewMockParseNode()

				elementVal, err := CreateElementValueFromDiscriminatorValue(parseNode)

				assert.Nil(t, err)
				assert.NotNil(t, elementVal)
				assert.IsType(t, &ElementValueModel{}, elementVal)
				assert.Nil(t, (elementVal.(*ElementValueModel)).val)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementValueModel_Serialize(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				writer := mocking.NewMockSerializationWriter()

				value := &ElementValueModel{}

				err := value.Serialize(writer)

				assert.Equal(t, errors.New("Serialize is not supported"), err)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				writer := mocking.NewMockSerializationWriter()

				value := (*ElementValueModel)(nil)
				err := value.Serialize(writer)

				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestElementValueModel_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementValueModel_IsNil(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Nil value",
			test: func(t *testing.T) {
				val := &ElementValueModel{val: nil}

				assert.True(t, val.IsNil())
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				val := (*ElementValueModel)(nil)

				assert.True(t, val.IsNil())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementValueModel_setValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValueModel{val: nil}

				err := model.setValue(value)

				assert.Nil(t, err)
				assert.Equal(t, value, model.val)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := (*ElementValueModel)(nil)

				err := model.setValue(value)

				assert.Nil(t, err)
			},
		},
		{
			name: "Non-pointer value",
			test: func(t *testing.T) {
				value := "test"

				model := &ElementValueModel{val: nil}

				err := model.setValue(value)

				assert.Equal(t, errors.New("val is not a pointer"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementValueModel_GetStringValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValueModel{val: value}

				ret, err := model.GetStringValue()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer(true)

				model := &ElementValueModel{val: value}

				ret, err := model.GetStringValue()

				assert.Equal(t, errors.New("type '*bool' is not compatible with type string"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValueModel)(nil)

				ret, err := model.GetStringValue()

				assert.Nil(t, err)
				assert.Nil(t, ret)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementValueModel_GetBoolValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := internal.ToPointer(true)

				model := &ElementValueModel{val: value}

				ret, err := model.GetBoolValue()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValueModel{val: value}

				ret, err := model.GetBoolValue()

				assert.Equal(t, errors.New("type '*string' is not compatible with type bool"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValueModel)(nil)

				ret, err := model.GetBoolValue()

				assert.Nil(t, err)
				assert.Nil(t, ret)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementValueModel_GetInt8Value(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := internal.ToPointer(int8(7))

				model := &ElementValueModel{val: value}

				ret, err := model.GetInt8Value()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValueModel{val: value}

				ret, err := model.GetInt8Value()

				assert.Equal(t, errors.New("value 'test' is not compatible with type int8"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValueModel)(nil)

				ret, err := model.GetInt8Value()

				assert.Nil(t, err)
				assert.Nil(t, ret)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementValueModel_GetByteValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := internal.ToPointer(byte(7))

				model := &ElementValueModel{val: value}

				ret, err := model.GetByteValue()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValueModel{val: value}

				ret, err := model.GetByteValue()

				assert.Equal(t, errors.New("value 'test' is not compatible with type uint8"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValueModel)(nil)

				ret, err := model.GetByteValue()

				assert.Nil(t, err)
				assert.Nil(t, ret)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementValueModel_GetFloat32Value(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := internal.ToPointer(float32(7))

				model := &ElementValueModel{val: value}

				ret, err := model.GetFloat32Value()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValueModel{val: value}

				ret, err := model.GetFloat32Value()

				assert.Equal(t, errors.New("value 'test' is not compatible with type float32"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValueModel)(nil)

				ret, err := model.GetFloat32Value()

				assert.Nil(t, err)
				assert.Nil(t, ret)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementValueModel_GetFloat64Value(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := internal.ToPointer(float64(7))

				model := &ElementValueModel{val: value}

				ret, err := model.GetFloat64Value()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValueModel{val: value}

				ret, err := model.GetFloat64Value()

				assert.Equal(t, errors.New("value 'test' is not compatible with type float64"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValueModel)(nil)

				ret, err := model.GetFloat64Value()

				assert.Nil(t, err)
				assert.Nil(t, ret)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementValueModel_GetInt32Value(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := internal.ToPointer(int32(7))

				model := &ElementValueModel{val: value}

				ret, err := model.GetInt32Value()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValueModel{val: value}

				ret, err := model.GetInt32Value()

				assert.Equal(t, errors.New("value 'test' is not compatible with type int32"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValueModel)(nil)

				ret, err := model.GetInt32Value()

				assert.Nil(t, err)
				assert.Nil(t, ret)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementValueModel_GetInt64Value(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := internal.ToPointer(int64(7))

				model := &ElementValueModel{val: value}

				ret, err := model.GetInt64Value()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValueModel{val: value}

				ret, err := model.GetInt64Value()

				assert.Equal(t, errors.New("value 'test' is not compatible with type int64"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValueModel)(nil)

				ret, err := model.GetInt64Value()

				assert.Nil(t, err)
				assert.Nil(t, ret)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestElementValueModel_GetEnumValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestElementValueModel_GetCollectionOfPrimitiveValues(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := []interface{}{true, false, true}

				model := &ElementValueModel{val: value}

				ret, err := model.GetCollectionOfPrimitiveValues(PrimitiveBool)

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestElementValueModel_getPrimitiveValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful - PrimitiveBool",
			test: func(t *testing.T) {
				value := internal.ToPointer(true)
				model := &ElementValueModel{val: value}

				val, err := model.getPrimitiveValue(PrimitiveBool)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveByte",
			test: func(t *testing.T) {
				value := internal.ToPointer(byte(1))
				model := &ElementValueModel{val: value}

				val, err := model.getPrimitiveValue(PrimitiveByte)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveFloat32",
			test: func(t *testing.T) {
				value := internal.ToPointer(float32(1))
				model := &ElementValueModel{val: value}

				val, err := model.getPrimitiveValue(PrimitiveFloat32)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveFloat64",
			test: func(t *testing.T) {
				value := internal.ToPointer(float64(1))
				model := &ElementValueModel{val: value}

				val, err := model.getPrimitiveValue(PrimitiveFloat64)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveInt32",
			test: func(t *testing.T) {
				value := internal.ToPointer(int32(1))
				model := &ElementValueModel{val: value}

				val, err := model.getPrimitiveValue(PrimitiveInt32)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveInt64",
			test: func(t *testing.T) {
				value := internal.ToPointer(int64(1))
				model := &ElementValueModel{val: value}

				val, err := model.getPrimitiveValue(PrimitiveInt64)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveInt8",
			test: func(t *testing.T) {
				value := internal.ToPointer(int8(1))
				model := &ElementValueModel{val: value}

				val, err := model.getPrimitiveValue(PrimitiveInt8)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveString",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")
				model := &ElementValueModel{val: value}

				val, err := model.getPrimitiveValue(PrimitiveString)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveDateOnly",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")
				model := &ElementValueModel{val: value}

				val, err := model.getPrimitiveValue(PrimitiveDateOnly)

				assert.Equal(t, errors.New("unknown primitive dateonly"), err)
				assert.Nil(t, val)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValueModel)(nil)

				ret, err := model.getPrimitiveValue(PrimitiveDateOnly)

				assert.Nil(t, err)
				assert.Nil(t, ret)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementValueModel_GetRawValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := internal.ToPointer(7)

				model := &ElementValueModel{val: value}

				ret, err := model.GetRawValue()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValueModel)(nil)

				ret, err := model.GetRawValue()

				assert.Nil(t, err)
				assert.Nil(t, ret)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
