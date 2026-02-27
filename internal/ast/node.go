//go:build preview.query

package ast

// Node represents a node in the ServiceNow encoded query AST.
type Node interface {
	Accept(Visitor)
}
