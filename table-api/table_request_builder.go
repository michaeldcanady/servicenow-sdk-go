package tableapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

type TableRequestBuilder = TableRequestBuilder2[TableEntry]

// Deprecated: depcrecated since v{version}. Please use `NewTableRequestBuilder2[T]` instead.
// NewTableRequestBuilder creates a new instance of the TableRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableRequestBuilder.
func NewTableRequestBuilder(client core.Client, pathParameters map[string]string) *TableRequestBuilder {

	return NewTableRequestBuilder2[tableEntry](client, pathParameters)
}
