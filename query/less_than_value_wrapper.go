package query

import ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func LessThan[T Numeric](val T) func(string) ast.Node {
	return valueWrapper2(ast.Operator("<"), val)
}
