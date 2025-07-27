package query

func LessThan[T Numeric](val T) func(string) *Condition {
	return valueWrapper1("<", val)
}
