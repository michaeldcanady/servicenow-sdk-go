//go:build preview.query

package query

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
	"github.com/stretchr/testify/mock"
)

type mockConditionBuilder struct {
	mock.Mock
}

func newMockConditionBuilder() *mockConditionBuilder {
	return &mockConditionBuilder{
		mock.Mock{},
	}
}

func (mock *mockConditionBuilder) addErrors(errs ...error) {
	_ = mock.Called(errs)
}
func (mock *mockConditionBuilder) binaryCondition(operator ast.Operator, value ast.Node) *T {
	mock.Called(operator, value)
}
func (mock *mockConditionBuilder) unaryCondition(operator ast.Operator) *T {
	mock.Called(operator)
}
