package ast

// NodeVisitor represents a visitor for a node tree.
type NodeVisitor interface {
	Visitor[Node]
	ArrayNodeVisitor
	BinaryNodeVisitor
	LiteralNodeVisitor
	PairNodeVisitor
	UnaryNodeVisitor
}
