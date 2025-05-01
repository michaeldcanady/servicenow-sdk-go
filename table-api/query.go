package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Query represents a ServiceNow query and its conditions.
//
// Deprecated: deprecated since v1.4.0. Please use core.Query instead.
type Query = core.Query

// NewQuery returns a new Query with no conditions.
//
// Deprecated: deprecated since v1.4.0. Please use core.NewQuery instead.
func NewQuery() *Query {
	return core.NewQuery()
}
