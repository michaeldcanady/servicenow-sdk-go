package ast

var _ Node = (*BinaryNode)(nil)

// BinaryNode Represents a binary expression node.
type BinaryNode struct {
	// LeftExpression The left-hand side expression of the binary operation.
	LeftExpression Node
	// Operator The binary operator.
	Operator Operator
	// Position The position of the operator within the source text.
	Position int
	// RightExpression The right-hand side expression of the binary operation.
	RightExpression Node
}

// Left The leftmost (starting) position of the node in source text.
func (expr *BinaryNode) Left() int {
	if expr.LeftExpression == nil {
		return -1
	}
	return expr.LeftExpression.Pos()
}

// Right The rightmost (ending) position of the node in source text.
func (expr *BinaryNode) Right() int {
	if expr.RightExpression == nil {
		return -1
	}
	return expr.RightExpression.Pos()
}

// Pos The actual position of the node.
func (expr *BinaryNode) Pos() int {
	return expr.Position
}

// Accept Accepts the provided visitor.
func (expr *BinaryNode) Accept(visitor NodeVisitor) {
	visitor.VisitBinaryNode(expr)
}
