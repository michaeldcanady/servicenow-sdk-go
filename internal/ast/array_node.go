package ast

type ArrayNode struct {
	LeftBrace  int
	Elements   []Node
	RightBrace int
	Incomplete bool
}

// Left The leftmost (starting) position of the node in source text.
func (expr *ArrayNode) Left() int {
	return expr.Pos()
}

// Right The rightmost (ending) position of the node in source text.
func (expr *ArrayNode) Right() int {
	return expr.RightBrace
}

// Pos The actual position of the node.
func (expr *ArrayNode) Pos() int {
	return expr.LeftBrace
}

// Accept Accepts the provided visitor.
func (expr *ArrayNode) Accept(visitor NodeVisitor) {
	visitor.VisitArrayNode(expr)
}
