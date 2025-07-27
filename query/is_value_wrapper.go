package query

import "fmt"

type Primitive interface {
	Numeric | ~string
}

func Is[T Primitive](val T) func(string) Node {
	return valueWrapper2(Operator("="), fmt.Sprintf("%v", val))
}
