package query

func LessThan[T Numeric](val T) func(string) Node {
	return valueWrapper2(Operator("<"), val)
}
