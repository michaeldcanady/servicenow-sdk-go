package tableapi

import (
	"github.com/RecoLabs/servicenow-sdk-go/core"
)

// OrderBy represents an order-by clause.
//
// Deprecated: deprecated since 1.4.0. Please use core.OrderBy instead.
type OrderBy = core.OrderBy

// NewOrderBy Creates new order by.
//
// Deprecated: deprecated since 1.4.0. Please use core.NewOrderBy instead.
func NewOrderBy() *OrderBy {
	return core.NewOrderBy()
}
