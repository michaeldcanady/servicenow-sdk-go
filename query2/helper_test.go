//go:build preview.query

package query2

import (
	"testing"
)

func TestConvertSliceToArrayNode(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{"Single", []string{"a"}, "a"},
		{"Multiple", []string{"a", "b"}, "a,b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := convertSliceToArrayNode(tt.input...)
			if node == nil {
				t.Fatal("Expected node, got nil")
			}
			// We can't easily check internal state without exporting or using a visitor, 
			// but we can check if it's correct via a Condition.
			c := String("f").IsOneOf(tt.input...)
			expectedFull := "fIN" + tt.expected
			if c.String() != expectedFull {
				t.Errorf("got %s, expected %s", c.String(), expectedFull)
			}
		})
	}
}
