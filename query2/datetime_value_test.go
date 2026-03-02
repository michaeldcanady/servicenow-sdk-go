//go:build preview.query

package query2

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
