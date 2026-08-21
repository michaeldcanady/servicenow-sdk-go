package ast

import (
	"testing"
)

func TestOperator_String(t *testing.T) {
	tests := []struct {
		name     string
		op       Operator
		expected string
	}{
		{"Is", OperatorIs, "="},
		{"IsNot", OperatorIsNot, "!="},
		{"Unknown", Operator(-99), "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.op.String() != tt.expected {
				t.Errorf("got %s, expected %s", tt.op.String(), tt.expected)
			}
		})
	}
}

func TestOperator_known(t *testing.T) {
	tests := []struct {
		name     string
		op       Operator
		expected bool
	}{
		{"Known", OperatorIs, true},
		{"UnknownSentinel", OperatorUnknown, false},
		{"OutOfRange", Operator(999), false},
		{"Negative", Operator(-99), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.op.known() != tt.expected {
				t.Errorf("got %t, expected %t", tt.op.known(), tt.expected)
			}
		})
	}
}
