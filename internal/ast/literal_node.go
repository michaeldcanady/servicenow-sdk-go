package ast

import (
	"fmt"
)

// LiteralNode represents a literal value (e.g., "hello", 1, true).
type LiteralNode struct {
	Value string
}

// Accept visits the node. It is a no-op on a nil receiver so that malformed
// trees (e.g., those embedding rejected subtrees) can be traversed safely by
// any visitor.
func (n *LiteralNode) Accept(v Visitor) {
	if n == nil {
		return
	}
	v.VisitLiteral(n)
}

// NewLiteralNode creates a new LiteralNode for the given value.
func NewLiteralNode(val any) *LiteralNode {
	return &LiteralNode{
		Value: fmt.Sprintf("%v", val),
	}
}
