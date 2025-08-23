//go:build preview.query

package query

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
	"github.com/stretchr/testify/assert"
)

func TestConvertSliceToArrayNode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				array := convertSliceToArrayNode("test1")

				assert.Equal(t, ast.NewArrayNode(ast.NewLiteralNode("test1")), array)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestEmpty(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Empty",
			test: func(t *testing.T) {
				assert.True(t, empty(""))
			},
		},
		{
			name: "Whitespace",
			test: func(t *testing.T) {
				assert.True(t, empty(" "))
			},
		},
		{
			name: "Big whitespace",
			test: func(t *testing.T) {
				assert.True(t, empty("                             "))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
