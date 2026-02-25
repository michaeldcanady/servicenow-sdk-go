//go:build preview.query

package ast2

import (
	"testing"
)

func TestArrayNode_Accept(t *testing.T) {
	node := NewArrayNode(NewLiteralNode("a"), NewLiteralNode("b"))
	visitor := NewStringerVisitor()
	node.Accept(visitor)
	if visitor.String() != "a,b" {
		t.Errorf("got %s, expected a,b", visitor.String())
	}
}

func TestNewArrayNode(t *testing.T) {
	tests := []struct {
		name  string
		nodes []Node
	}{
		{"Multiple", []Node{NewLiteralNode("a"), NewLiteralNode("b")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := NewArrayNode(tt.nodes...)
			if len(node.Nodes) != len(tt.nodes) {
				t.Error("NewArrayNode failed")
			}
		})
	}
}
