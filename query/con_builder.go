//go:build preview.query

package query

import "time"

type conBuilder[T Primitive | time.Time | any, Q QueryBuilder] interface {
	Is(value T) *Q
	IsAnything() *Q
	IsDynamic(sysID string) *Q
	IsEmpty() *Q
	IsNot(value T) *Q
	IsNotEmpty() *Q
	IsSame(sysID string) *Q
	ConditionBuilder[Q]
}
