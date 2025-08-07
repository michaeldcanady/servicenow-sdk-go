package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func LessThanField(field string) func(string) ast.Node {
	return BinaryCondition(ast.OperatorLessThanField, field)
}
