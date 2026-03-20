//go:build preview.query

package ast2

// Node represents a node in the ServiceNow encoded query AST.
type Node interface {
	Accept(Visitor)
}
