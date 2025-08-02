package query

import ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func IsAnythingCondition() func(string) ast.Node {
	return BinaryCondition(ast.OperatorIsAnything, nil)
}
