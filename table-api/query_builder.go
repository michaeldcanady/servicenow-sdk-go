package tableapi

import (
	"fmt"
	"strings"
)

// OrderDirection represents the order direction for sorting.
type OrderDirection int

const (
	Asc  OrderDirection = 1
	Desc OrderDirection = 2
)

// orderBy represents an order-by clause.
type orderBy struct {
	direction OrderDirection
	field     string
}

func (oB *orderBy) String() string {
	str := ""
	switch oB.direction {
	case Asc:
		str += "^ORDERBY"
	case Desc:
		str += "^ORDERBYDESC"
	}
	if str != "" {
		str += oB.field
	}
	return str
}

// Query represents a ServiceNow query and its conditions.
type Query struct {
	fragments []interface{}
	orderBy   *orderBy
}

func newOrderBy() *orderBy {
	return &orderBy{}
}

// NewQuery returns a new Query with no conditions.
func NewQuery() *Query {
	return &Query{
		fragments: []interface{}{},
		orderBy:   newOrderBy(),
	}
}

// AddQuery adds a query condition to the query.
func (q *Query) AddQuery(field string, operator Operator, value interface{}) *Query {
	if len(q.fragments) > 0 {
		q.fragments = append(q.fragments, "^")
	}
	q.fragments = append(q.fragments, field, operator, value)
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
func (q *Query) AddOrQuery(field string, operator Operator, value interface{}) *Query {
	if len(q.fragments) > 0 {
		q.fragments = append(q.fragments, "^OR")
	}
	q.fragments = append(q.fragments, field, operator, value)
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
	q.orderBy.direction = Asc
	q.orderBy.field = field
	return q
}

// AddOrderByDesc sets the order-by field in descending order.
func (q *Query) AddOrderByDesc(field string) *Query {
	q.orderBy.direction = Desc
	q.orderBy.field = field
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
	for _, v := range append(q.fragments, q.orderBy.String()) {
		sb.WriteString(fmt.Sprint(v)) // Convert each element to a string and append
	}
	return sb.String() // Get the final string
}
