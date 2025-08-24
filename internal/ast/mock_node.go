//go:build preview.query

package ast

import "github.com/stretchr/testify/mock"

type mockNode struct {
	mock.Mock
}

func NewMockNode() *mockNode {
	return &mockNode{
		mock.Mock{},
	}
}

func (node *mockNode) Left() int {
	args := node.Called()
	return args.Int(0)
}

func (node *mockNode) Right() int {
	args := node.Called()
	return args.Int(0)
}

func (node *mockNode) Pos() int {
	args := node.Called()
	return args.Int(0)
}

func (node *mockNode) Accept(visitor NodeVisitor) {
	_ = node.Called(visitor)
}
