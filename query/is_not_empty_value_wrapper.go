package query

import ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func IsNotEmpty() func(string) ast.Node {
	return valueWrapper2(ast.Operator("ISNOTEMPTY"), nil)
}
