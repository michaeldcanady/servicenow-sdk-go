package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func IsDifferent(field string) func(string) ast.Node {
	return BinaryCondition(ast.OperatorIsDifferent, field)
}
