package ast

type NodeVisitor interface {
	Visitor[Node]
	ArrayNodeVisitor
	BinaryNodeVisitor
	LiteralNodeVisitor
	PairNodeVisitor
	UnaryNodeVisitor
}
