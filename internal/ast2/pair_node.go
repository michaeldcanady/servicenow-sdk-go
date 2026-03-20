//go:build preview.query

package ast2

// PairNode represents a pair of values, typically used for BETWEEN.
type PairNode struct {
	Left  Node
	Right Node
}

func (n *PairNode) Accept(v Visitor) { v.VisitPair(n) }

// NewPairNode creates a new PairNode with the given left and right nodes.
func NewPairNode(left, right Node) *PairNode {
	return &PairNode{Left: left, Right: right}
}
