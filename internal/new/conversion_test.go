package internal

import (
	"errors"
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDereference(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "String pointer",
			test: func(t *testing.T) {
				input := ToPointer("testing")
				output := Dereference(reflect.ValueOf(input))

				assert.Equal(t, interface{}("testing"), output.Interface())
			},
		},
		{
			name: "String pointer pointer",
			test: func(t *testing.T) {
				input := ToPointer(ToPointer("testing"))
				output := Dereference(reflect.ValueOf(input))

				assert.Equal(t, interface{}("testing"), output.Interface())
			},
		},
		{
			name: "String",
			test: func(t *testing.T) {
				input := "testing"
				output := Dereference(reflect.ValueOf(input))

				assert.Equal(t, interface{}("testing"), output.Interface())
			},
		},
		{
			name: "nil",
			test: func(t *testing.T) {
				input := (*string)(nil)
				output := Dereference(reflect.ValueOf(input))

				assert.Equal(t, "", output.Interface())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestIsNumericKind(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Int8",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(int8(1)).Kind())

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Uint8",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(uint8(1)).Kind())

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Int16",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(int16(1)).Kind())

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Uint16",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(uint16(1)).Kind())

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Int32",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(int32(1)).Kind())

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Uint32",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(uint32(1)).Kind())

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Int64",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(int64(1)).Kind())

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Uint64",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(uint64(1)).Kind())

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Float32",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(float32(1.00)).Kind())

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Float64",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(float64(1.00)).Kind())

				assert.True(t, isNumeric)
			},
		},
		{
			name: "String",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf("string").Kind())

				assert.False(t, isNumeric)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestConvertNumeric(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Int8 to Int8",
			test: func(t *testing.T) {
				input := reflect.ValueOf(int8(0))
				desiredType := reflect.TypeOf(int8(0))
				result, err := convertNumeric(input, desiredType)

				assert.Nil(t, err)
				assert.Equal(t, int8(0), result)
			},
		},
		{
			name: "String to Int8",
			test: func(t *testing.T) {
				input := reflect.ValueOf("string")
				desiredType := reflect.TypeOf(int8(0))
				result, err := convertNumeric(input, desiredType)

				assert.Equal(t, errors.New("string is non-numeric"), err)
				assert.Nil(t, result)
			},
		},
		{
			name: "Int16 to Int8",
			test: func(t *testing.T) {
				input := reflect.ValueOf(math.MaxInt16)
				desiredType := reflect.TypeOf(int8(0))
				result, err := convertNumeric(input, desiredType)

				assert.Equal(t, errors.New("overflow or incompatible decimal converting to int8"), err)
				assert.Nil(t, result)
			},
		},
		{
			name: "Int16 to Int",
			test: func(t *testing.T) {
				input := reflect.ValueOf(math.MaxInt16)
				desiredType := reflect.TypeOf(int(0))
				result, err := convertNumeric(input, desiredType)

				assert.Nil(t, err)
				assert.Equal(t, int(32767), result)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestConvertValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Int8 to Int16",
			test: func(t *testing.T) {
				srcValue := reflect.ValueOf(int8(0))
				result, err := convertValue(srcValue, reflect.TypeOf(int16(0)))

				assert.Nil(t, err)
				assert.Equal(t, int16(0), result.Interface())
			},
		},
		{
			name: "Int String to Int16",
			test: func(t *testing.T) {
				srcValue := reflect.ValueOf("0")
				result, err := convertValue(srcValue, reflect.TypeOf(int16(0)))

				assert.Nil(t, err)
				assert.Equal(t, int16(0), result.Interface())
			},
		},
		{
			name: "Word String to Int16",
			test: func(t *testing.T) {
				srcValue := reflect.ValueOf("not number")
				result, err := convertValue(srcValue, reflect.TypeOf(int16(0)))

				assert.Equal(t, errors.New("unable to convert not number to float: strconv.ParseFloat: parsing \"not number\": invalid syntax"), err)
				assert.Equal(t, reflect.Value{}, result)
			},
		},
		{
			name: "Int16 to Int String",
			test: func(t *testing.T) {
				srcValue := reflect.ValueOf(int16(0))
				result, err := convertValue(srcValue, reflect.TypeOf(""))

				assert.Nil(t, err)
				assert.Equal(t, "0", result.Interface())
			},
		},
		{
			name: "String to String",
			test: func(t *testing.T) {
				srcValue := reflect.ValueOf("0")
				result, err := convertValue(srcValue, reflect.TypeOf(""))

				assert.Nil(t, err)
				assert.Equal(t, "0", result.Interface())
			},
		},
		{
			name: "Bool to String",
			test: func(t *testing.T) {
				srcValue := reflect.ValueOf(true)
				result, err := convertValue(srcValue, reflect.TypeOf(""))

				assert.Equal(t, errors.New("unsupported conversion: bool to string"), err)
				assert.Equal(t, reflect.Value{}, result)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Non-pointer output",
			test: func(t *testing.T) {
				input := ""
				output := ""

				err := Convert(input, output)

				assert.Equal(t, errors.New("output must be a non-nil pointer"), err)
			},
		},
		{
			name: "Nil input",
			test: func(t *testing.T) {
				output := ToPointer("random words")

				err := Convert(nil, output)

				assert.Nil(t, err)
				assert.Equal(t, "", *output)
			},
		},
		{
			name: "Nil output",
			test: func(t *testing.T) {
				err := Convert(nil, nil)

				assert.Equal(t, errors.New("output cannot be nil"), err)
			},
		},
		{
			name: "Int8 to Int16",
			test: func(t *testing.T) {
				input := int8(10)
				output := ToPointer(int16(0))

				err := Convert(input, output)

				assert.Nil(t, err)
				assert.Equal(t, int16(10), *output)
			},
		},
		{
			name: "String pointer",
			test: func(t *testing.T) {
				var input interface{} = ToPointer("test")
				var output *string

				err := Convert(input, &output)

				assert.Nil(t, err)
				assert.Equal(t, ToPointer("test"), output)
			},
		},
		{
			name: "int8",
			test: func(t *testing.T) {
				var input interface{} = int8(8)
				var output int8

				err := Convert(input, &output)

				assert.Nil(t, err)
				assert.Equal(t, int8(8), output)
			},
		},
		{
			name: "String pointer pointer",
			test: func(t *testing.T) {
				var input interface{} = ToPointer(ToPointer("test"))
				var output **string

				err := Convert(input, &output)

				assert.Nil(t, err)
				assert.Equal(t, ToPointer(ToPointer("test")), output)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
