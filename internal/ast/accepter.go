//go:build preview

package ast

// Accepter[T,V] Represents an accepter of a visitor.
type Accepter[T any, V Visitor[T]] interface {
	// Accept Accepts the provided visitor.
	Accept(V)
}
