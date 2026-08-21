package ast

type LiteralVisitor interface {
	VisitLiteral(node *LiteralNode)
}

type UnaryVisitor interface {
	VisitUnary(node *UnaryNode)
}

// Visitor represents a visitor for the AST.
type Visitor interface {
	LiteralVisitor
	UnaryVisitor
	VisitBinary(node *BinaryNode)
	VisitPair(node *PairNode)
	VisitArray(node *ArrayNode)
}
