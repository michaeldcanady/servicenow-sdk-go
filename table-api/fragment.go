package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: utilize core.NewFragment
// NewFragment creates a new query fragment with the specified field, operator, and value.
func NewFragment(field string, operator core.RelationalOperator, value interface{}) *core.Fragment {
	return core.NewFragment(field, operator, value)
}
