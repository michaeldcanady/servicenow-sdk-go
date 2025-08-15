package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayNode_Left(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Empty array",
			test: func(t *testing.T) {
				node := &ArrayNode{
					Elements: make([]Node, 0),
				}

				assert.Equal(t, -1, node.Left())
			},
		},
		{
			name: "Nil array",
			test: func(t *testing.T) {
				node := &ArrayNode{
					Elements: nil,
				}

				assert.Equal(t, -1, node.Left())
			},
		},
		{
			name: "One element array",
			test: func(t *testing.T) {
				mockNode := newMockNode()
				mockNode.On("Left").Return(1)

				node := &ArrayNode{
					Elements: []Node{mockNode},
				}

				assert.Equal(t, 1, node.Left())
				mockNode.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestArrayNode_Right(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Empty array",
			test: func(t *testing.T) {
				node := &ArrayNode{
					Elements: make([]Node, 0),
				}

				assert.Equal(t, -1, node.Right())
			},
		},
		{
			name: "Nil array",
			test: func(t *testing.T) {
				node := &ArrayNode{
					Elements: nil,
				}

				assert.Equal(t, -1, node.Right())
			},
		},
		{
			name: "One element array",
			test: func(t *testing.T) {
				mockNode := newMockNode()
				mockNode.On("Right").Return(1)

				node := &ArrayNode{
					Elements: []Node{mockNode},
				}

				assert.Equal(t, 1, node.Right())
				mockNode.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestArrayNode_Pos(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestArrayNode_Accept(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &ArrayNode{}

				visitor := newMockNodeVisitor()
				visitor.On("VisitArrayNode", node)

				visitor.VisitArrayNode(node)

				visitor.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
