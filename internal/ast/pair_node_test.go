//go:build preview.query

package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairNode_Left(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "No Element1",
			test: func(t *testing.T) {
				node := &PairNode{}

				assert.Equal(t, -1, node.Left())
			},
		},
		{
			name: "Has Element1",
			test: func(t *testing.T) {
				element1 := newMockNode()
				element1.On("Left").Return(1)

				node := &PairNode{
					Element1: element1,
				}

				assert.Equal(t, 1, node.Left())
				element1.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestPairNode_Right(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "No Element2",
			test: func(t *testing.T) {
				node := &PairNode{}

				assert.Equal(t, -1, node.Right())
			},
		},
		{
			name: "Has Element2",
			test: func(t *testing.T) {
				element1 := newMockNode()
				element1.On("Right").Return(1)

				node := &PairNode{
					Element2: element1,
				}

				assert.Equal(t, 1, node.Right())
				element1.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestPairNode_Pos(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "No Element1",
			test: func(t *testing.T) {
				node := &PairNode{}

				assert.Equal(t, -1, node.Pos())
			},
		},
		{
			name: "Has Element1",
			test: func(t *testing.T) {
				element1 := newMockNode()
				element1.On("Left").Return(1)

				node := &PairNode{
					Element1: element1,
				}

				assert.Equal(t, 1, node.Pos())
				element1.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestPairNode_Accept(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &PairNode{}
				visitor := newMockNodeVisitor()
				visitor.On("VisitPairNode", node)

				node.Accept(visitor)
				visitor.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
