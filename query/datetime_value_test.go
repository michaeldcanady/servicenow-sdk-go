package query

import (
	"testing"
	"time"
)

func TestDateTimeValue_String(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Standard", "v", "v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := DateTimeValue{literal: tt.input}
			if v.String() != tt.expected {
				t.Errorf("got %s, expected %s", v.String(), tt.expected)
			}
		})
	}
}

func TestNewDateTimeValue(t *testing.T) {
	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{"Standard", now, "2024-01-01 00:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewDateTimeValue(tt.input)
			if v.String() != tt.expected {
				t.Errorf("got %s, expected %s", v.String(), tt.expected)
			}
		})
	}
}

func TestJS(t *testing.T) {
	tests := []struct {
		name     string
		expr     string
		expected string
		wantErr  bool
	}{
		{"Standard", "gs.daysAgoStart(0)", "javascript:gs.daysAgoStart(0)", false},
		// Commas are literals inside a javascript: expression (multi-argument
		// calls); only "^" and "@" structure the surrounding term.
		{"MultiArgCall", "gs.dateGenerate('2024-01-01','00:00:00')", "javascript:gs.dateGenerate('2024-01-01','00:00:00')", false},
		{"CaretInjection", "gs.daysAgoStart(0)^ORactive=false", "", true},
		{"AtInjection", "a@b", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := JS(tt.expr)
			if tt.wantErr {
				if v.err == nil {
					t.Fatal("Expected validation error")
				}
				return
			}
			if v.err != nil {
				t.Fatalf("Unexpected error: %v", v.err)
			}
			if v.String() != tt.expected {
				t.Errorf("got %s, expected %s", v.String(), tt.expected)
			}
		})
	}
}
