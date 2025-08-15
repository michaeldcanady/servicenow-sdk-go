//go:build preview

package ast

import (
	"fmt"
	"io"
)

type StringerWriter interface {
	fmt.Stringer
	io.StringWriter
}
