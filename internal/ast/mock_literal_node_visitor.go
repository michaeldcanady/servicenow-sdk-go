package ast

import "github.com/stretchr/testify/mock"

type mockLiteralNodeVisitor struct {
	mock.Mock
}

func newMockLiteralNodeVisitor() *mockLiteralNodeVisitor {
	return &mockLiteralNodeVisitor{
		mock.Mock{},
	}
}

func (visitor *mockLiteralNodeVisitor) VisitLiteralNode(node *LiteralNode) {
	_ = visitor.Called(node)
}
