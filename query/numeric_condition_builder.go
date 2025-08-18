//go:build preview.query

package query

import (
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

// NumericConditionBuilder represents a condition builder for a numeric field.
type NumericConditionBuilder struct {
	*SharedConditionBuilder[float64]
}

// NewNumericConditionBuilder instantiates a new numeric condition builder.
func NewNumericConditionBuilder(field string, query *QueryBuilder) *NumericConditionBuilder {
	return &NumericConditionBuilder{
		NewSharedConditionBuilder[float64](field, query),
	}
}

// LessThan The numeric field is less than the provided value.
func (builder *NumericConditionBuilder) LessThan(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorLessThan, ast.NewLiteralNode(value))
}

// GreaterThan The numeric field is greater than the provided value.
func (builder *NumericConditionBuilder) GreaterThan(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorGreaterThan, ast.NewLiteralNode(value))
}

// LessThanOrIs The numeric field is less than or is the provided value.
func (builder *NumericConditionBuilder) LessThanOrIs(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorLessThanOrIs, ast.NewLiteralNode(value))
}

// GreaterThanOrIs The numeric field is greater than or is the provided value.
func (builder *NumericConditionBuilder) GreaterThanOrIs(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorGreaterThanOrIs, ast.NewLiteralNode(value))
}

// Between The numeric field is between provided lower and upper values.
func (builder *NumericConditionBuilder) Between(lower, upper float64) *QueryBuilder {
	if lower >= upper {
		builder.addErrors(fmt.Errorf("%v is greater or equal to %v", lower, upper))
	}

	return builder.binaryCondition(ast.OperatorBetween, ast.NewPairNode(
		ast.NewLiteralNode(lower),
		ast.NewLiteralNode(upper),
	))
}

// IsDifferent The numeric field is different the provided value.
func (builder *NumericConditionBuilder) IsDifferent(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsDifferent, ast.NewLiteralNode(value))
}

// GreaterThanField The numeric field is greater than the provided field.
func (builder *NumericConditionBuilder) GreaterThanField(sysID string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorGreaterThanField, ast.NewLiteralNode(sysID))
}

// LessThanField The numeric field is less than the provided field.
func (builder *NumericConditionBuilder) LessThanField(sysID string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorLessThanField, ast.NewLiteralNode(sysID))
}

// GreaterThanOrIsField The numeric field is greater than or is the provided field.
func (builder *NumericConditionBuilder) GreaterThanOrIsField(sysID string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorGreaterThanOrIsField, ast.NewLiteralNode(sysID))
}

// LessThanOrIsField The numeric field is less than or is the provided field.
func (builder *NumericConditionBuilder) LessThanOrIsField(sysID string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorLessThanOrIsField, ast.NewLiteralNode(sysID))
}

// IsMoreThan The numeric field is more than or is the provided value.
func (builder *NumericConditionBuilder) IsMoreThan(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsMoreThan, ast.NewLiteralNode(value))
}

// IsLessThan The numeric field is less than or is the provided value.
func (builder *NumericConditionBuilder) IsLessThan(value float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsLessThan, ast.NewLiteralNode(value))
}

// IsOneOf The numeric field is one of the provided values.
func (builder *NumericConditionBuilder) IsOneOf(values ...float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsOneOf, convertSliceToArrayNode(values...))
}

// IsNotOneOf The numeric field is not one of the provided values.
func (builder *NumericConditionBuilder) IsNotOneOf(values ...float64) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsNotOneOf, convertSliceToArrayNode(values...))
}
