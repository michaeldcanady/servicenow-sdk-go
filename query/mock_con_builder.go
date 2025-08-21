//go:build preview.query

package query

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

type mockConBuilder[T ast.Primitive | time.Time, Q QueryBuilder] struct {
	*mockConditionBuilder[Q]
}

func newMockConBuilder[T ast.Primitive | time.Time, Q QueryBuilder]() *mockConBuilder[T, Q] {
	return &mockConBuilder[T, Q]{
		newMockConditionBuilder[Q](),
	}
}

func (mock *mockConBuilder[T, Q]) Is(value T) *Q {
	args := mock.Called(value)

	return args.Get(0).(*Q)
}

func (mock *mockConBuilder[T, Q]) IsAnything() *Q {
	args := mock.Called()

	return args.Get(0).(*Q)
}

func (mock *mockConBuilder[T, Q]) IsDynamic(sysID string) *Q {
	args := mock.Called(sysID)

	return args.Get(0).(*Q)
}

func (mock *mockConBuilder[T, Q]) IsEmpty() *Q {
	args := mock.Called()

	return args.Get(0).(*Q)
}

func (mock *mockConBuilder[T, Q]) IsNot(value T) *Q {
	args := mock.Called(value)

	return args.Get(0).(*Q)
}

func (mock *mockConBuilder[T, Q]) IsNotEmpty() *Q {
	args := mock.Called()

	return args.Get(0).(*Q)
}

func (mock *mockConBuilder[T, Q]) IsSame(sysID string) *Q {
	args := mock.Called(sysID)

	return args.Get(0).(*Q)
}
