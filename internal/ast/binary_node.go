//go:build preview.query

package ast

// BinaryNode represents a binary operation (e.g., field=value, fieldLIKEvalue).
type BinaryNode struct {
	Left  Node
	Op    Operator
	Right Node
}

func (n *BinaryNode) Accept(v Visitor) { v.VisitBinary(n) }

// NewBinaryNode creates a new BinaryNode with the given left node, operator, and right node.
func NewBinaryNode(left Node, op Operator, right Node) *BinaryNode {
	return &BinaryNode{Left: left, Op: op, Right: right}
}
