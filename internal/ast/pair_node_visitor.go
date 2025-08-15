//go:build preview

package ast

// PairNodeVisitor represents a visitor to a pair node
type PairNodeVisitor interface {
	// VisitPairNode visits the provided pair node
	VisitPairNode(*PairNode)
}
