// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package ast

import (
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
)

type StringerVisitor struct {
	builder strings.Builder
}

var _ Visitor = (*StringerVisitor)(nil)

func NewStringerVisitor() *StringerVisitor {
	return &StringerVisitor{}
}

func (v *StringerVisitor) VisitLiteral(n *LiteralNode) error {
	v.builder.WriteString(n.Value)
	return nil
}

func (v *StringerVisitor) VisitUnary(n *UnaryNode) error {
	if conversion.IsNil(n.Left) {
		return ErrNilChild
	}
	if !n.Op.known() {
		return ErrUnknownOperator
	}
	if err := n.Left.Accept(v); err != nil {
		return err
	}
	v.builder.WriteString(n.Op.String())
	return nil
}

func (v *StringerVisitor) VisitBinary(n *BinaryNode) error {
	if conversion.IsNil(n.Left) {
		return ErrNilChild
	}
	if conversion.IsNil(n.Right) {
		return ErrNilChild
	}
	if !n.Op.known() {
		return ErrUnknownOperator
	}
	if err := n.Left.Accept(v); err != nil {
		return err
	}
	v.builder.WriteString(n.Op.String())
	return n.Right.Accept(v)
}

func (v *StringerVisitor) VisitPair(n *PairNode) error {
	if conversion.IsNil(n.Left) {
		return ErrNilChild
	}
	if conversion.IsNil(n.Right) {
		return ErrNilChild
	}
	if err := n.Left.Accept(v); err != nil {
		return err
	}
	v.builder.WriteString("@")
	return n.Right.Accept(v)
}

func (v *StringerVisitor) VisitArray(n *ArrayNode) error {
	for i, node := range n.Nodes {
		if conversion.IsNil(node) {
			return ErrNilChild
		}
		if i > 0 {
			v.builder.WriteString(",")
		}
		if err := node.Accept(v); err != nil {
			return err
		}
	}
	return nil
}

func (v *StringerVisitor) String() string {
	return v.builder.String()
}
