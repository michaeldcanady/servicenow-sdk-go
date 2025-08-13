package query

import (
	"fmt"

	ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

func IsCondition[T Primitive](val T) func(string) ast.Node {
	return BinaryCondition(ast.OperatorIs, fmt.Sprintf("%v", val))
}
