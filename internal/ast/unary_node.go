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
	if u.Node == nil {
		return -1
	}

	return u.Node.Left()
}

// Operator returns the expression operator.
func (u *UnaryNode) Operator() Operator {
	return u.Op
}
