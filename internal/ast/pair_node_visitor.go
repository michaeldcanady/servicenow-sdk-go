package ast

// PairNodeVisitor represents a visitor to a PairNode
type PairNodeVisitor interface {
	VisitPairNode(*PairNode)
}
