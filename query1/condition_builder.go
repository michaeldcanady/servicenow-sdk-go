package query

import (
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

type ConditionBuilder struct {
	field string
	query *QueryBuilder
}

func NewConditionBuilder(field string, query *QueryBuilder) *ConditionBuilder {
	return &ConditionBuilder{
		field: field,
		query: query,
	}
}

func (builder *ConditionBuilder) Is(value any) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIs, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *ConditionBuilder) unaryCondition(operator ast.Operator) *QueryBuilder {
	condition := &ast.UnaryNode{
		Node: &ast.LiteralNode{
			Value: builder.field,
		},
		Operator: operator,
	}

	return builder.query.addCondition(condition)
}

func (builder *ConditionBuilder) binaryCondition(operator ast.Operator, value ast.Node) *QueryBuilder {
	condition := &ast.BinaryNode{
		LeftExpression: &ast.LiteralNode{
			Value: builder.field,
		},
		Operator:        operator,
		RightExpression: value,
	}

	return builder.query.addCondition(condition)
}
