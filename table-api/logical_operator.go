package tableapi

import "github.com/RecoLabs/servicenow-sdk-go/core"

// LogicalOperator ...
//
// Deprecated: deprecated since 1.4.0. Please use core.LogicalOperator instead.
type LogicalOperator = core.LogicalOperator

const (
	// And ...
	//
	// Deprecated: deprecated since 1.4.0. Please use core.And instead.
	And LogicalOperator = core.And
	// Or ...
	//
	// Deprecated: deprecated since 1.4.0. Please use core.Or instead.
	Or LogicalOperator = core.Or
)
