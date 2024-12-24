package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since 1.4.0. Please use core.OrderDirection instead.
//
// OrderDirection represents the order direction for sorting.
type OrderDirection = core.OrderDirection

const (
	// Deprecated: deprecated since 1.4.0. Please use core.Unset instead.
	//
	// Unset ...
	Unset OrderDirection = core.Unset
	// Deprecated: deprecated since 1.4.0. Please use core.Asc instead.
	//
	// Asc ...
	Asc OrderDirection = core.Asc
	// Deprecated: deprecated since 1.4.0. Please use core.Desc instead.
	//
	// Desc ...
	Desc OrderDirection = core.Desc
)
