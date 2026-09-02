// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package ast

// Visitor represents a visitor for the AST.
//
// Every node's Accept takes this single interface: Go requires an exact
// method signature for interface satisfaction, so one shared Accept(Visitor)
// is what lets heterogeneous nodes be traversed uniformly through Node.
//
// Visit methods return an error so traversal can propagate structural faults
// (nil children, unknown operators) instead of panicking or rendering a
// silently-wrong query. Implementations should short-circuit on the first
// error and must not treat partially written output as valid.
type Visitor interface {
	VisitLiteral(node *LiteralNode) error
	VisitUnary(node *UnaryNode) error
	VisitBinary(node *BinaryNode) error
	VisitPair(node *PairNode) error
	VisitArray(node *ArrayNode) error
}
