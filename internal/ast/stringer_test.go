package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStringerVisitor(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Create visitor",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := NewStringerVisitor()
			assert.NotNil(t, v)
		})
	}
}

func TestStringerVisitor_Visit(t *testing.T) {
	tests := []struct {
		name     string
		visit    func(v Visitor) error
		expected string
	}{
		{
			name: "VisitLiteral",
			visit: func(v Visitor) error {
				return v.VisitLiteral(NewLiteralNode("v"))
			},
			expected: "v",
		},
		{
			name: "VisitUnary",
			visit: func(v Visitor) error {
				return v.VisitUnary(NewUnaryNode(OperatorIsEmpty, NewLiteralNode("f")))
			},
			expected: "fISEMPTY",
		},
		{
			name: "VisitBinary",
			visit: func(v Visitor) error {
				return v.VisitBinary(NewBinaryNode(NewLiteralNode("f"), OperatorIs, NewLiteralNode("v")))
			},
			expected: "f=v",
		},
		{
			name: "VisitPair",
			visit: func(v Visitor) error {
				return v.VisitPair(NewPairNode(NewLiteralNode("a"), NewLiteralNode("b")))
			},
			expected: "a@b",
		},
		{
			name: "VisitArray",
			visit: func(v Visitor) error {
				return v.VisitArray(NewArrayNode(NewLiteralNode("a"), NewLiteralNode("b")))
			},
			expected: "a,b",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := NewStringerVisitor()
			require.NoError(t, test.visit(v))
			assert.Equal(t, test.expected, v.String())
		})
	}
}

// TestStringerVisitor_VisitErrors asserts that structural faults — nil
// children (including typed-nil pointers), the unknown-operator sentinel, and
// out-of-range operators — surface as sentinel-wrapped errors instead of
// panics or silent corruption of the rendered query.
func TestStringerVisitor_VisitErrors(t *testing.T) {
	tests := []struct {
		name        string
		visit       func(v Visitor) error
		expectedErr error
	}{
		{
			name: "UnaryNilLeft",
			visit: func(v Visitor) error {
				return v.VisitUnary(&UnaryNode{Op: OperatorIsEmpty})
			},
			expectedErr: ErrNilChild,
		},
		{
			name: "UnaryTypedNilLeft",
			visit: func(v Visitor) error {
				return v.VisitUnary(&UnaryNode{Op: OperatorIsEmpty, Left: (*LiteralNode)(nil)})
			},
			expectedErr: ErrNilChild,
		},
		{
			name: "UnaryUnknownOperator",
			visit: func(v Visitor) error {
				return v.VisitUnary(&UnaryNode{Op: OperatorUnknown, Left: NewLiteralNode("f")})
			},
			expectedErr: ErrUnknownOperator,
		},
		{
			name: "UnaryOutOfRangeOperator",
			visit: func(v Visitor) error {
				return v.VisitUnary(&UnaryNode{Op: Operator(999), Left: NewLiteralNode("f")})
			},
			expectedErr: ErrUnknownOperator,
		},
		{
			name: "BinaryNilLeft",
			visit: func(v Visitor) error {
				return v.VisitBinary(&BinaryNode{Op: OperatorIs, Right: NewLiteralNode("v")})
			},
			expectedErr: ErrNilChild,
		},
		{
			name: "BinaryNilRight",
			visit: func(v Visitor) error {
				return v.VisitBinary(&BinaryNode{Op: OperatorIs, Left: NewLiteralNode("f")})
			},
			expectedErr: ErrNilChild,
		},
		{
			name: "BinaryUnknownOperator",
			visit: func(v Visitor) error {
				return v.VisitBinary(&BinaryNode{Op: OperatorUnknown, Left: NewLiteralNode("f"), Right: NewLiteralNode("v")})
			},
			expectedErr: ErrUnknownOperator,
		},
		{
			name: "PairNilLeft",
			visit: func(v Visitor) error {
				return v.VisitPair(&PairNode{Right: NewLiteralNode("b")})
			},
			expectedErr: ErrNilChild,
		},
		{
			name: "PairNilRight",
			visit: func(v Visitor) error {
				return v.VisitPair(&PairNode{Left: NewLiteralNode("a")})
			},
			expectedErr: ErrNilChild,
		},
		{
			name: "ArrayNilElement",
			visit: func(v Visitor) error {
				return v.VisitArray(NewArrayNode(NewLiteralNode("a"), nil))
			},
			expectedErr: ErrNilChild,
		},
		{
			name: "NestedFaultPropagates",
			visit: func(v Visitor) error {
				return v.VisitBinary(&BinaryNode{
					Left:  NewLiteralNode("f"),
					Op:    OperatorIs,
					Right: &UnaryNode{Op: OperatorIsEmpty}, // nil left child
				})
			},
			expectedErr: ErrNilChild,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := NewStringerVisitor()
			err := test.visit(v)
			assert.ErrorIs(t, err, test.expectedErr)
		})
	}
}

// TestNode_AcceptNilReceiver asserts that every node's Accept rejects a nil
// receiver with ErrNilNode instead of panicking, so malformed trees fail loud
// through the error channel rather than taking down the caller.
func TestNode_AcceptNilReceiver(t *testing.T) {
	tests := []struct {
		name   string
		accept func(v Visitor) error
	}{
		{"Literal", func(v Visitor) error { var n *LiteralNode; return n.Accept(v) }},
		{"Unary", func(v Visitor) error { var n *UnaryNode; return n.Accept(v) }},
		{"Binary", func(v Visitor) error { var n *BinaryNode; return n.Accept(v) }},
		{"Pair", func(v Visitor) error { var n *PairNode; return n.Accept(v) }},
		{"Array", func(v Visitor) error { var n *ArrayNode; return n.Accept(v) }},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var err error
			assert.NotPanics(t, func() {
				err = test.accept(NewStringerVisitor())
			})
			assert.ErrorIs(t, err, ErrNilNode)
		})
	}
}
