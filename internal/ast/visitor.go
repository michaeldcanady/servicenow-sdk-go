package ast

type Visitor[T any] interface {
	Visit(T)
}
