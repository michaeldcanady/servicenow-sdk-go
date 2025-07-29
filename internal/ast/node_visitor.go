package ast

type LiteralNodeVisitor interface {
	VisitLiteralNode(*LiteralNode)
}

type BinaryNodeVisitor interface {
	VisitBinaryNode(*BinaryNode)
}

type NodeVisitor[T any] interface {
	Visitor[T]
	BinaryNodeVisitor
	LiteralNodeVisitor
}
