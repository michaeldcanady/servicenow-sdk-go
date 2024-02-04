package tableapi

import "github.com/RecoLabs/servicenow-sdk-go/core"

// Deprecated: deprecated since 1.4.0. Please use core.RelationalOperator instead.
//
// RelationalOperator ...
type RelationalOperator = core.RelationalOperator

const (
	// Deprecated: deprecated since 1.4.0. Please use core.Null instead.
	//
	// Null ...
	Null RelationalOperator = core.Null
	// Deprecated: deprecated since 1.4.0. Please use core.Is instead.
	//
	// Is ...
	Is RelationalOperator = core.Is
	// Deprecated: deprecated since 1.4.0. Please use core.IsNot instead.
	//
	// IsNot ...
	IsNot RelationalOperator = core.IsNot
	// Deprecated: deprecated since 1.4.0. Please use core.GreaterThan instead.
	//
	// GreaterThan ...
	GreaterThan RelationalOperator = core.GreaterThan
	// Deprecated: deprecated since 1.4.0. Please use core.GreaterOrEqual instead.
	//
	// GreaterOrEqual ...
	GreaterOrEqual RelationalOperator = core.GreaterOrEqual
	// Deprecated: deprecated since 1.4.0. Please use core.LessThan instead.
	//
	// LessThan ...
	LessThan RelationalOperator = core.LessOrEqual
	// Deprecated: deprecated since 1.4.0. Please use core.LessOrEqual instead.
	//
	// LessOrEqual ...
	LessOrEqual RelationalOperator = core.LessOrEqual
	// Deprecated: deprecated since 1.4.0. Please use core.Contains instead.
	//
	// Contains ...
	Contains RelationalOperator = core.Contains
	// Deprecated: deprecated since 1.4.0. Please use core.NotContains instead.
	//
	// NotContains ...
	NotContains RelationalOperator = core.NotContains
	// Deprecated: deprecated since 1.4.0. Please use core.StartsWith instead.
	//
	// StartsWith ...
	StartsWith RelationalOperator = core.StartsWith
	// Deprecated: deprecated since 1.4.0. Please use core.EndsWith instead.
	//
	// EndsWith ...
	EndsWith RelationalOperator = core.EndsWith
	// Deprecated: deprecated since 1.4.0. Please use core.Between instead.
	//
	// Between ...
	Between RelationalOperator = core.Between
	// Deprecated: deprecated since 1.4.0. Please use core.IsSame instead.
	//
	// IsSame ...
	IsSame RelationalOperator = core.IsSame
	// Deprecated: deprecated since 1.4.0. Please use core.IsDifferent instead.
	//
	// IsDifferent ...
	IsDifferent RelationalOperator = core.IsDifferent
	// Deprecated: deprecated since 1.4.0. Please use core.IsEmpty instead.
	//
	// IsEmpty ...
	IsEmpty RelationalOperator = core.IsEmpty
)
