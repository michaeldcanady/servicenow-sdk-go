//go:build preview

package ast

// ArrayNodeVisitor represents a visitor for an array node.
type ArrayNodeVisitor interface {
	// VisitArrayNode Visits the provided array node.
	VisitArrayNode(*ArrayNode)
}
