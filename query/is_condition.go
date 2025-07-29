package query

import (
	"fmt"

	ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

type Primitive interface {
	Numeric | ~string
}

func IsCondition[T Primitive](val T) func(string) ast.Node {
	return Condition(ast.OperatorIs, fmt.Sprintf("%v", val))
}
