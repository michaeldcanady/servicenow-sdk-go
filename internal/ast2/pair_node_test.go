//go:build preview.query

package ast2

import (
	"testing"
)

func TestPairNode_Accept(t *testing.T) {
	node := NewPairNode(NewLiteralNode("a"), NewLiteralNode("b"))
	visitor := NewStringerVisitor()
	node.Accept(visitor)
	if visitor.String() != "a@b" {
		t.Errorf("got %s, expected a@b", visitor.String())
	}
}

func TestNewPairNode(t *testing.T) {
	tests := []struct {
		name  string
		left  Node
		right Node
	}{
		{"Basic", NewLiteralNode("a"), NewLiteralNode("b")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := NewPairNode(tt.left, tt.right)
			if node.Left != tt.left || node.Right != tt.right {
				t.Error("NewPairNode failed")
			}
		})
	}
}
