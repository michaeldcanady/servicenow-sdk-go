package types

import "fmt"

type Integer interface {
	Type
}

type integerValue struct {
	val int64
}

func NewInteger(val int64) Integer {
	return &integerValue{val}
}

func (i *integerValue) String() string {
	return fmt.Sprint(i.val)
}
