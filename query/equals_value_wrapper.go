package query

import "fmt"

type Primitive interface {
	Numeric | ~string
}

type equalsWrapper[T Primitive] struct {
	value T
}

func Equals[T Primitive](val T) ValueWrapper {
	return &equalsWrapper[T]{value: val}
}

func (g *equalsWrapper[T]) ToCondition(field string) Condition {
	return Condition{
		Field:    field,
		Operator: "=",
		Value:    fmt.Sprintf("%v", g.value),
	}
}
