package query

import (
	"fmt"

	ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

func GreaterThanCondition[T Numeric](val T) func(string) ast.Node {
	return Condition(ast.OperatorGreaterThan, fmt.Sprintf("%v", val))
}
