package query

func IsEmpty() func(string) *Condition {
	return valueWrapper1("ISEMPTY", nil)
}
