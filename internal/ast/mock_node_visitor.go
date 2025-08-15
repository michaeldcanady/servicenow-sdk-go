//go:build preview

package ast

import "github.com/stretchr/testify/mock"

type mockNodeVisitor struct {
	mock.Mock
}

func newMockNodeVisitor() *mockNodeVisitor {
	return &mockNodeVisitor{
		mock.Mock{},
	}
}

func (visitor *mockNodeVisitor) Visit(node Node) {
	_ = visitor.Called(node)
}

func (visitor *mockNodeVisitor) VisitArrayNode(node *ArrayNode) {
	_ = visitor.Called(node)
}

func (visitor *mockNodeVisitor) VisitBinaryNode(node *BinaryNode) {
	_ = visitor.Called(node)
}

func (visitor *mockNodeVisitor) VisitLiteralNode(node *LiteralNode) {
	_ = visitor.Called(node)
}

func (visitor *mockNodeVisitor) VisitPairNode(node *PairNode) {
	_ = visitor.Called(node)
}

func (visitor *mockNodeVisitor) VisitUnaryNode(node *UnaryNode) {
	_ = visitor.Called(node)
}
