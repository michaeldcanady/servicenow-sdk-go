package tableapi

import (
	"github.com/RecoLabs/servicenow-sdk-go/core"
)

// Fragment represents a query fragment with a field, operator, and value.
//
// Deprecated: deprecated since 1.4.0. Please use core.Fragment instead.
type Fragment = core.Fragment

// NewFragment creates a new query fragment with the specified field, operator, and value.
//
// Deprecated: deprecated since 1.4.0. Please use core.NewFragment instead.
func NewFragment(field string, operator RelationalOperator, value interface{}) *Fragment {
	return core.NewFragment(field, operator, value)
}
