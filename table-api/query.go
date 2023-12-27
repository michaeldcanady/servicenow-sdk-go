package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since {version}. Please use core.Query instead.
//
// Query represents a ServiceNow query and its conditions.
type Query = core.Query

// Deprecated: deprecated since {version}. Please use core.NewQuery instead.
//
// NewQuery returns a new Query with no conditions.
func NewQuery() *Query {
	return core.NewQuery()
}
