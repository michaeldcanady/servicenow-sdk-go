//go:build preview.query

package ast

type ExpressionNode interface {
	Node
	Operator() Operator
}
