package query

import (
	"fmt"

	ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

func BinaryCondition(operator ast.Operator, value any) func(string) ast.Node {
	return nodeCondition(operator, &ast.LiteralNode{
		Position: 0,
		Value:    fmt.Sprintf("%v", value),
		Kind:     ast.KindString,
	})
}

func ArrayCondition[T Primitive](operator ast.Operator, values []T) func(string) ast.Node {
	node := &ast.ArrayNode{
		LeftBrace:  0,
		Elements:   make([]ast.Node, len(values)),
		RightBrace: 0,
	}

	for index, value := range values {
		node.Elements[index] = &ast.LiteralNode{
			Position: 0,
			Value:    fmt.Sprintf("%v", value),
		}
	}

	return nodeCondition(operator, node)
}

func nodeCondition(operator ast.Operator, node ast.Node) func(string) ast.Node {
	return func(field string) ast.Node {
		return &ast.BinaryNode{
			LeftExpression: &ast.LiteralNode{
				Position: 0,
				Value:    field,
				Kind:     ast.KindString,
			},
			Operator:        operator,
			RightExpression: node,
		}
	}
}
