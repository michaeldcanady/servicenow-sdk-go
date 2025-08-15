//go:build preview

package ast

// LiteralNodeVisitor Represents a visitor of a literal node
type LiteralNodeVisitor interface {
	// VisitLiteralNode visits the provided literal node
	VisitLiteralNode(*LiteralNode)
}
