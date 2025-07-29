package query

import (
	"fmt"

	ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

func valueWrapper2(operator ast.Operator, value any) func(string) ast.Node {
	return func(field string) ast.Node {
		return &ast.BinaryNode{
			LeftExpression: &ast.LiteralNode{
				Position: 0,
				Value:    field,
				Kind:     ast.KindString,
			},
			Operator: operator,
			RightExpression: &ast.LiteralNode{
				Position: 0,
				Value:    fmt.Sprintf("%v", value),
				Kind:     ast.KindString,
			},
		}
	}
}
