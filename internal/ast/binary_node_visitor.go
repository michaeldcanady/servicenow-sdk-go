//go:build preview

package ast

// BinaryNodeVisitor Represents a visitor for a binary node.
type BinaryNodeVisitor interface {
	// VisitBinaryNode Visits the provided binary node.
	VisitBinaryNode(*BinaryNode)
}
