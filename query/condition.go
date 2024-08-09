package query

import "fmt"

type condition[T any] struct {
	field    string
	operator relationalOperator
	value    *T
}

func newCondition[T any](operand1 string, oper relationalOperator, operand2 *T) condition[T] {
	return condition[T]{
		field:    operand1,
		operator: oper,
		value:    operand2,
	}
}

func (c condition[T]) String() string {
	//TODO: add handling for time
	if c.value == nil {
		return fmt.Sprintf("%v%v", c.field, c.operator)
	}
	return fmt.Sprintf("%v%v%v", c.field, c.operator, c.value)
}
