package query

func IsNotEmpty() func(string) Node {
	return valueWrapper2(Operator("ISNOTEMPTY"), nil)
}
