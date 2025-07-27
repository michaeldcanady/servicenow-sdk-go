package query

import "fmt"

func valueWrapper1(operator string, value any) func(string) *Condition {
	return func(field string) *Condition {
		return NewCondition(field, operator, value)
	}
}

func valueWrapper2(operator Operator, value any) func(string) Node {
	return func(field string) Node {
		return &BinaryExpression{
			LeftExpression: &Literal{
				Position: 0,
				Value:    field,
				Kind:     KindString,
			},
			Operator: operator,
			RightExpression: &Literal{
				Position: 0,
				Value:    fmt.Sprintf("%v", value),
				Kind:     KindString,
			},
		}
	}
}
