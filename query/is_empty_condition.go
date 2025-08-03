package query

import ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func IsEmptyCondition() func(string) ast.Node {
	return unaryCondition(ast.OperatorIsEmpty)
}
