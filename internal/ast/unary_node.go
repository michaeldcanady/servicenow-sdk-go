//go:build preview.query

package ast

var _ Node = (*UnaryNode)(nil)

// UnaryNode represents a node with only an operator and a node
type UnaryNode struct {
	// Op The unary operator.
	Op Operator
	// Position The position of the unary operator.
	Position int
	// Node
	Node Node
}

// NewUnaryNode instantiates a new unary node expression.
func NewUnaryNode(operator Operator, node Node) *UnaryNode {
	return &UnaryNode{
		Op:       operator,
		Position: -1,
		Node:     node,
	}
}

// Accept implements Node.
func (u *UnaryNode) Accept(node NodeVisitor) {
	node.VisitUnaryNode(u)
}

// Right returns the right most position of the operator.
func (u *UnaryNode) Right() int {
	return u.Pos()
}

// Pos returns the position of the operator.
func (u *UnaryNode) Pos() int {
	return u.Position
}

// Left return the left most value.
func (u *UnaryNode) Left() int {
	if u.Node == nil {
		return -1
	}

	return u.Node.Left()
}

// Operator returns the operator value.
func (u *UnaryNode) Operator() Operator {
	return u.Op
}
