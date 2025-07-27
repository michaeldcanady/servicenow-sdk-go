package query

import "fmt"

func GreaterThan[T Numeric](val T) func(string) *Condition {
	return valueWrapper1(">", fmt.Sprintf("%v", val))
}
