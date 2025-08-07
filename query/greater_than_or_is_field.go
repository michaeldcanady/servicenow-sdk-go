package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func GreaterThanOrIsField(field string) func(string) ast.Node {
	return BinaryCondition(ast.OperatorGreaterThanOrIsField, field)
}
