package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func IsEmptyString() func(string) ast.Node {
	return unaryCondition(ast.OperatorContains)
}
