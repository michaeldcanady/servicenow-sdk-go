//go:build preview.query

package ast

// UnaryNode represents a unary operation (e.g., fieldISEMPTY).
type UnaryNode struct {
	Op   Operator
	Left Node
}

func (n *UnaryNode) Accept(v Visitor) { v.VisitUnary(n) }

// NewUnaryNode creates a new UnaryNode with the given operator and left node.
func NewUnaryNode(op Operator, left Node) *UnaryNode {
	return &UnaryNode{Op: op, Left: left}
}
