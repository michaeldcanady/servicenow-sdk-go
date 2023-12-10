package tableapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

// OrderBy represents an order-by clause.
//
// Deprecated: deprecated since {version}. Please use core.OrderBy instead.
type OrderBy = core.OrderBy

// NewOrderBy Creates new order by.
//
// Deprecated: deprecated since {version}. Please use core.NewOrderBy instead.
func NewOrderBy() *OrderBy {
	return core.NewOrderBy()
}
