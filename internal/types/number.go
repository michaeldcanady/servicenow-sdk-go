package types

import "fmt"

type Number interface {
	Type
}

type numberValue struct {
	val float64
}

func NewNumber(val float64) Number {
	return &numberValue{val}
}

func (n *numberValue) String() string {
	return fmt.Sprint(n.val)
}
