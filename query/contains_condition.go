package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func Contains(substring string) func(string) ast.Node {
	return BinaryCondition(ast.OperatorContains, substring)
}
