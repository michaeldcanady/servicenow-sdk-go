package query

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
	"github.com/stretchr/testify/assert"
)

func TestStringerVisitor_VisitArrayNode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				node := &ast.ArrayNode{
					LeftBrace: 0,
					Elements: []ast.Node{
						&ast.LiteralNode{
							Position: 0,
							Value:    "test",
						},
						&ast.LiteralNode{
							Position: 0,
							Value:    "test",
						},
					},
				}

				visitor := NewStringerVisitor()
				visitor.VisitArrayNode(node)

				assert.Equal(t, "test,test", visitor.String())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
