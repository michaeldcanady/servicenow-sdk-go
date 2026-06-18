package ast

// Node represents a node in the ServiceNow encoded query AST.
type Node interface {
	Accept(Visitor)
}
