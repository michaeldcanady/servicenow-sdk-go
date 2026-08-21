package ast

// PairNode represents a pair of values, typically used for BETWEEN.
type PairNode struct {
	Left  Node
	Right Node
}

// Accept visits the node. It is a no-op on a nil receiver so that malformed
// trees (e.g., those embedding rejected subtrees) can be traversed safely by
// any visitor.
func (n *PairNode) Accept(v Visitor) {
	if n == nil {
		return
	}
	v.VisitPair(n)
}

// NewPairNode creates a new PairNode with the given left and right nodes.
func NewPairNode(left, right Node) *PairNode {
	return &PairNode{Left: left, Right: right}
}
