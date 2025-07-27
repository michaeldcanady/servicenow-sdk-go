package query

func valueWrapper1(operator string, value any) func(string) *Condition {
	return func(field string) *Condition {
		return NewCondition(field, operator, value)
	}
}
