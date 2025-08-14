package ast

type UnaryNodeVisitor interface {
	VisitUnaryNode(*UnaryNode)
}
