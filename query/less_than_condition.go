package query

import ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

func LessThanCondition[T Numeric](val T) func(string) ast.Node {
	return Condition(ast.OperatorLessThan, val)
}
