package ast

// BinaryNode represents a binary operation (e.g., field=value, fieldLIKEvalue).
type BinaryNode struct {
	Left  Node
	Op    Operator
	Right Node
}

// Accept visits the node. It is a no-op on a nil receiver so that malformed
// trees (e.g., those embedding rejected subtrees) can be traversed safely by
// any visitor.
func (n *BinaryNode) Accept(v Visitor) {
	if n == nil {
		return
	}
	v.VisitBinary(n)
}

// NewBinaryNode creates a new BinaryNode with the given left node, operator, and right node.
func NewBinaryNode(left Node, op Operator, right Node) *BinaryNode {
	return &BinaryNode{Left: left, Op: op, Right: right}
}
