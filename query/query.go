// Package query provides a type-safe and fluent API for building ServiceNow encoded queries.
// It is a redesign of the query package, focusing on usability and immutability.
//
// Condition values are validated at construction time: a value containing one
// of the reserved encoded-query characters ("^", ",", "@") yields an error
// Condition whose [Condition.Error] is non-nil, because ServiceNow encoded queries offer
// no way to escape those characters inside a value.
package query

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
)

// Where starts a new query on the specified field as a string field.
func Where(name string) StringField {
	return String(name)
}

// String starts a query on a string field.
func String(name string) StringField {
	return StringField{BaseField{name: name}}
}

// Number starts a query on a numeric field.
func Number(name string) NumberField {
	return NumberField{BaseField{name: name}}
}

// Boolean starts a query on a boolean field.
func Boolean(name string) BooleanField {
	return BooleanField{BaseField{name: name}}
}

// Date starts a query on a date-time field.
func Date(name string) DateTimeField {
	return DateTime(name)
}

// DateTime starts a query on a date-time field.
func DateTime(name string) DateTimeField {
	return DateTimeField{BaseField{name: name}}
}

// And combines multiple conditions with an AND operator.
//
// Without any conditions it returns an error Condition matching
// [ErrNoConditions]; a nil element yields one matching [ErrNilCondition].
func And(conds ...Condition) Condition {
	if len(conds) == 0 {
		return NewErrorCondition(ErrNoConditions)
	}
	if conversion.IsNil(conds[0]) {
		return NewErrorCondition(ErrNilCondition)
	}
	res := conds[0]
	for _, c := range conds[1:] {
		res = res.And(c)
	}
	return res
}

// Or combines multiple conditions with an OR operator.
//
// Without any conditions it returns an error Condition matching
// [ErrNoConditions]; a nil element yields one matching [ErrNilCondition].
func Or(conds ...Condition) Condition {
	if len(conds) == 0 {
		return NewErrorCondition(ErrNoConditions)
	}
	if conversion.IsNil(conds[0]) {
		return NewErrorCondition(ErrNilCondition)
	}
	res := conds[0]
	for _, c := range conds[1:] {
		res = res.Or(c)
	}
	return res
}
