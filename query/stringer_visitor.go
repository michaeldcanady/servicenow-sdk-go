package query

import (
	"strings"

	ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

var _ ast.NodeVisitor = (*StringerVisitor)(nil)

// StringerVisitor Represents a visitor to convert the tree to a string.
type StringerVisitor struct {
	// builder Builds the returned string.
	builder strings.Builder
}

// VisitPairNode implements ast.NodeVisitor.
func (v *StringerVisitor) VisitPairNode(pair *ast.PairNode) {
	v.Visit(pair.Element1)
	v.builder.WriteString("@")
	v.Visit(pair.Element2)
}

// VisitArrayNode implements ast.NodeVisitor.
func (v *StringerVisitor) VisitArrayNode(array *ast.ArrayNode) {
	for index, element := range array.Elements {
		v.Visit(element)
		if index != len(array.Elements)-1 {
			v.builder.WriteString(",")
		}
	}
}

// NewStringerVisitor Instantiates new StringerVisitor.
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
	v.builder.WriteString(node.Operator.String())
	v.Visit(node.RightExpression)
}

// VisitLiteralNode implements NodeVisitor.
func (v *StringerVisitor) VisitLiteralNode(node *ast.LiteralNode) {
	v.builder.WriteString(node.Value)
}

func (v *StringerVisitor) String() string {
	return v.builder.String()
}
