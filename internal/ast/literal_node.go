package ast

import (
	"fmt"
)

var _ Node = (*LiteralNode)(nil)

// LiteralNode Represents a literal value node (e.g. numeric constant, string, or keyword) in the expression tree.
type LiteralNode struct {
	// Position The position of the literal in the source text.
	Position int
	// Kind The category or type of literal.
	Kind Kind
	// The literal value itself.
	Value string
}

// NewLiteralNode instantiates a new literal node of the specified value
func NewLiteralNode(value any) *LiteralNode {
	return &LiteralNode{
		Position: -1,
		Kind:     KindUnknown,
		Value:    fmt.Sprintf("%v", value),
	}
}

// Left The leftmost (starting) position of the node in source text.
func (expr *LiteralNode) Left() int {
	return expr.Pos()
}

// Right The rightmost (ending) position of the node in source text.
func (expr *LiteralNode) Right() int {
	return expr.Pos() + len(expr.Value)
}

// Pos The actual position of the node.
func (expr *LiteralNode) Pos() int {
	return expr.Position
}

// Accept Accepts the provided visitor.
func (expr *LiteralNode) Accept(visitor NodeVisitor) {
	visitor.VisitLiteralNode(expr)
}
