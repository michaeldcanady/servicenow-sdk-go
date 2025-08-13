package query

import (
	"fmt"

	ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

func IsMoreThanCondition[T Numeric](val T) func(string) ast.Node {
	return BinaryCondition(ast.OperatorIsMoreThan, fmt.Sprintf("%v", val))
}
