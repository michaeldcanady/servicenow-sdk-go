//go:build preview.query

package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

type conditionAdder[T logicalConditionBuilder] interface {
	addCondition(condition ast.Node) T
	ErrorAdder
}
