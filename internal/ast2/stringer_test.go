//go:build preview.query

package ast2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStringerVisitor(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Create visitor",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := NewStringerVisitor()
			assert.NotNil(t, v)
		})
	}
}

func TestStringerVisitor_Visit(t *testing.T) {
	tests := []struct {
		name     string
		visit    func(v Visitor)
		expected string
	}{
		{
			name: "VisitLiteral",
			visit: func(v Visitor) {
				n := NewLiteralNode("v")
				v.VisitLiteral(n)
			},
			expected: "v",
		},
		{
			name: "VisitUnary",
			visit: func(v Visitor) {
				n := NewUnaryNode(OperatorIsEmpty, NewLiteralNode("f"))
				v.VisitUnary(n)
			},
			expected: "fISEMPTY",
		},
		{
			name: "VisitBinary",
			visit: func(v Visitor) {
				n := NewBinaryNode(NewLiteralNode("f"), OperatorIs, NewLiteralNode("v"))
				v.VisitBinary(n)
			},
			expected: "f=v",
		},
		{
			name: "VisitPair",
			visit: func(v Visitor) {
				n := NewPairNode(NewLiteralNode("a"), NewLiteralNode("b"))
				v.VisitPair(n)
			},
			expected: "a@b",
		},
		{
			name: "VisitArray",
			visit: func(v Visitor) {
				n := NewArrayNode(NewLiteralNode("a"), NewLiteralNode("b"))
				v.VisitArray(n)
			},
			expected: "a,b",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := NewStringerVisitor()
			test.visit(v)
			assert.Equal(t, test.expected, v.String())
		})
	}
}
