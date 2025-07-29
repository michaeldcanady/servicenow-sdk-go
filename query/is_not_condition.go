package query

import ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func IsNotCondition[T Primitive](val T) func(string) ast.Node {
	return Condition(ast.OperatorIsNot, val)
}
