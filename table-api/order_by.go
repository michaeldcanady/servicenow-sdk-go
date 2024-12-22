package tableapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

// Deprecated: deprecated since 1.4.0. Please use core.OrderBy instead.
//
// OrderBy represents an order-by clause.
type OrderBy = core.OrderBy

// Deprecated: deprecated since 1.4.0. Please use core.NewOrderBy instead.
//
// NewOrderBy Creates new order by.
func NewOrderBy() *OrderBy {
	return core.NewOrderBy()
}
