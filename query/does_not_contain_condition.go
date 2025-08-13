package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func DoesNotContain(substring string) func(string) ast.Node {
	return BinaryCondition(ast.OperatorDoesNotContain, substring)
}
