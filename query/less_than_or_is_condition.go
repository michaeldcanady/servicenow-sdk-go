package query

import ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func LessThanOrIsCondition[T Numeric](val T) func(string) ast.Node {
	return BinaryCondition(ast.OperatorLessThanOrIs, val)
}
