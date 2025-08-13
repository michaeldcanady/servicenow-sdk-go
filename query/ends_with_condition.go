package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func EndsWith(substring string) func(string) ast.Node {
	return BinaryCondition(ast.OperatorEndsWith, substring)
}
