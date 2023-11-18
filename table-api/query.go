package tableapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

// Deprecated: utilize core.NewQuery
// NewQuery returns a new Query with no conditions.
func NewQuery() *core.Query {
	return core.NewQuery()
}
