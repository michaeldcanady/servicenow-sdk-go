package query

import "fmt"

type greaterThanWrapper[T Numeric] struct {
	value T
}

func GreaterThan[T Numeric](val T) ValueWrapper {
	return &greaterThanWrapper[T]{value: val}
}

func (g *greaterThanWrapper[T]) ToCondition(field string) Condition {
	return Condition{
		Field:    field,
		Operator: ">",
		Value:    fmt.Sprintf("%v", g.value),
	}
}
