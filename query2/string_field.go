//go:build preview.query

package query2

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast2"
)

// StringField represents a string field in ServiceNow.
type StringField struct {
	BaseField
}

// Is query that field is the provided value.
func (f StringField) Is(val string) Condition {
	return f.binary(ast2.OperatorIs, val)
}

// IsNot query that field is not the provided value.
func (f StringField) IsNot(val string) Condition {
	return f.binary(ast2.OperatorIsNot, val)
}

// StartsWith query that string field starts with the provided value.
func (f StringField) StartsWith(val string) Condition {
	return f.binary(ast2.OperatorStartsWith, val)
}

// EndsWith query that string field ends with the provided value.
func (f StringField) EndsWith(val string) Condition {
	return f.binary(ast2.OperatorEndsWith, val)
}

// Contains query that string field contains the provided value.
func (f StringField) Contains(val string) Condition {
	return f.binary(ast2.OperatorContains, val)
}

// DoesNotContain query that string field does not contain the provided value.
func (f StringField) DoesNotContain(val string) Condition {
	return f.binary(ast2.OperatorDoesNotContain, val)
}

// IsOneOf query that field is one of the provided values.
func (f StringField) IsOneOf(values ...string) Condition {
	return f.multi(ast2.OperatorIsOneOf, convertSliceToArrayNode(values...))
}

// IsNotOneOf query that field is not one of the provided values.
func (f StringField) IsNotOneOf(values ...string) Condition {
	return f.multi(ast2.OperatorIsNotOneOf, convertSliceToArrayNode(values...))
}

// IsEmptyString query that string field is empty.
func (f StringField) IsEmptyString() Condition {
	return f.unary(ast2.OperatorIsEmptyString)
}

// MatchesPattern query that string field matches the provided pattern.
func (f StringField) MatchesPattern(pattern string) Condition {
	return f.binary(ast2.OperatorMatchesPattern, pattern)
}

// Between query that string field is between the provided lower and upper values.
func (f StringField) Between(lower, upper string) Condition {
	return f.pair(ast2.OperatorBetween, lower, upper)
}
