package ast

type ArrayNodeVisitor interface {
	VisitArrayNode(*ArrayNode)
}
