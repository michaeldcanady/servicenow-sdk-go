//go:build preview.query

package query

import "errors"

var (
	UnknownOperatorErr = errors.New("operator is unknown")
)
