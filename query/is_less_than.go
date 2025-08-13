package query

import (
	"fmt"

	ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

func IsLessThanCondition[T Numeric](val T) func(string) ast.Node {
	return BinaryCondition(ast.OperatorIsLessThan, fmt.Sprintf("%v", val))
}
