//go:build preview.query

package ast

var _ Node = (*ArrayNode)(nil)

// ArrayNode represents an array of elements.
type ArrayNode struct {
	// Elements elements that make up the array.
	Elements []Node
}

func NewArrayNode(nodes ...Node) *ArrayNode {
	return &ArrayNode{
		Elements: nodes,
	}
}

// Left The leftmost (starting) position of the node in source text.
func (expr *ArrayNode) Left() int {
	return expr.Pos()
}

// Right The rightmost (ending) position of the node in source text.
func (expr *ArrayNode) Right() int {
	if len(expr.Elements) == 0 {
		return -1
	}
	return expr.Elements[len(expr.Elements)-1].Right()
}

// Pos The actual position of the node.
func (expr *ArrayNode) Pos() int {
	if len(expr.Elements) == 0 {
		return -1
	}
	return expr.Elements[0].Left()
}

// Accept Accepts the provided visitor.
func (expr *ArrayNode) Accept(visitor NodeVisitor) {
	visitor.VisitArrayNode(expr)
}
