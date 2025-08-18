//go:build preview.query

package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

type BinaryConditionBuilder[T QueryBuilder] interface {
	binaryCondition(operator ast.Operator, value ast.Node) *T
}
