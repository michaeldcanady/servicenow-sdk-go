package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// LogicalOperator ...
//
// Deprecated: deprecated since v1.4.0. Please use core.LogicalOperator instead.
type LogicalOperator = core.LogicalOperator

const (
	// And ...
	//
	// Deprecated: deprecated since v1.4.0. Please use core.And instead.
	And LogicalOperator = core.And
	// Or ...
	//
	// Deprecated: deprecated since v1.4.0. Please use core.Or instead.
	Or LogicalOperator = core.Or
)
