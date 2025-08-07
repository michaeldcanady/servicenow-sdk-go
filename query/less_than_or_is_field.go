package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func LessThanOrIsField(field string) func(string) ast.Node {
	return BinaryCondition(ast.OperatorLessThanOrIsField, field)
}
