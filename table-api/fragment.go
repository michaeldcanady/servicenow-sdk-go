package tableapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

// Deprecated: deprecated since 1.4.0. Please use core.Fragment instead.
//
// Fragment represents a query fragment with a field, operator, and value.
type Fragment = core.Fragment

// Deprecated: deprecated since 1.4.0. Please use core.NewFragment instead.
//
// NewFragment creates a new query fragment with the specified field, operator, and value.
func NewFragment(field string, operator RelationalOperator, value interface{}) *Fragment {
	return core.NewFragment(field, operator, value)
}
