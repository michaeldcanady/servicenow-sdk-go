package ast

import "github.com/stretchr/testify/mock"

type mockArrayNodeVisitor struct {
	mock.Mock
}

func newMockArrayNodeVisitor() *mockArrayNodeVisitor {
	return &mockArrayNodeVisitor{
		mock.Mock{},
	}
}

func (visitor *mockArrayNodeVisitor) VisitArrayNode(node *ArrayNode) {
	_ = visitor.Called(node)
}
