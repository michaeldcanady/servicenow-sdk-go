package query

import "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func Between[T Numeric](value1, value2 T) func(string) ast.Node {
	return PairCondition(ast.OperatorBetween, value1, value2)
}
