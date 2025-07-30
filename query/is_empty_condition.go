package query

import ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func IsEmptyCondition() func(string) ast.Node {
	return BinaryCondition(ast.OperatorIsEmpty, nil)
}
