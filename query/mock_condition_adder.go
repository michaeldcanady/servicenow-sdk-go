//go:build preview.query

package query

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

type mockConditionAdder[T logicalConditionBuilder] struct {
	*mockErrorAdder
}

func newMockConditionAdder[T logicalConditionBuilder]() *mockConditionAdder[T] {
	return &mockConditionAdder[T]{
		newMockErrorAdder(),
	}
}

func (mock *mockConditionAdder[T]) addCondition(condition ast.Node) T {
	arg := mock.Called(condition)

	return arg.Get(0).(T)
}
