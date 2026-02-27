//go:build preview.query

package ast

import (
	"testing"
)

func TestBinaryNode_Accept(t *testing.T) {
	node := NewBinaryNode(NewLiteralNode("f"), OperatorIs, NewLiteralNode("v"))
	visitor := NewStringerVisitor()
	node.Accept(visitor)
	if visitor.String() != "f=v" {
		t.Errorf("got %s, expected f=v", visitor.String())
	}
}

func TestNewBinaryNode(t *testing.T) {
	tests := []struct {
		name  string
		left  Node
		op    Operator
		right Node
	}{
		{"Standard", NewLiteralNode("f"), OperatorIs, NewLiteralNode("v")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := NewBinaryNode(tt.left, tt.op, tt.right)
			if node.Left != tt.left || node.Op != tt.op || node.Right != tt.right {
				t.Error("NewBinaryNode failed")
			}
		})
	}
}
