// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package ast

// PairNode represents a pair of values, typically used for BETWEEN.
type PairNode struct {
	Left  Node
	Right Node
}

// Accept visits the node, propagating any error the visitor reports. It
// rejects a nil receiver so malformed trees surface as errors rather than
// panics.
func (n *PairNode) Accept(v Visitor) error {
	if n == nil {
		return ErrNilNode
	}
	return v.VisitPair(n)
}

// NewPairNode creates a new PairNode with the given left and right nodes.
func NewPairNode(left, right Node) *PairNode {
	return &PairNode{Left: left, Right: right}
}
