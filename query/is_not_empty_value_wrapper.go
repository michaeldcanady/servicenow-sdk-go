package query

func IsNotEmpty() func(string) *Condition {
	return valueWrapper1("ISNOTEMPTY", nil)
}
