package query

import ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func IsEmpty() func(string) ast.Node {
	return valueWrapper2(ast.Operator("ISEMPTY"), nil)
}
