package query

func IsNot[T Primitive](val T) func(string) *Condition {
	return valueWrapper1("!=", val)
}
