package ast

import "github.com/stretchr/testify/mock"

type mockBinaryNodeVisitor struct {
	mock.Mock
}

func newMockBinaryNodeVisitor() *mockBinaryNodeVisitor {
	return &mockBinaryNodeVisitor{
		mock.Mock{},
	}
}

func (visitor *mockBinaryNodeVisitor) VisitBinaryNode(node *BinaryNode) {
	_ = visitor.Called(node)
}
