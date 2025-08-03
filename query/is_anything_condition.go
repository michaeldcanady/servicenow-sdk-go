package query

import ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func IsAnythingCondition() func(string) ast.Node {
	return unaryCondition(ast.OperatorIsAnything)
}
