package query

import ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func IsNot[T Primitive](val T) func(string) ast.Node {
	return valueWrapper2(ast.Operator("!="), val)
}
