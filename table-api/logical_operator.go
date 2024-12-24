package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since 1.4.0. Please use core.LogicalOperator instead.
//
// LogicalOperator ...
type LogicalOperator = core.LogicalOperator

const (
	// Deprecated: deprecated since 1.4.0. Please use core.And instead.
	//
	// And ...
	And LogicalOperator = core.And
	// Deprecated: deprecated since 1.4.0. Please use core.Or instead.
	//
	// Or ...
	Or LogicalOperator = core.Or
)
