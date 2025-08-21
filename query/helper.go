//go:build preview.query

package query

import (
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

// convertSliceToArrayNode converts the provided slice of values to an array of literal nodes.
func convertSliceToArrayNode[T Primitive](values ...T) *ast.ArrayNode {
	nodes := make([]ast.Node, len(values))
	for index, value := range values {
		nodes[index] = ast.NewLiteralNode(value)
	}

	return ast.NewArrayNode(nodes...)
}

func empty(s string) bool {
	return (s == "" || strings.TrimSpace(s) == "")
}
