// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package ast

// ArrayNode represents a list of values, typically used for IN/NOT IN.
type ArrayNode struct {
	Nodes []Node
}

// Accept visits the node, propagating any error the visitor reports. It
// rejects a nil receiver so malformed trees surface as errors rather than
// panics.
func (n *ArrayNode) Accept(v Visitor) error {
	if n == nil {
		return ErrNilNode
	}
	return v.VisitArray(n)
}

// NewArrayNode creates a new ArrayNode with the given nodes.
func NewArrayNode(nodes ...Node) *ArrayNode {
	return &ArrayNode{Nodes: nodes}
}
