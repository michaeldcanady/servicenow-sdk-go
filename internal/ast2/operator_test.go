//go:build preview.query

package ast2

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
