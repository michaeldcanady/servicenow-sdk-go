// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package ast

// UnaryNode represents a unary operation (e.g., fieldISEMPTY).
type UnaryNode struct {
	Op   Operator
	Left Node
}

// Accept visits the node, propagating any error the visitor reports. It
// rejects a nil receiver so malformed trees surface as errors rather than
// panics.
func (n *UnaryNode) Accept(v Visitor) error {
	if n == nil {
		return ErrNilNode
	}
	return v.VisitUnary(n)
}

// NewUnaryNode creates a new UnaryNode with the given operator and left node.
func NewUnaryNode(op Operator, left Node) *UnaryNode {
	return &UnaryNode{Op: op, Left: left}
}
