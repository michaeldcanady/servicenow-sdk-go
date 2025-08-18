//go:build preview.query

package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

type UnaryConditionBuilder[T QueryBuilder] interface {
	unaryCondition(operator ast.Operator) *T
}
