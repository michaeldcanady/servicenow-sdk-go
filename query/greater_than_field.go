package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func GreaterThanField(field string) func(string) ast.Node {
	return BinaryCondition(ast.OperatorGreaterThanField, field)
}
