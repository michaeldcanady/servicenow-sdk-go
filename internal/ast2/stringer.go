//go:build preview.query

package ast2

import (
	"strings"
)

type StringerVisitor struct {
	builder strings.Builder
}

func NewStringerVisitor() *StringerVisitor {
	return &StringerVisitor{}
}

func (v *StringerVisitor) VisitLiteral(n *LiteralNode) {
	v.builder.WriteString(n.Value)
}

func (v *StringerVisitor) VisitUnary(n *UnaryNode) {
	n.Left.Accept(v)
	v.builder.WriteString(n.Op.String())
}

func (v *StringerVisitor) VisitBinary(n *BinaryNode) {
	n.Left.Accept(v)
	v.builder.WriteString(n.Op.String())
	n.Right.Accept(v)
}

func (v *StringerVisitor) VisitPair(n *PairNode) {
	n.Left.Accept(v)
	v.builder.WriteString("@")
	n.Right.Accept(v)
}

func (v *StringerVisitor) VisitArray(n *ArrayNode) {
	for i, node := range n.Nodes {
		if i > 0 {
			v.builder.WriteString(",")
		}
		node.Accept(v)
	}
}

func (v *StringerVisitor) String() string {
	return v.builder.String()
}
