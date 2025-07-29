package query

import (
	"fmt"

	ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

type Primitive interface {
	Numeric | ~string
}

func Is[T Primitive](val T) func(string) ast.Node {
	return valueWrapper2(ast.Operator("="), fmt.Sprintf("%v", val))
}
