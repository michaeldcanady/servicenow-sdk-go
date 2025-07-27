package query

import "fmt"

func GreaterThan[T Numeric](val T) func(string) Node {
	return valueWrapper2(Operator(">"), fmt.Sprintf("%v", val))
}
