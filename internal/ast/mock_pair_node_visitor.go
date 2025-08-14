package ast

import "github.com/stretchr/testify/mock"

type mockPairNodeVisitor struct {
	mock.Mock
}

func newMockPairNodeVisitor() *mockPairNodeVisitor {
	return &mockPairNodeVisitor{
		mock.Mock{},
	}
}

func (visitor *mockPairNodeVisitor) VisitPairNode(node *PairNode) {
	_ = visitor.Called(node)
}
