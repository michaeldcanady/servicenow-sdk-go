package query

import (
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

type NumericConditionBuilder struct {
	*BaseConditionBuilder[float64]
}

func NewNumericConditionBuilder(field string, query *QueryBuilder) *NumericConditionBuilder {
	return &NumericConditionBuilder{
		NewBaseConditionBuilder[float64](field, query),
	}
}

func (builder *NumericConditionBuilder) LessThan(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorLessThan, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *NumericConditionBuilder) GreaterThan(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorGreaterThan, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *NumericConditionBuilder) LessThanOrIs(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorLessThanOrIs, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *NumericConditionBuilder) GreaterThanOrIs(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorGreaterThanOrIs, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *NumericConditionBuilder) Between(lower, upper float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorBetween, &ast.PairNode{
		Element1: &ast.LiteralNode{
			Value: fmt.Sprintf("%v", lower),
		},
		Element2: &ast.LiteralNode{
			Value: fmt.Sprintf("%v", upper),
		},
	})
}

func (builder *NumericConditionBuilder) IsDifferent(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsDifferent, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *NumericConditionBuilder) GreaterThanField(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorGreaterThanField, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *NumericConditionBuilder) LessThanField(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorLessThanField, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *NumericConditionBuilder) GreaterThanOrIsField(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorGreaterThanOrIsField, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *NumericConditionBuilder) LessThanOrIsField(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorLessThanOrIsField, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *NumericConditionBuilder) IsMoreThan(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsMoreThan, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *NumericConditionBuilder) IsLessThan(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsLessThan, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *NumericConditionBuilder) IsOneOf(values ...float64) *QueryBuilder {
	nodes := make([]ast.Node, len(values))
	for index, value := range values {
		node := ast.LiteralNode{
			Value: fmt.Sprintf("%v", value),
		}
		nodes[index] = &node
	}

	return builder.binaryCondition(ast.OperatorIsOneOf, &ast.ArrayNode{
		Elements: nodes,
	})
}

func (builder *NumericConditionBuilder) IsNotOneOf(values ...float64) *QueryBuilder {
	nodes := make([]ast.Node, len(values))
	for index, value := range values {
		node := ast.LiteralNode{
			Value: fmt.Sprintf("%v", value),
		}
		nodes[index] = &node
	}

	return builder.binaryCondition(ast.OperatorIsNotOneOf, &ast.ArrayNode{
		Elements: nodes,
	})
}
