package query

import "fmt"

// A single condition in a query
type Condition struct {
	Field    string
	Operator Operator
	Value    any
}

func NewCondition(field string, operator string, value any) *Condition {
	return &Condition{
		Field:    field,
		Operator: Operator(operator),
		Value:    value,
	}
}

func (c Condition) Serialize() string {
	return fmt.Sprintf("%s%s%s", c.Field, c.Operator, c.Value)
}
