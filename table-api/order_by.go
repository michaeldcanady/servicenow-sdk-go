package tableapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

// OrderBy represents an order-by clause.
//
// Deprecated: deprecated since v1.4.0. Please use core.OrderBy instead.
type OrderBy = core.OrderBy

// NewOrderBy Creates new order by.
//
// Deprecated: deprecated since v1.4.0. Please use core.NewOrderBy instead.
func NewOrderBy() *OrderBy {
	return core.NewOrderBy()
}
