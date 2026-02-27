package utils

import (
	"testing"
)

func TestIsNil(t *testing.T) {
	var s *string
	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{"UntypedNil", nil, true},
		{"NilPtr", s, true},
		{"NotNil", "v", false},
		{"IntNotNil", 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if IsNil(tt.input) != tt.expected {
				t.Errorf("got %v, expected %v", IsNil(tt.input), tt.expected)
			}
		})
	}
}

func TestToPointer(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"String", "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := ToPointer(tt.input)
			if *res != tt.input {
				t.Errorf("got %v, expected %v", *res, tt.input)
			}
		})
	}
}

func TestIsPointer(t *testing.T) {
	s := "v"
	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{"Pointer", &s, true},
		{"NotPointer", s, false},
		{"Nil", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if IsPointer(tt.input) != tt.expected {
				t.Errorf("got %v, expected %v", IsPointer(tt.input), tt.expected)
			}
		})
	}
}
