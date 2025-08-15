package ast

// UnaryNodeVisitor represents a visitor to the unary node.
type UnaryNodeVisitor interface {
	// VisitUnaryNode visits the unary node.
	VisitUnaryNode(*UnaryNode)
}
