// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package query

import (
	"errors"
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/ast"
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
		// A rejected condition must render the fail-closed placeholder instead
		// of panicking on its missing subtree.
		{"ErrorCondition", Number("n").Between(10, 5), invalidQueryPlaceholder},
		{"ErrorOnLeft", Number("n").Between(10, 5).And(String("b").Is("2")), "<invalid query>^b=2"},
		{"ErrorOnRight", String("a").Is("1").Or(Number("n").Between(10, 5)), "a=1^OR<invalid query>"},
		{"NilNode", NewCondition(nil), invalidQueryPlaceholder},
		// A hand-built malformed tree must degrade to the placeholder instead
		// of panicking or emitting a truncated query.
		{"MalformedTree", NewCondition(ast.NewBinaryNode(nil, ast.OperatorIs, ast.NewLiteralNode("v"))), invalidQueryPlaceholder},
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

// stubCondition lets tests exercise the typed-nil trap: (*stubCondition)(nil)
// is a non-nil interface wrapping a nil pointer, which a plain == nil check
// would miss.
type stubCondition struct {
	baseCondition
}

func TestCombinatorNilSafety(t *testing.T) {
	tests := []struct {
		name          string
		c             Condition
		expectedCause error
	}{
		// Degenerate combination inputs must fail closed to an error
		// condition rather than panicking on the nil interface.
		{"MethodAnd", String("f").Is("v").And(nil), ErrNilCondition},
		{"MethodOr", String("f").Is("v").Or(nil), ErrNilCondition},
		{"VariadicFirst", And(nil), ErrNilCondition},
		{"VariadicOrFirst", Or(nil), ErrNilCondition},
		{"VariadicSecond", And(String("f").Is("v"), nil), ErrNilCondition},
		{"VariadicOrSecond", Or(String("f").Is("v"), nil), ErrNilCondition},
		{"TypedNilVariadic", And((*stubCondition)(nil)), ErrNilCondition},
		{"TypedNilMethod", String("f").Is("v").And((*stubCondition)(nil)), ErrNilCondition},
		{"EmptyAnd", And(), ErrNoConditions},
		{"EmptyOr", Or(), ErrNoConditions},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.c == nil {
				t.Fatal("combinator returned nil Condition")
			}
			if err := tt.c.Error(); !errors.Is(err, tt.expectedCause) {
				t.Errorf("got %v, expected cause %v", err, tt.expectedCause)
			}
			if rendered := tt.c.String(); rendered != invalidQueryPlaceholder {
				t.Errorf("got %q, expected %q", rendered, invalidQueryPlaceholder)
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
			if !errors.Is(c.Error(), tt.err) {
				t.Error("Failed to store error")
			}
		})
	}
}
