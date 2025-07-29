package ast

// LiteralNode Represents a literal value node (e.g. numeric constant, string, or keyword) in the expression tree.
type LiteralNode struct {
	// Position The position of the literal in the source text.
	Position int
	// Kind The category or type of literal.
	Kind Kind
	// The literal value itself.
	Value string
}

func (expr *LiteralNode) Left() int {
	return expr.Position
}

func (expr *LiteralNode) Right() int {
	return expr.Position + len(expr.Value)
}

func (expr *LiteralNode) Pos() int {
	return expr.Position
}

func (expr *LiteralNode) String() string {
	return expr.Value
}

func (expr *LiteralNode) Accept(visitor NodeVisitor[Node]) {
	visitor.VisitLiteralNode(expr)
}
