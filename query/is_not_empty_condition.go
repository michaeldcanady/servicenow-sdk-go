package query

import ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func IsNotEmptyCondition() func(string) ast.Node {
	return unaryCondition(ast.OperatorIsNotEmpty)
}
