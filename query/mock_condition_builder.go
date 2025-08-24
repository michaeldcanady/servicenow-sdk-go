//go:build preview.query

package query

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
	"github.com/stretchr/testify/mock"
)

type mockConditionBuilder[T QueryBuilder] struct {
	mock.Mock
}

func newMockConditionBuilder[T QueryBuilder]() *mockConditionBuilder[T] {
	return &mockConditionBuilder[T]{
		mock.Mock{},
	}
}

func (mock *mockConditionBuilder[T]) addErrors(errs ...error) {
	_ = mock.Called(errs)
}

func (mock *mockConditionBuilder[T]) binaryCondition(operator ast.Operator, value ast.Node) *T {
	args := mock.Called(operator, value)
	return args.Get(0).(*T)
}

func (mock *mockConditionBuilder[T]) unaryCondition(operator ast.Operator) *T {
	args := mock.Called(operator)
	return args.Get(0).(*T)
}
