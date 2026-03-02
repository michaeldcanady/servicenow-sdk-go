package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// OrderDirection represents the order direction for sorting.
//
// Deprecated: deprecated since v1.4.0. Please use core.OrderDirection instead.
type OrderDirection = core.OrderDirection

const (
	// Deprecated: deprecated since v1.4.0. Please use [core.Unset] instead.
	//
	// Unset ...
	Unset OrderDirection = core.Unset

	// Deprecated: deprecated since v1.4.0. Please use [core.Asc] instead.
	//
	// Asc ...
	Asc OrderDirection = core.Asc

	// Deprecated: deprecated since v1.4.0. Please use [core.Desc] instead.
	//
	// Desc ...
	Desc OrderDirection = core.Desc
)
