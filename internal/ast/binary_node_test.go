//go:build preview.query

package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBinaryNode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				operator := OperatorBefore
				node1 := NewLiteralNode("value1")
				node2 := NewLiteralNode("value")
				expr := NewBinaryNode(node1, operator, node2)

				assert.Equal(t, operator, expr.Op)
				assert.Equal(t, node1, expr.LeftExpression)
				assert.Equal(t, node2, expr.RightExpression)
				assert.Equal(t, -1, expr.Position)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBinaryNode_Left(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Has left expression",
			test: func(t *testing.T) {
				leftNode := NewMockNode()
				leftNode.On("Pos").Return(1)

				node := &BinaryNode{
					LeftExpression: leftNode,
				}

				assert.Equal(t, 1, node.Left())
				leftNode.AssertExpectations(t)
			},
		},
		{
			name: "No left expression",
			test: func(t *testing.T) {
				node := &BinaryNode{
					LeftExpression: nil,
				}

				assert.Equal(t, -1, node.Left())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBinaryNode_Right(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Has left expression",
			test: func(t *testing.T) {
				rightNode := NewMockNode()
				rightNode.On("Pos").Return(1)

				node := &BinaryNode{
					RightExpression: rightNode,
				}

				assert.Equal(t, 1, node.Right())
				rightNode.AssertExpectations(t)
			},
		},
		{
			name: "No right expression",
			test: func(t *testing.T) {
				node := &BinaryNode{
					RightExpression: nil,
				}

				assert.Equal(t, -1, node.Right())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBinaryNode_Pos(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &BinaryNode{
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

func TestBinaryNode_Accept(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &BinaryNode{}
				visitor := newMockNodeVisitor()
				visitor.On("VisitBinaryNode", node)

				node.Accept(visitor)
				visitor.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBinaryNode_Operator(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &BinaryNode{
					Op: OperatorAnd,
				}

				assert.Equal(t, OperatorAnd, node.Operator())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
