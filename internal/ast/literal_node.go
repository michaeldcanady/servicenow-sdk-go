//go:build preview.query

package ast

import (
	"fmt"
)

// LiteralNode represents a literal value (e.g., "hello", 1, true).
type LiteralNode struct {
	Value string
}

func (n *LiteralNode) Accept(v Visitor) { v.VisitLiteral(n) }

// NewLiteralNode creates a new LiteralNode for the given value.
func NewLiteralNode(val any) *LiteralNode {
	return &LiteralNode{
		Value: fmt.Sprintf("%v", val),
	}
}
