//go:build preview.query

package query

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

// DateTimeConditionBuilder represents a condition builder for a date-time field.
type DateTimeConditionBuilder struct {
	conBuilder[time.Time, QueryBuilder]
}

// NewDateTimeConditionBuilder  a new date-time condition builder, of the provided query, for the provided field.
func NewDateTimeConditionBuilder(field string, query *QueryBuilder) *DateTimeConditionBuilder {
	return &DateTimeConditionBuilder{
		NewSharedConditionBuilder[time.Time](field, query),
	}
}

// On query the date-time field is on a specific date-time.
func (builder *DateTimeConditionBuilder) On(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorOn, ast.NewLiteralNode(dateTime.String()))
}

// NotOn query the date-time field is not on a specific date-time.
func (builder *DateTimeConditionBuilder) NotOn(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorNotOn, ast.NewLiteralNode(dateTime.String()))
}

// Before query the date-time field is before a specific date-time.
func (builder *DateTimeConditionBuilder) Before(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorBefore, ast.NewLiteralNode(dateTime.String()))
}

// AtOrBefore query the date-time field is at or before a specific date-time.
func (builder *DateTimeConditionBuilder) AtOrBefore(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorAtOrBefore, ast.NewLiteralNode(dateTime.String()))
}

// AtOrAfter query the date-time field is at or after a specific date-time.
func (builder *DateTimeConditionBuilder) AtOrAfter(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorAtOrAfter, ast.NewLiteralNode(dateTime.String()))
}

// After query the date-time field is after a specific date-time.
func (builder *DateTimeConditionBuilder) After(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorAfter, ast.NewLiteralNode(dateTime.String()))
}

// TrendOnOrAfter query the date-time field that trends on or after a specific date-time.
func (builder *DateTimeConditionBuilder) TrendOnOrAfter(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorTrendOnOrAfter, ast.NewLiteralNode(dateTime.String()))
}

// TrendOnOrBefore query the date-time field that trends on or before a specific date-time.
func (builder *DateTimeConditionBuilder) TrendOnOrBefore(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorTrendOnOrBefore, ast.NewLiteralNode(dateTime.String()))
}

// TrendAfter query the date-time field that trends after a specific date-time.
func (builder *DateTimeConditionBuilder) TrendAfter(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorTrendAfter, ast.NewLiteralNode(dateTime.String()))
}

// TrendBefore query the date-time field that trends before a specific date-time.
func (builder *DateTimeConditionBuilder) TrendBefore(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorTrendBefore, ast.NewLiteralNode(dateTime.String()))
}

// TrendOn query the date-time field that trends on a specific date-time.
func (builder *DateTimeConditionBuilder) TrendOn(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorTrendOn, ast.NewLiteralNode(dateTime.String()))
}

// RelativeAfter query the date-time field that is relatively after a specific date-time.
func (builder *DateTimeConditionBuilder) RelativeAfter(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorRelativeAfter, ast.NewLiteralNode(dateTime.String()))
}

// RelativeBefore query the date-time field that is relatively before a specific date-time.
func (builder *DateTimeConditionBuilder) RelativeBefore(dateTime time.Time) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorRelativeBefore, ast.NewLiteralNode(dateTime.String()))
}
