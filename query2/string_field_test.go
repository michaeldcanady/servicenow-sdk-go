//go:build preview.query

package query2

import (
	"testing"
)

func TestStringField_Is(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "f=v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).Is(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).Is(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestStringField_IsNot(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "f!=v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsNot(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsNot(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestStringField_StartsWith(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fSTARTSWITHv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).StartsWith(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).StartsWith(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestStringField_EndsWith(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fENDSWITHv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).EndsWith(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).EndsWith(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestStringField_Contains(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fLIKEv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).Contains(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).Contains(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestStringField_DoesNotContain(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "v", "fNOT LIKEv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).DoesNotContain(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).DoesNotContain(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestStringField_IsOneOf(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		vals     []string
		expected string
	}{
		{"Multiple", "f", []string{"a", "b"}, "fINa,b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsOneOf(tt.vals...).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsOneOf(tt.vals...).String(), tt.expected)
			}
		})
	}
}

func TestStringField_IsNotOneOf(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		vals     []string
		expected string
	}{
		{"Multiple", "f", []string{"a", "b"}, "fNOT INa,b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsNotOneOf(tt.vals...).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsNotOneOf(tt.vals...).String(), tt.expected)
			}
		})
	}
}

func TestStringField_IsEmptyString(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"Basic", "f", "fEMPTYSTRING"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).IsEmptyString().String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).IsEmptyString().String(), tt.expected)
			}
		})
	}
}

func TestStringField_MatchesPattern(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      string
		expected string
	}{
		{"Standard", "f", "p", "fMATCHES PATTERNp"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).MatchesPattern(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).MatchesPattern(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestStringField_Between(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		l, u     string
		expected string
	}{
		{"Standard", "f", "a", "b", "fBETWEENa@b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if String(tt.field).Between(tt.l, tt.u).String() != tt.expected {
				t.Errorf("got %s, expected %s", String(tt.field).Between(tt.l, tt.u).String(), tt.expected)
			}
		})
	}
}
