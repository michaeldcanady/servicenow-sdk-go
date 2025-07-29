package query

import (
	"strings"

	ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

var _ ast.NodeVisitor[ast.Node] = (*StringerVisitor)(nil)

type StringerVisitor struct {
	builder strings.Builder
}

func NewStringerVisitor() *StringerVisitor {
	return &StringerVisitor{
		builder: strings.Builder{},
	}
}

// Visit implements NodeVisitor.
func (v *StringerVisitor) Visit(node ast.Node) {
	node.Accept(v)
}

// VisitBinaryNode implements NodeVisitor.
func (v *StringerVisitor) VisitBinaryNode(node *ast.BinaryNode) {
	v.Visit(node.LeftExpression)
	v.builder.WriteString(string(node.Operator))
	v.Visit(node.RightExpression)
}

// VisitLiteralNode implements NodeVisitor.
func (v *StringerVisitor) VisitLiteralNode(node *ast.LiteralNode) {
	v.builder.WriteString(node.Value)
}

func (v *StringerVisitor) String() string {
	return v.builder.String()
}
