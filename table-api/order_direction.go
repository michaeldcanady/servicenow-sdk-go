package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// OrderDirection represents the order direction for sorting.
//
// Deprecated: deprecated since v1.4.0. Please use core.OrderDirection instead.
type OrderDirection = core.OrderDirection

const (
	// Unset ...
	//
	// Deprecated: deprecated since v1.4.0. Please use core.Unset instead.
	Unset OrderDirection = core.Unset
	// Asc ...
	//
	// Deprecated: deprecated since v1.4.0. Please use core.Asc instead.
	Asc OrderDirection = core.Asc
	// Desc ...
	//
	// Deprecated: deprecated since v1.4.0. Please use core.Desc instead.
	Desc OrderDirection = core.Desc
)
