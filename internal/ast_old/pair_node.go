//go:build preview.query

package ast

var _ Node = (*PairNode)(nil)

// PairNode represents a pair of nodes
type PairNode struct {
	// Element1 represents the first node in the pair.
	Element1 Node
	// Element2 represents the second node in the pair.
	Element2 Node
}

// NewPairNode instantiates a new node pair of the provided nodes.
func NewPairNode(node1, node2 Node) *PairNode {
	return &PairNode{
		Element1: node1,
		Element2: node2,
	}
}

// Accept accepts the NodeVisitor and visits the visitor.
func (p *PairNode) Accept(visitor NodeVisitor) {
	visitor.VisitPairNode(p)
}

// Left returns the left most position of the first element.
func (p *PairNode) Left() int {
	return p.Pos()
}

// Pos returns the left most position of the first element.
func (p *PairNode) Pos() int {
	if p.Element1 == nil {
		return -1
	}

	return p.Element1.Left()
}

// Right returns the right most position of the second element.
func (p *PairNode) Right() int {
	if p.Element2 == nil {
		return -1
	}

	return p.Element2.Right()
}
