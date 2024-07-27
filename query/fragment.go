package query

import "fmt"

type fragment interface {
	setNext(fragment fragment, operator logicalOperator)
	getNext() fragment
	setLogicalOperator(operator logicalOperator)
	getLogicalOperator() logicalOperator
	String() string
}

// fragmentImpl represents a query fragmentImpl with a field, operator, and value.
type fragmentImpl[T any] struct {
	// Field represents the field on which the condition is applied.
	Field string
	// RelationalOperator represents the comparison operator for the condition.
	RelationalOperator relationalOperator
	// Value represents the value to compare against.
	Value *T
	// LogicalOperator represents the operator connection to the next fragment
	LogicalOperator logicalOperator
	next            fragment
}

// newFragment creates a new query fragment with the specified field, operator, and value.
func newFragment[t any](field string, operator relationalOperator, value *t) fragment {
	return &fragmentImpl[t]{
		Field:              field,
		RelationalOperator: operator,
		Value:              value,
		LogicalOperator:    unset,
		next:               nil,
	}
}

// setNext sets the next and Logical Operator value
func (f *fragmentImpl[t]) setNext(fragment fragment, operator logicalOperator) {
	f.next = fragment
	f.LogicalOperator = operator
}

func (f *fragmentImpl[T]) getNext() fragment {
	return f.next
}

func (f *fragmentImpl[T]) setLogicalOperator(operator logicalOperator) {
	f.LogicalOperator = operator
}

func (f *fragmentImpl[T]) getLogicalOperator() logicalOperator {
	return f.LogicalOperator
}

// String returns a string representation of the query fragment in the format "field operator value".
func (f *fragmentImpl[t]) String() string {
	value := f.Value
	if f.Value == nil {
		return fmt.Sprintf("%s%s", f.Field, f.RelationalOperator)
	}

	return fmt.Sprintf("%s%s%v", f.Field, f.RelationalOperator, *value)
}
