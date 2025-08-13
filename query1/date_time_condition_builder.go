package query

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

type DateTimeConditionBuilder struct {
	*BaseConditionBuilder[time.Time]
}

func NewDateTimeConditionBuilder(name string, query *QueryBuilder) *DateTimeConditionBuilder {
	return &DateTimeConditionBuilder{
		NewBaseConditionBuilder[time.Time](name, query),
	}
}

func (builder *DateTimeConditionBuilder) On(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorOn, &ast.LiteralNode{
		Value: dateTime.String(),
	})
}

func (builder *DateTimeConditionBuilder) NotOn(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorNotOn, &ast.LiteralNode{
		Value: dateTime.String(),
	})
}

func (builder *DateTimeConditionBuilder) Before(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorBefore, &ast.LiteralNode{
		Value: dateTime.String(),
	})
}

func (builder *DateTimeConditionBuilder) AtOrBefore(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorAtOrBefore, &ast.LiteralNode{
		Value: dateTime.String(),
	})
}

func (builder *DateTimeConditionBuilder) AtOrAfter(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorAtOrAfter, &ast.LiteralNode{
		Value: dateTime.String(),
	})
}

func (builder *DateTimeConditionBuilder) After(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorAfter, &ast.LiteralNode{
		Value: dateTime.String(),
	})
}

func (builder *DateTimeConditionBuilder) TrendOnOrAfter(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorTrendOnOrAfter, &ast.LiteralNode{
		Value: dateTime.String(),
	})
}

func (builder *DateTimeConditionBuilder) TrendOnOrBefore(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorTrendOnOrBefore, &ast.LiteralNode{
		Value: dateTime.String(),
	})
}

func (builder *DateTimeConditionBuilder) TrendAfter(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorTrendAfter, &ast.LiteralNode{
		Value: dateTime.String(),
	})
}

func (builder *DateTimeConditionBuilder) TrendBefore(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorTrendBefore, &ast.LiteralNode{
		Value: dateTime.String(),
	})
}

func (builder *DateTimeConditionBuilder) TrendOn(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorTrendOn, &ast.LiteralNode{
		Value: dateTime.String(),
	})
}

func (builder *DateTimeConditionBuilder) RelativeAfter(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorRelativeAfter, &ast.LiteralNode{
		Value: dateTime.String(),
	})
}

func (builder *DateTimeConditionBuilder) RelativeBefore(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorRelativeBefore, &ast.LiteralNode{
		Value: dateTime.String(),
	})
}
