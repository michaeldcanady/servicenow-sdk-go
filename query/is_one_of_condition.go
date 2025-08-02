package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func IsOneOf[T Primitive](values ...T) func(string) ast.Node {
	return ArrayCondition(ast.OperatorIsOneOf, values)
}
