//go:build preview.query

package ast

import (
	"testing"
)

func TestUnaryNode_Accept(t *testing.T) {
	node := NewUnaryNode(OperatorIsEmpty, NewLiteralNode("f"))
	visitor := NewStringerVisitor()
	node.Accept(visitor)
	if visitor.String() != "fISEMPTY" {
		t.Errorf("got %s, expected fISEMPTY", visitor.String())
	}
}

func TestNewUnaryNode(t *testing.T) {
	tests := []struct {
		name     string
		op       Operator
		left     Node
		expected string
	}{
		{"IsEmpty", OperatorIsEmpty, NewLiteralNode("f"), "fISEMPTY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := NewUnaryNode(tt.op, tt.left)
			if node.Op != tt.op || node.Left != tt.left {
				t.Errorf("NewUnaryNode failed")
			}
		})
	}
}
