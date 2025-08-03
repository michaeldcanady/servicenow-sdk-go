package ast

// NodeVisitor[T] Represents a visit for a node
type NodeVisitor interface {
	Visitor[Node]
	BinaryNodeVisitor
	LiteralNodeVisitor[LiteralNode]
	ArrayNodeVisitor
	PairNodeVisitor
	UnaryNodeVisitor
}
