package query

import "fmt"

type Node interface {
	Left() int
	Right() int
	Pos() int
	String() string
}

type BinaryExpression struct {
	LeftExpression  Node
	Operator        Operator
	Position        int
	RightExpression Node
}

func (expr *BinaryExpression) Left() int {
	return expr.LeftExpression.Pos()
}

func (expr *BinaryExpression) Right() int {
	return expr.RightExpression.Pos()
}

func (expr *BinaryExpression) Pos() int {
	return expr.Position
}

func (expr *BinaryExpression) String() string {
	return fmt.Sprintf("%s%s%s", expr.LeftExpression.String(), expr.Operator, expr.RightExpression.String())
}

type Literal struct {
	Position int
	Kind     Kind
	Value    string
}

func (expr *Literal) Left() int {
	return expr.Position
}

func (expr *Literal) Right() int {
	return expr.Position + len(expr.Value)
}

func (expr *Literal) Pos() int {
	return expr.Position
}

func (expr *Literal) String() string {
	return expr.Value
}
