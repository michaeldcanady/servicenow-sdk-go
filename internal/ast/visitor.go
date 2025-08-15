package ast

// Visitor Represents
type Visitor[T any] interface {
	// Visit Visits the provided value.
	Visit(T)
}
