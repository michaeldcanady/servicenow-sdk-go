//go:build preview.query

package query

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

type conBuilder[T ast.Primitive | time.Time, Q QueryBuilder] interface {
	Is(value T) *Q
	IsAnything() *Q
	IsDynamic(sysID string) *Q
	IsEmpty() *Q
	IsNot(value T) *Q
	IsNotEmpty() *Q
	IsSame(sysID string) *Q
	ConditionBuilder[Q]
}
