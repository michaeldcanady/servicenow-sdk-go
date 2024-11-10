package core

import (
	"fmt"
	"strings"
)

// Query represents a ServiceNow query and its conditions.
type Query struct {
	head    *Fragment
	tail    *Fragment
	orderBy *OrderBy
}

// NewQuery returns a new Query with no conditions.
func NewQuery() *Query {
	return &Query{
		head:    nil,
		tail:    nil,
		orderBy: NewOrderBy(),
	}
}

// AddFragment adds a fragment to the query.
func (q *Query) AddFragment(fragment *Fragment, operator LogicalOperator) *Query {
	if q.head == nil {
		q.head = fragment
	}

	if q.tail != nil {
		q.tail.SetNext(fragment, operator)
	}

	q.tail = fragment

	return q
}

// AddQuery adds a query condition to the query.
func (q *Query) AddQuery(field string, operator RelationalOperator, value interface{}) *Query {
	q.AddFragment(NewFragment(field, operator, value), And)
	return q
}

// AddEqual adds an equality condition to the query.
func (q *Query) AddEqual(field string, value interface{}) *Query {
	q.AddQuery(field, Is, value)
	return q
}

// AddNotEqual adds a not-equal condition to the query.
func (q *Query) AddNotEqual(field string, value interface{}) *Query {
	q.AddQuery(field, IsNot, value)
	return q
}

// AddGreaterThan adds a greater-than condition to the query.
func (q *Query) AddGreaterThan(field string, value interface{}) *Query {
	q.AddQuery(field, GreaterThan, value)
	return q
}

// AddLessThan adds a less-than condition to the query.
func (q *Query) AddLessThan(field string, value interface{}) *Query {
	q.AddQuery(field, LessThan, value)
	return q
}

// AddContains adds a contains condition to the query.
func (q *Query) AddContains(field string, value interface{}) *Query {
	q.AddQuery(field, Contains, value)
	return q
}

// AddNotContains adds a not-contains condition to the query.
func (q *Query) AddNotContains(field string, value interface{}) *Query {
	q.AddQuery(field, NotContains, value)
	return q
}

// AddStartsWith adds a starts-with condition to the query.
func (q *Query) AddStartsWith(field string, value interface{}) *Query {
	q.AddQuery(field, StartsWith, value)
	return q
}

// AddEndsWith adds an ends-with condition to the query.
func (q *Query) AddEndsWith(field string, value interface{}) *Query {
	q.AddQuery(field, EndsWith, value)
	return q
}

// AddBetween adds a between condition to the query.
func (q *Query) AddBetween(field string, start, end interface{}) *Query {
	value := fmt.Sprintf("%v@%v", start, end)
	q.AddQuery(field, Between, value)
	return q
}

// AddIsSame adds a condition to check if two fields are the same.
func (q *Query) AddIsSame(startField, endField string) *Query {
	q.AddQuery(startField, IsSame, endField)
	return q
}

// AddIsDifferent adds a condition to check if two fields are different.
func (q *Query) AddIsDifferent(startField, endField string) *Query {
	q.AddQuery(startField, IsDifferent, endField)
	return q
}

// IsEmpty adds an "is empty" condition to the query.
func (q *Query) IsEmpty(field string) *Query {
	q.AddQuery(field, IsEmpty, "")
	return q
}

// AddOrQuery adds an "OR" query condition to the query.
func (q *Query) AddOrQuery(field string, operator RelationalOperator, value interface{}) *Query {
	q.AddFragment(NewFragment(field, operator, value), Or)
	return q
}

// AddOrEqual adds an "OR" equality condition to the query.
func (q *Query) AddOrEqual(field string, value interface{}) *Query {
	q.AddOrQuery(field, Is, value)
	return q
}

// AddOrNotEqual adds an "OR" not-equal condition to the query.
func (q *Query) AddOrNotEqual(field string, value interface{}) *Query {
	q.AddOrQuery(field, IsNot, value)
	return q
}

// AddOrGreaterThan adds an "OR" greater-than condition to the query.
func (q *Query) AddOrGreaterThan(field string, value interface{}) *Query {
	q.AddOrQuery(field, GreaterThan, value)
	return q
}

// AddOrLessThan adds an "OR" less-than condition to the query.
func (q *Query) AddOrLessThan(field string, value interface{}) *Query {
	q.AddOrQuery(field, LessThan, value)
	return q
}

// AddOrContains adds an "OR" contains condition to the query.
func (q *Query) AddOrContains(field string, value interface{}) *Query {
	q.AddOrQuery(field, Contains, value)
	return q
}

// AddOrNotContains adds an "OR" not-contains condition to the query.
func (q *Query) AddOrNotContains(field string, value interface{}) *Query {
	q.AddOrQuery(field, NotContains, value)
	return q
}

// AddOrderBy sets the order-by field in ascending order.
func (q *Query) AddOrderBy(field string) *Query {
	q.orderBy.Direction = Asc
	q.orderBy.Field = field
	return q
}

// AddOrderByDesc sets the order-by field in descending order.
func (q *Query) AddOrderByDesc(field string) *Query {
	q.orderBy.Direction = Desc
	q.orderBy.Field = field
	return q
}

// Encoded returns the encoded query as a string.
func (q *Query) Encoded() string {
	// Implement query encoding logic here
	return ""
}

// String returns the query as a string.
func (q *Query) String() string {
	var sb strings.Builder // Create a new string builder

	if q.head == nil {
		return ""
	}

	q.head.Iterate(func(f *Fragment) bool {
		sb.WriteString(f.String() + string(f.LogicalOperator))

		return true
	})

	sb.WriteString(q.orderBy.String())

	return sb.String()
}
