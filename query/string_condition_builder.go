//go:build preview.query

package query

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

// StringConditionBuilder represents a condition builder for a string field.
type StringConditionBuilder struct {
	*SharedConditionBuilder[string]
}

// NewStringConditionBuilder instantiates a new string condition builder, of the provided query, for the provided field.
func NewStringConditionBuilder(field string, query *QueryBuilder) *StringConditionBuilder {
	return &StringConditionBuilder{
		NewSharedConditionBuilder[string](field, query),
	}
}

// StartsWith query that string field starts with the provided value.
func (builder *StringConditionBuilder) StartsWith(value string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorStartsWith, ast.NewLiteralNode(value))
}

// EndsWith query that string field ends with the provided value.
func (builder *StringConditionBuilder) EndsWith(value string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorEndsWith, ast.NewLiteralNode(value))
}

// Contains query that string field contains the provided value.
func (builder *StringConditionBuilder) Contains(value string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorContains, ast.NewLiteralNode(value))
}

// DoesNotContain query that string field does not contain the provided value.
func (builder *StringConditionBuilder) DoesNotContain(value string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorDoesNotContain, ast.NewLiteralNode(value))
}

// IsEmptyString query that string field is empty.
func (builder *StringConditionBuilder) IsEmptyString() *QueryBuilder {
	return builder.unaryCondition(ast.OperatorIsEmptyString)
}
