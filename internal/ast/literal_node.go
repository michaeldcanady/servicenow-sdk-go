// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package ast

import (
	"fmt"
)

// LiteralNode represents a literal value (e.g., "hello", 1, true).
type LiteralNode struct {
	Value string
}

// Accept visits the node, propagating any error the visitor reports. It
// rejects a nil receiver so malformed trees surface as errors rather than
// panics.
func (n *LiteralNode) Accept(v Visitor) error {
	if n == nil {
		return ErrNilNode
	}
	return v.VisitLiteral(n)
}

// NewLiteralNode creates a new LiteralNode for the given value.
func NewLiteralNode(val any) *LiteralNode {
	return &LiteralNode{
		Value: fmt.Sprintf("%v", val),
	}
}
