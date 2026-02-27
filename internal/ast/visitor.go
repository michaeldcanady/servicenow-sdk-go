//go:build preview.query

package ast

// Visitor represents a visitor for the AST.
type Visitor interface {
	VisitLiteral(node *LiteralNode)
	VisitUnary(node *UnaryNode)
	VisitBinary(node *BinaryNode)
	VisitPair(node *PairNode)
	VisitArray(node *ArrayNode)
}
