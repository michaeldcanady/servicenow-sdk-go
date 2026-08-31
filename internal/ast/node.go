// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package ast

// Node represents a node in the ServiceNow encoded query AST.
type Node interface {
	Accept(Visitor)
}
