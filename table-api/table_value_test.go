package tableapi

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsNotNilError(err error) bool {
	return err != nil
}

type test[T any] struct {
	title     string
	value     interface{}
	expected  T
	expectErr bool
	Err       error
}

func (te test[T]) checkError(t *testing.T, err error) {
	if te.expectErr && err == nil {
		if te.Err != nil {
			assert.ErrorIs(t, err, te.Err)
		} else {
			assert.Error(t, err)
		}
	}

	if !te.expectErr && err != nil {
		assert.Nil(t, err)
	}
}

func TestTableValueToInt64(t *testing.T) {
	tests := []test[int64]{
		{"Int64", int64(42), 42, false, nil},
		{"Int32", int32(123), 123, false, nil},
		{"Float32", float32(3.14), 0, true, nil},
		{"String", "not an integer", 0, true, nil},
		{"Nil", nil, 0, true, nil},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result, err := tableValue.ToInt64()

			test.checkError(t, err)

			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestTableValueInt(t *testing.T) {
	tests := []test[int64]{
		{"Int64", int64(4142), 4142, false, nil},
		{"Int32", int32(123), 123, false, nil},
		{"Int16", int16(4870), 4870, false, nil},
		{"Int8", int8(98), 98, false, nil},
		{"Int", int(9), 9, false, nil},
		{"Float32", float32(3.14), 0, true, nil},
		{"String", "not an integer", 0, true, nil},
		{"Nil", nil, 0, true, nil},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result, err := tableValue.Int()

			test.checkError(t, err)

			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestTableValueToFloat64(t *testing.T) {
	tests := []test[float64]{
		{"Float32", float32(3.14), 3.1400000, false, nil},
		{"Float64", float64(2.71828), 2.71828, false, nil},
		{"Int", int(42), 0, true, nil},
		{"String", "not a float", 0, true, nil},
		{"Nil", nil, 0, true, nil},
	}

	const tolerance = 1e-6 // adjust this as needed

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result, err := tableValue.ToFloat64()

			test.checkError(t, err)

			if math.Abs(result-test.expected) > tolerance {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

func TestTableValueFloat(t *testing.T) {
	tests := []test[float64]{
		{"Float32", float32(3.14), 3.1400000, false, nil},
		{"Float64", float64(2.71828), 2.71828, false, nil},
		{"Int", int(42), 0, true, nil},
		{"String", "not a float", 0, true, nil},
		{"Nil", nil, 0, true, nil},
	}

	const tolerance = 1e-6 // adjust this as needed

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result, err := tableValue.Float()

			test.checkError(t, err)

			if math.Abs(result-test.expected) > tolerance {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

func TestTableValueToString(t *testing.T) {
	tests := []test[string]{
		{"String", "hello", "hello", false, nil},
		{"String", "world", "world", false, nil},
		{"Int", 42, "", true, nil},
		{"Nil", nil, "", true, nil},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result, err := tableValue.ToString()

			test.checkError(t, err)

			if result != test.expected {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

func TestTableValueString(t *testing.T) {
	tests := []test[string]{
		{"String", "hello", "hello", false, nil},
		{"String", "world", "world", false, nil},
		{"Int", 42, "", true, nil},
		{"Nil", nil, "", true, nil},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result, err := tableValue.String()

			test.checkError(t, err)

			if result != test.expected {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

func TestTableValueToBool(t *testing.T) {
	tests := []test[bool]{
		{"Bool", true, true, false, nil},
		{"Int", 0, false, true, nil},
		{"String", "true", false, true, nil},
		{"Nil", nil, false, true, nil},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result, err := tableValue.ToBool()

			test.checkError(t, err)

			if result != test.expected {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

func TestTableValueBool(t *testing.T) {
	tests := []test[bool]{
		{"Bool", true, true, false, nil},
		{"Int", 0, false, true, nil},
		{"String", "true", false, true, nil},
		{"Nil", nil, false, true, nil},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result, err := tableValue.Bool()

			test.checkError(t, err)

			if result != test.expected {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

func TestTableValueGetType(t *testing.T) {
	tests := []test[reflect.Type]{
		{"", int(42), reflect.TypeOf(int(0)), false, nil},
		{"", float64(3.14), reflect.TypeOf(float64(0)), false, nil},
		{"", "hello", reflect.TypeOf(""), false, nil},
		{"", true, reflect.TypeOf(true), false, nil},
		{"", nil, nil, false, nil},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result := tableValue.GetType()

			if result != test.expected {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

func TestTableValueType(t *testing.T) {
	tests := []struct {
		title    string
		value    interface{}
		expected reflect.Type
	}{
		{"", int(42), reflect.TypeOf(int(0))},
		{"", float64(3.14), reflect.TypeOf(float64(0))},
		{"", "hello", reflect.TypeOf("")},
		{"", true, reflect.TypeOf(true)},
		{"", nil, nil}, // Expected type for nil value
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result := tableValue.Type()

			if result != test.expected {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}
