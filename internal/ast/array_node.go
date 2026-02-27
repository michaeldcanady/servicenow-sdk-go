//go:build preview.query

package ast

// ArrayNode represents a list of values, typically used for IN/NOT IN.
type ArrayNode struct {
	Nodes []Node
}

func (n *ArrayNode) Accept(v Visitor) { v.VisitArray(n) }

// NewArrayNode creates a new ArrayNode with the given nodes.
func NewArrayNode(nodes ...Node) *ArrayNode {
	return &ArrayNode{Nodes: nodes}
}
