//go:build preview.query

package query2

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast2"
)

// convertSliceToArrayNode converts the provided slice of values to an array of literal nodes.
func convertSliceToArrayNode[T ast2.Primitive](values ...T) *ast2.ArrayNode {
	nodes := make([]ast2.Node, len(values))
	for index, value := range values {
		nodes[index] = ast2.NewLiteralNode(value)
	}

	return ast2.NewArrayNode(nodes...)
}
