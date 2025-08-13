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

func unaryCondition(operator ast.Operator) func(string) ast.Node {
	return func(field string) ast.Node {
		return &ast.UnaryNode{
			Operator: operator,
			Node: &ast.LiteralNode{
				Position: 0,
				Value:    field,
			},
		}
	}
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

func PairCondition[T Numeric](operator ast.Operator, value1, value2 T) func(string) ast.Node {
	node := &ast.PairNode{
		Element1: &ast.LiteralNode{
			Position: 0,
			Value:    fmt.Sprintf("%v", value1),
		},
		Element2: &ast.LiteralNode{
			Position: 0,
			Value:    fmt.Sprintf("%v", value2),
		},
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
