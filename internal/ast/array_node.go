package ast

// ArrayNode represents a list of values, typically used for IN/NOT IN.
type ArrayNode struct {
	Nodes []Node
}

// Accept visits the node. It is a no-op on a nil receiver so that malformed
// trees (e.g., those embedding rejected subtrees) can be traversed safely by
// any visitor.
func (n *ArrayNode) Accept(v Visitor) {
	if n == nil {
		return
	}
	v.VisitArray(n)
}

// NewArrayNode creates a new ArrayNode with the given nodes.
func NewArrayNode(nodes ...Node) *ArrayNode {
	return &ArrayNode{Nodes: nodes}
}
