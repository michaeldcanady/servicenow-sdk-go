package query

import "fmt"

type Primitive interface {
	Numeric | ~string
}

func Is[T Primitive](val T) func(string) *Condition {
	return valueWrapper1("=", fmt.Sprintf("%v", val))
}
