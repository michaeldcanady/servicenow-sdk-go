package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func IsNotOneOf[T Numeric](values ...T) func(string) ast.Node {
	return ArrayCondition(ast.OperatorIsNotOneOf, values)
}
