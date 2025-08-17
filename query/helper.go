//go:build preview.query

package query

import (
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

// convertSliceToArrayNode converts the provided slice of values to an array of literal nodes.
func convertSliceToArrayNode[T Primitive](values ...T) *ast.ArrayNode {
	nodes := make([]ast.Node, len(values))
	for index, value := range values {
		node := ast.LiteralNode{
			Value: fmt.Sprintf("%v", value),
		}
		nodes[index] = &node
	}

	return ast.NewArrayNode(nodes...)
}
