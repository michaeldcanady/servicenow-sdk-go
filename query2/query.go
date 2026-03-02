//go:build preview.query

// Package query2 provides a type-safe and fluent API for building ServiceNow encoded queries.
// It is a redesign of the query package, focusing on usability and immutability.
package query2

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
func And(conds ...Condition) Condition {
	if len(conds) == 0 {
		return nil
	}
	res := conds[0]
	for _, c := range conds[1:] {
		res = res.And(c)
	}
	return res
}

// Or combines multiple conditions with an OR operator.
func Or(conds ...Condition) Condition {
	if len(conds) == 0 {
		return nil
	}
	res := conds[0]
	for _, c := range conds[1:] {
		res = res.Or(c)
	}
	return res
}
