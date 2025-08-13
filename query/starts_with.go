package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func StartWith(substring string) func(string) ast.Node {
	return BinaryCondition(ast.OperatorStartsWith, substring)
}
