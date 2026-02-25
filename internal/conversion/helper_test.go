package conversion

import (
	"reflect"
	"testing"
)

func TestHasDecimalPlace(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected bool
	}{
		{"NoDecimal", 1.0, false},
		{"HasDecimal", 1.5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if hasDecimalPlace(tt.input) != tt.expected {
				t.Errorf("got %v, expected %v", hasDecimalPlace(tt.input), tt.expected)
			}
		})
	}
}

func TestIsNumericType(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{"Int", 1, true},
		{"IntPtr", reflect.TypeOf(int(0)), true},
		{"String", "s", false},
		{"Nil", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if isNumericType(tt.input) != tt.expected {
				t.Errorf("failed for %v", tt.input)
			}
		})
	}
}

func TestIsCompatibleInt(t *testing.T) {
	tests := []struct {
		name     string
		in       any
		tp       reflect.Type
		expected bool
	}{
		{"Int8Ok", 10, reflect.TypeOf(int8(0)), true},
		{"Int8Overflow", 200, reflect.TypeOf(int8(0)), false},
		{"Int8Decimal", 10.5, reflect.TypeOf(int8(0)), false},
		{"NonNumeric", "s", reflect.TypeOf(int(0)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if isCompatibleInt(tt.in, tt.tp) != tt.expected {
				t.Errorf("%s failed", tt.name)
			}
		})
	}
}

func TestIsCompatible(t *testing.T) {
	tests := []struct {
		name     string
		val      any
		tp       reflect.Type
		strict   bool
		expected bool
	}{
		{"Numeric", 10, reflect.TypeOf(int16(0)), false, true},
		{"StrictMatch", "s", reflect.TypeOf(""), true, true},
		{"StrictMismatch", 10, reflect.TypeOf(int16(0)), true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if isCompatible(tt.val, tt.tp, tt.strict) != tt.expected {
				t.Errorf("%s failed: got %v, expected %v", tt.name, isCompatible(tt.val, tt.tp, tt.strict), tt.expected)
			}
		})
	}
}

func TestAs(t *testing.T) {
	var i int16
	tests := []struct {
		name  string
		in    any
		out   any
		err   bool
	}{
		{"Successful", int16(10), &i, false},
		{"NilInput", nil, &i, false},
		{"NilOutput", 10, nil, true},
		{"NonPtrOutput", 10, i, true},
		{"Incompatible", "s", &i, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := As(tt.in, tt.out)
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

func TestAs2(t *testing.T) {
	var i16 int16
	var s string
	ps := &s
	tests := []struct {
		name   string
		in     any
		out    any
		strict bool
		err    bool
	}{
		{"DirectMatch", "v", &s, true, false},
		{"DirectPtrMatch", ps, &ps, true, false},
		{"CompatibleNumeric", int16(10), &i16, false, false},
		{"IncompatibleStrict", 10, &i16, true, true},
		{"NilInput", nil, &s, true, false},
		{"NilOutput", "v", nil, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := As2(tt.in, tt.out, tt.strict)
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
