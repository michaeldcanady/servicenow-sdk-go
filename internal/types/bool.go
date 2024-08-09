package types

import "fmt"

type Bool interface {
	Type
}

type boolValue struct {
	b bool
}

func NewBool(b bool) Bool {
	return &boolValue{b}
}

func (b *boolValue) String() string {
	return fmt.Sprint(b.b)
}
