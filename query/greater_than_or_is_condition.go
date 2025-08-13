package query

import (
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

func GreaterThanOrIs[T Numeric](val T) func(string) ast.Node {
	return BinaryCondition(ast.OperatorGreaterThanOrIs, fmt.Sprintf("%v", val))
}
