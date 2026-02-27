package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestDereference(t *testing.T) {
	s := "test"
	ps := &s
	pps := &ps
	var nilS *string

	tests := []struct {
		name     string
		input    any
		expected any
	}{
		{"String", s, "test"},
		{"Ptr", ps, "test"},
		{"PtrPtr", pps, "test"},
		{"NilPtr", nilS, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Dereference(reflect.ValueOf(tt.input))
			if res.Interface() != tt.expected {
				t.Errorf("got %v, expected %v", res.Interface(), tt.expected)
			}
		})
	}
}

func TestIsNumericKind(t *testing.T) {
	tests := []struct {
		name     string
		input    reflect.Kind
		expected bool
	}{
		{"Int8", reflect.Int8, true},
		{"Float64", reflect.Float64, true},
		{"String", reflect.String, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if isNumericKind(tt.input) != tt.expected {
				t.Errorf("%s failed", tt.name)
			}
		})
	}
}

func TestConvertNumeric(t *testing.T) {
	tests := []struct {
		name       string
		srcVal     any
		targetType reflect.Type
		expected   any
		err        bool
	}{
		{"Int8ToInt8", int8(1), reflect.TypeOf(int8(0)), int8(1), false},
		{"NonNumeric", "string", reflect.TypeOf(int8(0)), nil, true},
		{"UnsupportedTarget", int8(1), reflect.TypeOf(complex64(0)), nil, true},
		{"Overflow", int16(1000), reflect.TypeOf(int8(0)), nil, true},
		{"DecimalToInteger", float64(1.5), reflect.TypeOf(int(0)), nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := convertNumeric(reflect.ValueOf(tt.srcVal), tt.targetType)
			if tt.err {
				if err == nil {
					t.Error("Expected error")
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if res != tt.expected {
					t.Errorf("got %v, expected %v", res, tt.expected)
				}
			}
		})
	}
}

func TestConvertValue(t *testing.T) {
	tests := []struct {
		name       string
		srcVal     any
		targetType reflect.Type
		expected   any
		err        bool
	}{
		{"Assignable", "s", reflect.TypeOf(""), "s", false},
		{"NumericToNumeric", int8(1), reflect.TypeOf(int16(0)), int16(1), false},
		{"StringToNumeric", "123", reflect.TypeOf(int(0)), int(123), false},
		{"InvalidStringToNumeric", "bad", reflect.TypeOf(int(0)), nil, true},
		{"NumericToString", 123, reflect.TypeOf(""), "123", false},
		{"Unsupported", true, reflect.TypeOf(""), nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := convertValue(reflect.ValueOf(tt.srcVal), tt.targetType)
			if tt.err {
				if err == nil {
					t.Error("Expected error")
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if res.Interface() != tt.expected {
					t.Errorf("got %v, expected %v", res.Interface(), tt.expected)
				}
			}
		})
	}
}

func TestConvert(t *testing.T) {
	var s string
	var ps *string
	var i int

	tests := []struct {
		name   string
		input  any
		output any
		err    bool
	}{
		{"OutputNil", "v", nil, true},
		{"OutputNonPtr", "v", s, true},
		{"InputNil", nil, &s, false},
		{"DirectAssign", "v", &s, false},
		{"PtrTarget", 123, &ps, false},
		{"NonPtrTarget", "123", &i, false},
		{"ConvertError", "bad", &i, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Convert(tt.input, tt.output)
			if tt.err {
				if err == nil {
					t.Error("Expected error")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestStringToTime(t *testing.T) {
	format := time.RFC3339
	val := "2024-01-01T12:00:00Z"
	mutator := StringToTime(format)

	tests := []struct {
		name  string
		input string
		err   bool
	}{
		{"Valid", val, false},
		{"Invalid", "bad", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := mutator(tt.input)
			if tt.err {
				if err == nil {
					t.Error("Expected error")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}
