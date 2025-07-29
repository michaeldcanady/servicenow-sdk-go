package ast

type Accepter[T any, V Visitor[T]] interface {
	Accept(V)
}
