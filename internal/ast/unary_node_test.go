//go:build preview.query

package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnaryNode_Left(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "No Node",
			test: func(t *testing.T) {
				node := &UnaryNode{}

				assert.Equal(t, -1, node.Left())
			},
		},
		{
			name: "Has Node",
			test: func(t *testing.T) {
				node := newMockNode()
				node.On("Left").Return(1)

				unaryNode := &UnaryNode{
					Node: node,
				}

				assert.Equal(t, 1, unaryNode.Left())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestUnaryNode_Right(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &UnaryNode{
					Position: 1,
				}

				assert.Equal(t, 1, node.Right())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestUnaryNode_Pos(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &UnaryNode{
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

func TestUnaryNode_Accept(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &UnaryNode{}
				visitor := newMockNodeVisitor()
				visitor.On("VisitUnaryNode", node)

				node.Accept(visitor)
				visitor.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
