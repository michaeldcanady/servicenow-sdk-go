package tableapi

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTableValueToInt64(t *testing.T) {
	tests := []struct {
		value      interface{}
		expected   int64
		expectErr  bool
		errorCheck func(error) bool
	}{
		{int64(42), 42, false, nil},
		{int64(123), 123, false, nil},
		{float32(3.14), 0, true, IsNotNilError},
		{"not an integer", 0, true, IsNotNilError},
		{nil, 0, true, IsNotNilError},
	}

	for _, test := range tests {
		t.Run("ToInt64", func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result, err := tableValue.ToInt64()

			if test.errorCheck != nil {
				if !test.errorCheck(err) {
					t.Errorf("Expected error, got nil")
				}
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}

func IsNotNilError(err error) bool {
	return err != nil
}

func TestTableValueToFloat64(t *testing.T) {
	tests := []struct {
		value      interface{}
		expected   float64
		expectErr  bool
		errorCheck func(error) bool
	}{
		{float32(3.14), 0, true, IsNotNilError},
		{float64(2.71828), 2.71828, false, nil},
		{int(42), 0, true, IsNotNilError},
		{"not a float", 0, true, IsNotNilError},
		{nil, 0, true, IsNotNilError},
	}

	for _, test := range tests {
		t.Run("ToFloat64", func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result, err := tableValue.ToFloat64()

			if test.errorCheck != nil {
				if !test.errorCheck(err) {
					t.Errorf("Expected error, got nil")
				}
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if result != test.expected {
				t.Errorf("Expected %f, got %f", test.expected, result)
			}
		})
	}
}

func TestTableValueToString(t *testing.T) {
	tests := []struct {
		value      interface{}
		expected   string
		expectErr  bool
		errorCheck func(error) bool
	}{
		{"hello", "hello", false, nil},
		{"world", "world", false, nil},
		{42, "", true, IsNotNilError},
		{nil, "", true, IsNotNilError},
	}

	for _, test := range tests {
		t.Run("ToString", func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result, err := tableValue.ToString()

			if test.errorCheck != nil {
				if !test.errorCheck(err) {
					t.Errorf("Expected error, got nil")
				}
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}

func TestTableValueGetType(t *testing.T) {
	tests := []struct {
		value    interface{}
		expected reflect.Type
	}{
		{int(42), reflect.TypeOf(int(0))},
		{float64(3.14), reflect.TypeOf(float64(0))},
		{"hello", reflect.TypeOf("")},
		{true, reflect.TypeOf(true)},
		{nil, nil}, // Expected type for nil value
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Value=%v", test.value), func(t *testing.T) {
			tableValue := &TableValue{value: test.value}
			result := tableValue.GetType()

			if result != test.expected {
				t.Errorf("Expected type %v, got %v", test.expected, result)
			}
		})
	}
}
