package query

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

func IsSameCondition(field string) func(string) ast.Node {
	return BinaryCondition(ast.OperatorIsSame, field)
}
