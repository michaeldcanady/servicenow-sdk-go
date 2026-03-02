//go:build preview.query

package ast2

import (
	"testing"
)

func TestNewStringerVisitor(t *testing.T) {
	v := NewStringerVisitor()
	if v == nil {
		t.Error("NewStringerVisitor returned nil")
	}
}

func TestStringerVisitor_VisitLiteral(t *testing.T) {
	v := NewStringerVisitor()
	n := NewLiteralNode("v")
	v.VisitLiteral(n)
	if v.String() != "v" {
		t.Errorf("got %s, expected v", v.String())
	}
}

func TestStringerVisitor_VisitUnary(t *testing.T) {
	v := NewStringerVisitor()
	n := NewUnaryNode(OperatorIsEmpty, NewLiteralNode("f"))
	v.VisitUnary(n)
	if v.String() != "fISEMPTY" {
		t.Errorf("got %s, expected fISEMPTY", v.String())
	}
}

func TestStringerVisitor_VisitBinary(t *testing.T) {
	v := NewStringerVisitor()
	n := NewBinaryNode(NewLiteralNode("f"), OperatorIs, NewLiteralNode("v"))
	v.VisitBinary(n)
	if v.String() != "f=v" {
		t.Errorf("got %s, expected f=v", v.String())
	}
}

func TestStringerVisitor_VisitPair(t *testing.T) {
	v := NewStringerVisitor()
	n := NewPairNode(NewLiteralNode("a"), NewLiteralNode("b"))
	v.VisitPair(n)
	if v.String() != "a@b" {
		t.Errorf("got %s, expected a@b", v.String())
	}
}

func TestStringerVisitor_VisitArray(t *testing.T) {
	v := NewStringerVisitor()
	n := NewArrayNode(NewLiteralNode("a"), NewLiteralNode("b"))
	v.VisitArray(n)
	if v.String() != "a,b" {
		t.Errorf("got %s, expected a,b", v.String())
	}
}
