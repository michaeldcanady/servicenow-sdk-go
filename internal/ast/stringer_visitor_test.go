package ast

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStringerVisitor(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				visitor := NewStringerVisitor()

				assert.IsType(t, &strings.Builder{}, visitor.builder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestStringerVisitor_VisitUnaryNode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				stringerWriter := newMockStringerWriter()
				stringerWriter.On("WriteString", "unknown").Return(0, nil).Twice()

				visitor := &StringerVisitor{
					builder: stringerWriter,
				}

				node1 := newMockNode()
				node1.On("Accept", visitor)

				node := &UnaryNode{
					Op:   OperatorUnknown,
					Node: node1,
				}

				visitor.VisitUnaryNode(node)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestStringerVisitor_VisitPairNode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				stringerWriter := newMockStringerWriter()
				stringerWriter.On("WriteString", "@").Return(0, nil).Twice()

				visitor := &StringerVisitor{
					builder: stringerWriter,
				}

				node1 := newMockNode()
				node1.On("Accept", visitor)

				node2 := newMockNode()
				node2.On("Accept", visitor)

				node := &PairNode{
					Element1: node1,
					Element2: node2,
				}

				visitor.VisitPairNode(node)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestStringerVisitor_VisitArrayNode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				stringerWriter := newMockStringerWriter()
				stringerWriter.On("WriteString", ",").Return(0, nil).Twice()

				visitor := &StringerVisitor{
					builder: stringerWriter,
				}

				node1 := newMockNode()
				node1.On("Accept", visitor)

				node2 := newMockNode()
				node2.On("Accept", visitor)

				node := &ArrayNode{
					Elements: []Node{node1, node2},
				}

				visitor.VisitArrayNode(node)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestStringerVisitor_Visit(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				stringerWriter := newMockStringerWriter()
				stringerWriter.On("String").Return("example")

				visitor := &StringerVisitor{
					builder: stringerWriter,
				}
				node := newMockNode()
				node.On("Accept", visitor)

				visitor.Visit(node)
				node.AssertExpectations(t)
				stringerWriter.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestStringerVisitor_VisitBinaryNode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				stringerWriter := newMockStringerWriter()
				stringerWriter.On("WriteString", "unknown").Return(0, nil)

				visitor := &StringerVisitor{
					builder: stringerWriter,
				}

				leftNode := newMockNode()
				leftNode.On("Accept", visitor)

				rightNode := newMockNode()
				rightNode.On("Accept", visitor)

				node := &BinaryNode{
					LeftExpression:  leftNode,
					Operator:        OperatorUnknown,
					RightExpression: rightNode,
				}

				visitor.VisitBinaryNode(node)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestStringerVisitor_VisitLiteralNode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				stringerWriter := newMockStringerWriter()
				stringerWriter.On("WriteString", "example1").Return(0, nil)

				visitor := &StringerVisitor{
					builder: stringerWriter,
				}

				node := &LiteralNode{
					Value: "example1",
				}

				visitor.VisitLiteralNode(node)
				stringerWriter.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
