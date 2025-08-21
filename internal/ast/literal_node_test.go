//go:build preview.query

package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLiteralNode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "String",
			test: func(t *testing.T) {
				node := NewLiteralNode("value")

				assert.Equal(t, "value", node.Value)
				assert.Equal(t, KindString, node.Kind)
				assert.Equal(t, -1, node.Position)
			},
		},
		{
			name: "Reference",
			test: func(t *testing.T) {
				node := NewLiteralNode("fec4d97893d32210806af40bdd03d66f")

				assert.Equal(t, "fec4d97893d32210806af40bdd03d66f", node.Value)
				assert.Equal(t, KindReference, node.Kind)
				assert.Equal(t, -1, node.Position)
			},
		},
		{
			name: "bool",
			test: func(t *testing.T) {
				node := NewLiteralNode(true)

				assert.Equal(t, "true", node.Value)
				assert.Equal(t, KindBoolean, node.Kind)
				assert.Equal(t, -1, node.Position)
			},
		},
		{
			name: "int",
			test: func(t *testing.T) {
				node := NewLiteralNode(int(8))

				assert.Equal(t, "8", node.Value)
				assert.Equal(t, KindNumeric, node.Kind)
				assert.Equal(t, -1, node.Position)
			},
		},
		{
			name: "int32",
			test: func(t *testing.T) {
				node := NewLiteralNode(int32(8))

				assert.Equal(t, "8", node.Value)
				assert.Equal(t, KindNumeric, node.Kind)
				assert.Equal(t, -1, node.Position)
			},
		},
		{
			name: "int64",
			test: func(t *testing.T) {
				node := NewLiteralNode(int64(8))

				assert.Equal(t, "8", node.Value)
				assert.Equal(t, KindNumeric, node.Kind)
				assert.Equal(t, -1, node.Position)
			},
		},
		{
			name: "uint",
			test: func(t *testing.T) {
				node := NewLiteralNode(uint(8))

				assert.Equal(t, "8", node.Value)
				assert.Equal(t, KindNumeric, node.Kind)
				assert.Equal(t, -1, node.Position)
			},
		},
		{
			name: "uint8",
			test: func(t *testing.T) {
				node := NewLiteralNode(uint8(8))

				assert.Equal(t, "8", node.Value)
				assert.Equal(t, KindNumeric, node.Kind)
				assert.Equal(t, -1, node.Position)
			},
		},
		{
			name: "uint16",
			test: func(t *testing.T) {
				node := NewLiteralNode(uint16(8))

				assert.Equal(t, "8", node.Value)
				assert.Equal(t, KindNumeric, node.Kind)
				assert.Equal(t, -1, node.Position)
			},
		},
		{
			name: "uint32",
			test: func(t *testing.T) {
				node := NewLiteralNode(uint32(8))

				assert.Equal(t, "8", node.Value)
				assert.Equal(t, KindNumeric, node.Kind)
				assert.Equal(t, -1, node.Position)
			},
		},
		{
			name: "uint64",
			test: func(t *testing.T) {
				node := NewLiteralNode(uint64(8))

				assert.Equal(t, "8", node.Value)
				assert.Equal(t, KindNumeric, node.Kind)
				assert.Equal(t, -1, node.Position)
			},
		},
		{
			name: "float32",
			test: func(t *testing.T) {
				node := NewLiteralNode(float32(8))

				assert.Equal(t, "8", node.Value)
				assert.Equal(t, KindNumeric, node.Kind)
				assert.Equal(t, -1, node.Position)
			},
		},
		{
			name: "float64",
			test: func(t *testing.T) {
				node := NewLiteralNode(float64(8))

				assert.Equal(t, "8", node.Value)
				assert.Equal(t, KindNumeric, node.Kind)
				assert.Equal(t, -1, node.Position)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestLiteralNode_Left(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &LiteralNode{
					Position: 1,
				}

				assert.Equal(t, 1, node.Left())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestLiteralNode_Right(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := "I am a value"
				node := &LiteralNode{
					Position: 1,
					Value:    value,
				}

				assert.Equal(t, 1+len(value), node.Right())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestLiteralNode_Pos(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &LiteralNode{
					Position: 1,
				}

				assert.Equal(t, 1, node.Pos())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestLiteralNode_Accept(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &LiteralNode{}
				visitor := newMockNodeVisitor()
				visitor.On("VisitLiteralNode", node)

				node.Accept(visitor)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
