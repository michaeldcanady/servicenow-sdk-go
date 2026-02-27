//go:build preview.query

package query2

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

// StringField represents a string field in ServiceNow.
type StringField struct {
	BaseField
}

// Is query that field is the provided value.
func (f StringField) Is(val string) Condition {
	return f.binary(ast.OperatorIs, val)
}

// IsNot query that field is not the provided value.
func (f StringField) IsNot(val string) Condition {
	return f.binary(ast.OperatorIsNot, val)
}

// StartsWith query that string field starts with the provided value.
func (f StringField) StartsWith(val string) Condition {
	return f.binary(ast.OperatorStartsWith, val)
}

// EndsWith query that string field ends with the provided value.
func (f StringField) EndsWith(val string) Condition {
	return f.binary(ast.OperatorEndsWith, val)
}

// Contains query that string field contains the provided value.
func (f StringField) Contains(val string) Condition {
	return f.binary(ast.OperatorContains, val)
}

// DoesNotContain query that string field does not contain the provided value.
func (f StringField) DoesNotContain(val string) Condition {
	return f.binary(ast.OperatorDoesNotContain, val)
}

// IsOneOf query that field is one of the provided values.
func (f StringField) IsOneOf(values ...string) Condition {
	return f.multi(ast.OperatorIsOneOf, convertSliceToArrayNode(values...))
}

// IsNotOneOf query that field is not one of the provided values.
func (f StringField) IsNotOneOf(values ...string) Condition {
	return f.multi(ast.OperatorIsNotOneOf, convertSliceToArrayNode(values...))
}

// IsEmptyString query that string field is empty.
func (f StringField) IsEmptyString() Condition {
	return f.unary(ast.OperatorIsEmptyString)
}

// MatchesPattern query that string field matches the provided pattern.
func (f StringField) MatchesPattern(pattern string) Condition {
	return f.binary(ast.OperatorMatchesPattern, pattern)
}

// Between query that string field is between the provided lower and upper values.
func (f StringField) Between(lower, upper string) Condition {
	return f.pair(ast.OperatorBetween, lower, upper)
}
