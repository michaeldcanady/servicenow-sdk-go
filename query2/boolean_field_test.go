//go:build preview.query

package query2

import (
	"testing"
)

func TestBooleanField_Is(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      bool
		expected string
	}{
		{"True", "f", true, "f=true"},
		{"False", "f", false, "f=false"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Boolean(tt.field).Is(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Boolean(tt.field).Is(tt.val).String(), tt.expected)
			}
		})
	}
}
