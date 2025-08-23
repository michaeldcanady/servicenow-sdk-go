//go:build preview.query

package ast

// ExpressionNode represents an expression of (a) node(s).
type ExpressionNode interface {
	Node
	// Operator returns the expression operator.
	Operator() Operator
}
