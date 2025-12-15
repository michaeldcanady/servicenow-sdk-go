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

				elementValue, err := NewElementValue(val)

				assert.Nil(t, err)
				assert.NotNil(t, elementValue)
				assert.IsType(t, &ElementValue{}, elementValue)
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
				assert.IsType(t, &ElementValue{}, elementVal)
				assert.Nil(t, (elementVal.(*ElementValue)).val)
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

				value := &ElementValue{}

				err := value.Serialize(writer)

				assert.Equal(t, errors.New("Serialize is not supported"), err)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				writer := mocking.NewMockSerializationWriter()

				value := (*ElementValue)(nil)
				err := value.Serialize(writer)

				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementValueModel_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				model := &ElementValue{}
				serializers := model.GetFieldDeserializers()

				assert.Len(t, serializers, 0)
			},
		},
	}

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
				val := &ElementValue{val: nil}

				assert.True(t, val.IsNil())
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				val := (*ElementValue)(nil)

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

				model := &ElementValue{val: nil}

				err := model.setValue(value)

				assert.Nil(t, err)
				assert.Equal(t, value, model.val)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := (*ElementValue)(nil)

				err := model.setValue(value)

				assert.Nil(t, err)
			},
		},
		{
			name: "Non-pointer value",
			test: func(t *testing.T) {
				value := "test"

				model := &ElementValue{val: nil}

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

				model := &ElementValue{val: value}

				ret, err := model.GetStringValue()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer(true)

				model := &ElementValue{val: value}

				ret, err := model.GetStringValue()

				assert.Equal(t, errors.New("cannot convert 'true' to type string"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValue)(nil)

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

				model := &ElementValue{val: value}

				ret, err := model.GetBoolValue()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValue{val: value}

				ret, err := model.GetBoolValue()

				assert.Equal(t, errors.New("cannot convert 'test' to type bool"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValue)(nil)

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

				model := &ElementValue{val: value}

				ret, err := model.GetInt8Value()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValue{val: value}

				ret, err := model.GetInt8Value()

				assert.Equal(t, errors.New("cannot convert 'test' to type int8"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValue)(nil)

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

				model := &ElementValue{val: value}

				ret, err := model.GetByteValue()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValue{val: value}

				ret, err := model.GetByteValue()

				assert.Equal(t, errors.New("cannot convert 'test' to type uint8"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValue)(nil)

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

				model := &ElementValue{val: value}

				ret, err := model.GetFloat32Value()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValue{val: value}

				ret, err := model.GetFloat32Value()

				assert.Equal(t, errors.New("cannot convert 'test' to type float32"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValue)(nil)

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

				model := &ElementValue{val: value}

				ret, err := model.GetFloat64Value()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValue{val: value}

				ret, err := model.GetFloat64Value()

				assert.Equal(t, errors.New("cannot convert 'test' to type float64"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValue)(nil)

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

				model := &ElementValue{val: value}

				ret, err := model.GetInt32Value()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValue{val: value}

				ret, err := model.GetInt32Value()

				assert.Equal(t, errors.New("cannot convert 'test' to type int32"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValue)(nil)

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

				model := &ElementValue{val: value}

				ret, err := model.GetInt64Value()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				value := internal.ToPointer("test")

				model := &ElementValue{val: value}

				ret, err := model.GetInt64Value()

				assert.Equal(t, errors.New("cannot convert 'test' to type int64"), err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValue)(nil)

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

func TestElementValueModel_GetEnumValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Valid Parser",
			test: func(t *testing.T) {
				strct := mocking.NewMockEnumFactory()
				strct.On("Factory", "value").Return(0, nil)

				model := &ElementValue{val: "value"}
				enum, err := model.GetEnumValue(strct.Factory)

				assert.Nil(t, err)
				assert.Equal(t, 0, enum)
			},
		},
		{
			name: "Nil Parser",
			test: func(t *testing.T) {
				model := &ElementValue{val: "value"}
				enum, err := model.GetEnumValue(nil)

				assert.Equal(t, errors.New("parser is nil"), err)
				assert.Nil(t, enum)
			},
		},
		{
			name: "Parsing Error",
			test: func(t *testing.T) {
				strct := mocking.NewMockEnumFactory()

				model := &ElementValue{val: 1}
				enum, err := model.GetEnumValue(strct.Factory)

				assert.Equal(t, errors.New("cannot convert '1' to type string"), err)
				assert.Nil(t, enum)
			},
		},
		{
			name: "Empty String",
			test: func(t *testing.T) {
				strct := mocking.NewMockEnumFactory()

				model := &ElementValue{val: ""}
				enum, err := model.GetEnumValue(strct.Factory)

				assert.Nil(t, err)
				assert.Nil(t, enum)
			},
		},
		{
			name: "Nil Model",
			test: func(t *testing.T) {
				strct := mocking.NewMockEnumFactory()

				model := (*ElementValue)(nil)
				enum, err := model.GetEnumValue(strct.Factory)

				assert.Nil(t, err)
				assert.Nil(t, enum)
			},
		},
	}

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
				value := []interface{}{internal.ToPointer(true), internal.ToPointer(false), internal.ToPointer(true)}

				model, err := NewElementValue(value)
				assert.Nil(t, err)

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

func TestElementValueModel_getPrimitiveValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful - PrimitiveBool",
			test: func(t *testing.T) {
				value := internal.ToPointer(true)
				model := &ElementValue{val: value}

				val, err := model.getPrimitiveValue(PrimitiveBool)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveByte",
			test: func(t *testing.T) {
				value := internal.ToPointer(byte(1))
				model := &ElementValue{val: value}

				val, err := model.getPrimitiveValue(PrimitiveByte)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveFloat32",
			test: func(t *testing.T) {
				value := internal.ToPointer(float32(1))
				model := &ElementValue{val: value}

				val, err := model.getPrimitiveValue(PrimitiveFloat32)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveFloat64",
			test: func(t *testing.T) {
				value := internal.ToPointer(float64(1))
				model := &ElementValue{val: value}

				val, err := model.getPrimitiveValue(PrimitiveFloat64)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveInt32",
			test: func(t *testing.T) {
				value := internal.ToPointer(int32(1))
				model := &ElementValue{val: value}

				val, err := model.getPrimitiveValue(PrimitiveInt32)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveInt64",
			test: func(t *testing.T) {
				value := internal.ToPointer(int64(1))
				model := &ElementValue{val: value}

				val, err := model.getPrimitiveValue(PrimitiveInt64)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveInt8",
			test: func(t *testing.T) {
				value := internal.ToPointer(int8(1))
				model := &ElementValue{val: value}

				val, err := model.getPrimitiveValue(PrimitiveInt8)

				assert.Nil(t, err)
				assert.Equal(t, value, val)
			},
		},
		{
			name: "Successful - PrimitiveString",
			test: func(t *testing.T) {
				value := "test"
				model := &ElementValue{val: value}

				val, err := model.getPrimitiveValue(PrimitiveString)

				assert.Nil(t, err)
				assert.Equal(t, &value, val)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValue)(nil)

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

				model := &ElementValue{val: value}

				ret, err := model.GetRawValue()

				assert.Nil(t, err)
				assert.Equal(t, value, ret)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ElementValue)(nil)

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
