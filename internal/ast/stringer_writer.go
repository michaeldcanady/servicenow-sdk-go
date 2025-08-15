//go:build preview.query

package ast

import (
	"fmt"
	"io"
)

type StringerWriter interface {
	fmt.Stringer
	io.StringWriter
}
