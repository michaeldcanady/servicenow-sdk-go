package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiteralNode_Left(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &LiteralNode{
					Position: 1,
				}

				assert.Equal(t, 1, node.Left())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestLiteralNode_Right(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := "I am a value"
				node := &LiteralNode{
					Position: 1,
					Value:    value,
				}

				assert.Equal(t, 1+len(value), node.Right())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestLiteralNode_Pos(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &LiteralNode{
					Position: 1,
				}

				assert.Equal(t, 1, node.Pos())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestLiteralNode_Accept(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &LiteralNode{}
				visitor := newMockLiteralNodeVisitor()
				visitor.On("VisitLiteralNode", node)

				node.Accept(visitor)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
