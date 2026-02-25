//go:build preview.query

package query2

import (
	"testing"
	"time"
)

func TestCondition_And(t *testing.T) {
	tests := []struct {
		name     string
		c1, c2   Condition
		expected string
	}{
		{"Standard", String("a").Is("1"), String("b").Is("2"), "a=1^b=2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.c1.And(tt.c2).String() != tt.expected {
				t.Errorf("got %s, expected %s", tt.c1.And(tt.c2).String(), tt.expected)
			}
		})
	}
}

func TestCondition_Or(t *testing.T) {
	tests := []struct {
		name     string
		c1, c2   Condition
		expected string
	}{
		{"Standard", String("a").Is("1"), String("b").Is("2"), "a=1^ORb=2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.c1.Or(tt.c2).String() != tt.expected {
				t.Errorf("got %s, expected %s", tt.c1.Or(tt.c2).String(), tt.expected)
			}
		})
	}
}

func TestCondition_ToNode(t *testing.T) {
	tests := []struct {
		name string
		c    Condition
	}{
		{"Basic", String("f").Is("v")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.c.ToNode() == nil {
				t.Error("ToNode returned nil")
			}
		})
	}
}

func TestCondition_String(t *testing.T) {
	tests := []struct {
		name     string
		c        Condition
		expected string
	}{
		{"Standard", String("f").Is("v"), "f=v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.c.String() != tt.expected {
				t.Errorf("got %s, expected %s", tt.c.String(), tt.expected)
			}
		})
	}
}

func TestCondition_Error(t *testing.T) {
	tests := []struct {
		name string
		c    Condition
	}{
		{"InvalidRange", Number("f").Between(10, 5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.c.Error() == nil {
				t.Error("Expected error")
			}
		})
	}
}

func TestNewCondition(t *testing.T) {
	// Already tested but let's be explicit
	c := NewCondition(nil)
	if c == nil {
		t.Error("NewCondition should not return nil even for nil node")
	}
}

func TestNewErrorCondition(t *testing.T) {
	_, err := time.ParseDuration("invalid")
	tests := []struct {
		name string
		err  error
	}{
		{"Basic", err},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewErrorCondition(tt.err)
			if c.Error() != tt.err {
				t.Error("Failed to store error")
			}
		})
	}
}
