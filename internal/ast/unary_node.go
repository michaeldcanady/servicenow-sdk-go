package ast

var _ Node = (*UnaryNode)(nil)

type UnaryNode struct {
	// Operator The unary operator.
	Operator Operator
	// Position The position of the unary operator.
	Position int
	// Node
	Node Node
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

// Left return the left most value of the Node.
func (u *UnaryNode) Left() int {
	return u.Node.Left()
}
