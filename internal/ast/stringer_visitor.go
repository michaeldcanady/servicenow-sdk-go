//go:build preview.query

package ast

import (
	"strings"
)

const (
	arraySeparator string = ","
	pairSeparator  string = "@"
)

var _ NodeVisitor = (*StringerVisitor)(nil)

// StringerVisitor Represents a visitor to convert the tree to a string.
type StringerVisitor struct {
	// builder Builds the returned string.
	builder StringerWriter
}

// NewStringerVisitor Instantiates new StringerVisitor.
func NewStringerVisitor() *StringerVisitor {
	return &StringerVisitor{
		builder: &strings.Builder{},
	}
}

// VisitUnaryNode implements NodeVisitor.
func (v *StringerVisitor) VisitUnaryNode(unary *UnaryNode) {
	v.Visit(unary.Node)
	_, _ = v.builder.WriteString(unary.Op.String())
}

// VisitPairNode implements NodeVisitor.
func (v *StringerVisitor) VisitPairNode(pair *PairNode) {
	v.Visit(pair.Element1)
	_, _ = v.builder.WriteString(pairSeparator)
	v.Visit(pair.Element2)
}

// VisitArrayNode implements NodeVisitor.
func (v *StringerVisitor) VisitArrayNode(array *ArrayNode) {
	for index, element := range array.Elements {
		v.Visit(element)
		if index != len(array.Elements)-1 {
			_, _ = v.builder.WriteString(arraySeparator)
		}
	}
}

// Visit implements NodeVisitor.
func (v *StringerVisitor) Visit(node Node) {
	node.Accept(v)
}

// VisitBinaryNode implements NodeVisitor.
func (v *StringerVisitor) VisitBinaryNode(node *BinaryNode) {
	v.Visit(node.LeftExpression)
	_, _ = v.builder.WriteString(node.Operator.String())
	v.Visit(node.RightExpression)
}

// VisitLiteralNode implements NodeVisitor.
func (v *StringerVisitor) VisitLiteralNode(node *LiteralNode) {
	_, _ = v.builder.WriteString(node.Value)
}

func (v *StringerVisitor) String() string {
	return v.builder.String()
}
