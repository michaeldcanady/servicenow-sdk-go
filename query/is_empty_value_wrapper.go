package query

func IsEmpty() func(string) Node {
	return valueWrapper2(Operator("ISEMPTY"), nil)
}
