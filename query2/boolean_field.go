//go:build preview.query

package query2

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast2"
)

// BooleanField represents a boolean field in ServiceNow.
type BooleanField struct {
	BaseField
}

// Is query that field is the provided value.
func (f BooleanField) Is(val bool) Condition {
	return f.binary(ast2.OperatorIs, val)
}
