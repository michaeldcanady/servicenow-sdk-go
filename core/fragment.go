package core

import "fmt"

// Fragment represents a query fragment with a field, operator, and value.
type Fragment struct {
	// Field represents the field on which the condition is applied.
	Field string
	// RelationalOperator represents the comparison operator for the condition.
	RelationalOperator RelationalOperator
	// Value represents the value to compare against.
	Value interface{}
	// LogicalOperator represents the operator connection to the next fragment
	LogicalOperator LogicalOperator
	next            *Fragment
}

// NewFragment creates a new query fragment with the specified field, operator, and value.
func NewFragment(field string, operator RelationalOperator, value interface{}) *Fragment {
	return &Fragment{
		Field:              field,
		RelationalOperator: operator,
		Value:              value,
		next:               nil,
	}
}

// SetNext sets the next and Logical Operator value
func (f *Fragment) SetNext(fragment *Fragment, operator LogicalOperator) {
	f.next = fragment
	f.LogicalOperator = operator
}

// Iterate iterates over the fragments
func (f *Fragment) Iterate(callback func(*Fragment) bool) {
	current := f

	for current != nil {
		if !callback(current) {
			break
		}
		current = current.next
	}
}

// String returns a string representation of the query fragment in the format "field operator value".
func (f *Fragment) String() string {
	value := f.Value
	if f.Value == nil {
		value = ""
	}

	return fmt.Sprintf("%s%s%v", f.Field, f.RelationalOperator, value)
}
