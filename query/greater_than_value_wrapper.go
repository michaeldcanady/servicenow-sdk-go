package query

import (
	"fmt"

	ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

func GreaterThan[T Numeric](val T) func(string) ast.Node {
	return valueWrapper2(ast.Operator(">"), fmt.Sprintf("%v", val))
}
