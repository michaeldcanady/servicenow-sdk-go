//go:build preview.query

package ast

import (
	"testing"
)

func TestLiteralNode_Accept(t *testing.T) {
	node := NewLiteralNode("v")
	visitor := NewStringerVisitor()
	node.Accept(visitor)
	if visitor.String() != "v" {
		t.Errorf("got %s, expected v", visitor.String())
	}
}

func TestNewLiteralNode(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{"String", "hello", "hello"},
		{"Int", 123, "123"},
		{"Bool", true, "true"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := NewLiteralNode(tt.input)
			if node.Value != tt.expected {
				t.Errorf("got %s, expected %s", node.Value, tt.expected)
			}
		})
	}
}
