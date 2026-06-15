package tableapi

import (
	"testing"
)

func TestConvertType(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int
		err      bool
	}{
		{"Ok", 42, 42, false},
		{"Bad", "s", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := convertType[int](tt.input)
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if res != tt.expected {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
		})
	}
}

func TestIsNil(t *testing.T) {
	var s *string
	tests := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{"Nil", nil, true},
		{"NilPtr", s, true},
		{"NotNil", "v", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if isNil(tt.input) != tt.expected {
				t.Errorf("failed for %v", tt.input)
			}
		})
	}
}
