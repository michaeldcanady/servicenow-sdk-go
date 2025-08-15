//go:build preview

package ast

type Kind int64

const (
	KindUnknown Kind = iota - 1
	KindString
	KindReference
	KindDateTime
	KindNumeric
	KindBoolean
)
