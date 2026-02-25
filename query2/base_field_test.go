//go:build preview.query

package query2

import (
	"testing"
)

func TestBaseField_IsAnything(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "fANYTHING"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsAnything().String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsAnything().String(), tt.expected)
			}
		})
	}
}

func TestBaseField_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "fISEMPTY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsEmpty().String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsEmpty().String(), tt.expected)
			}
		})
	}
}

func TestBaseField_IsNotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "fISNOTEMPTY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsNotEmpty().String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsNotEmpty().String(), tt.expected)
			}
		})
	}
}

func TestBaseField_IsDynamic(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		sid      string
		expected string
	}{
		{"Standard", "f", "sid", "fDYNAMICsid"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsDynamic(tt.sid).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsDynamic(tt.sid).String(), tt.expected)
			}
		})
	}
}

func TestBaseField_IsSame(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		other    string
		expected string
	}{
		{"Standard", "f", "o", "fSAMEASo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsSame(tt.other).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsSame(tt.other).String(), tt.expected)
			}
		})
	}
}

func TestBaseField_IsDifferent(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		other    string
		expected string
	}{
		{"Standard", "f", "o", "fNSAMEASo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsDifferent(tt.other).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsDifferent(tt.other).String(), tt.expected)
			}
		})
	}
}

func TestBaseField_IsInHierarchy(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "fIN HIERARCHY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsInHierarchy().String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsInHierarchy().String(), tt.expected)
			}
		})
	}
}
