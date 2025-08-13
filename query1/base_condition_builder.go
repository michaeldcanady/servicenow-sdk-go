package query

import (
	"fmt"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

type BaseConditionBuilder[T Primitive | time.Time] struct {
	field string
	query *QueryBuilder
}

func NewBaseConditionBuilder[T Primitive | time.Time](field string, query *QueryBuilder) *BaseConditionBuilder[T] {
	return &BaseConditionBuilder[T]{
		field: field,
		query: query,
	}
}

func (builder *BaseConditionBuilder[T]) Is(value T) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIs, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *BaseConditionBuilder[T]) IsNot(value T) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsNot, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *BaseConditionBuilder[T]) IsEmpty() *QueryBuilder {
	return builder.unaryCondition(ast.OperatorIsEmpty)
}

func (builder *BaseConditionBuilder[T]) IsNotsEmpty() *QueryBuilder {
	return builder.unaryCondition(ast.OperatorIsNotEmpty)
}

func (builder *BaseConditionBuilder[T]) IsAnything() *QueryBuilder {
	return builder.unaryCondition(ast.OperatorIsAnything)
}

func (builder *BaseConditionBuilder[T]) IsDynamic(sysID string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsNot, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", sysID),
	})
}

func (builder *BaseConditionBuilder[T]) IsSame(sysID string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsSame, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", sysID),
	})
}

func (builder *BaseConditionBuilder[T]) unaryCondition(operator ast.Operator) *QueryBuilder {
	condition := &ast.UnaryNode{
		Node: &ast.LiteralNode{
			Value: builder.field,
		},
		Operator: operator,
	}

	return builder.query.addCondition(condition)
}

func (builder *BaseConditionBuilder[T]) binaryCondition(operator ast.Operator, value ast.Node) *QueryBuilder {
	condition := &ast.BinaryNode{
		LeftExpression: &ast.LiteralNode{
			Value: builder.field,
		},
		Operator:        operator,
		RightExpression: value,
	}

	return builder.query.addCondition(condition)
}
