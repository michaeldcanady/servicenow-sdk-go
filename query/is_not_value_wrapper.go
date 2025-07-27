package query

func IsNot[T Primitive](val T) func(string) Node {
	return valueWrapper2(Operator("!="), val)
}
