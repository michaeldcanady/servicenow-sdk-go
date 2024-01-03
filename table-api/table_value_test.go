package tableapi

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const tolerance = 1e-6 // adjust this as needed

var (
	int64TestSet = []test[int64]{
		{"Int64", int64(42), 42, false, nil},
		{"Int32", int32(123), 123, false, nil},
		{"Float32", float32(3.14), 0, true, nil},
		{"String", "not an integer", 0, true, nil},
		{"Nil", nil, 0, true, nil},
	}
	float64TestSet = []test[float64]{
		{"Float32", float32(3.14), 3.1400000, false, nil},
		{"Float64", float64(2.71828), 2.71828, false, nil},
		{"Int", int(42), 0, true, nil},
		{"String", "not a float", 0, true, nil},
		{"Nil", nil, 0, true, nil},
	}
	stringTestSet = []test[string]{
		{"String", "hello", "hello", false, nil},
		{"String", "world", "world", false, nil},
		{"Int", 42, "", true, nil},
		{"Nil", nil, "", true, nil},
	}
	boolTestSet = []test[bool]{
		{"Bool", true, true, false, nil},
		{"Int", 0, false, true, nil},
		{"String", "true", false, true, nil},
		{"Nil", nil, false, true, nil},
	}
	getTypeTestSet = []test[reflect.Type]{
		{"", int(42), reflect.TypeOf(int(0)), false, nil},
		{"", float64(3.14), reflect.TypeOf(float64(0)), false, nil},
		{"", "hello", reflect.TypeOf(""), false, nil},
		{"", true, reflect.TypeOf(true), false, nil},
		{"", nil, nil, false, nil},
	}
)

type test[T any] struct {
	title     string
	value     interface{}
	expected  T
	expectErr bool
	err       error
}

func (te test[T]) checkError(t *testing.T, err error) {
	if te.expectErr && err == nil {
		if te.err != nil {
			assert.ErrorIs(t, err, te.err)
		} else {
			assert.Error(t, err)
		}
	}

	if !te.expectErr && err != nil {
		assert.Nil(t, err)
	}
}

func TestTableValueToInt64(t *testing.T) {
	for _, test := range int64TestSet {
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
	for _, test := range int64TestSet {
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
	for _, test := range float64TestSet {
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
	for _, test := range float64TestSet {
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
	for _, test := range stringTestSet {
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
	for _, test := range stringTestSet {
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
	for _, test := range boolTestSet {
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
	for _, test := range boolTestSet {
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
	for _, test := range getTypeTestSet {
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
	for _, test := range getTypeTestSet {
		t.Run(test.title, func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result := tableValue.Type()

			if result != test.expected {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}
