// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package query

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/ast"
)

// BooleanField represents a boolean field in ServiceNow.
type BooleanField struct {
	BaseField
}

// Is query that field is the provided value.
func (f BooleanField) Is(val bool) Condition {
	return f.binary(ast.OperatorIs, val)
}
