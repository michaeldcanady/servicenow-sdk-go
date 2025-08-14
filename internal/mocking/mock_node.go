package mocking

import "github.com/stretchr/testify/mock"

type MockNode struct {
	mock.Mock
}

func NewMockNode() *MockNode {
	return &MockNode{
		mock.Mock{},
	}
}

func (node *MockNode) Left() int {
	args := node.Called()
	return args.Int(0)
}

func (node *MockNode) Right() int {
	args := node.Called()
	return args.Int(0)
}

func (node *MockNode) Pos() int {
	args := node.Called()
	return args.Int(0)
}
