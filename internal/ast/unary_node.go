package ast

// UnaryNode represents a unary operation (e.g., fieldISEMPTY).
type UnaryNode struct {
	Op   Operator
	Left Node
}

// Accept visits the node. It is a no-op on a nil receiver so that malformed
// trees (e.g., those embedding rejected subtrees) can be traversed safely by
// any visitor.
func (n *UnaryNode) Accept(v Visitor) {
	if n == nil {
		return
	}
	v.VisitUnary(n)
}

// NewUnaryNode creates a new UnaryNode with the given operator and left node.
func NewUnaryNode(op Operator, left Node) *UnaryNode {
	return &UnaryNode{Op: op, Left: left}
}
