package ast

// LiteralNodeVisitor Represents a visitor of a literal node
type LiteralNodeVisitor[T LiteralNode] interface {
	// VisitLiteralNode visits the provided literal node
	VisitLiteralNode(*T)
}
