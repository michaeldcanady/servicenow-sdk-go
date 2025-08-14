package ast

import "github.com/stretchr/testify/mock"

type mockUnaryNodeVisitor struct {
	mock.Mock
}

func newMockUnaryNodeVisitor() *mockUnaryNodeVisitor {
	return &mockUnaryNodeVisitor{
		mock.Mock{},
	}
}

func (visitor *mockUnaryNodeVisitor) VisitUnaryNode(node *UnaryNode) {
	_ = visitor.Called(node)
}
