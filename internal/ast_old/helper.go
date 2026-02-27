//go:build preview.query

package ast

import (
	"regexp"
	"time"
)

var sysIDRegEx = regexp.MustCompile("^[0-9a-f]{32}$")

// kindOf determines the kind of the value.
func kindOf[T Primitive | time.Time](value T) Kind {
	switch typedValue := interface{}(value).(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return KindNumeric
	case string:
		if sysIDRegEx.MatchString(typedValue) {
			return KindReference
		}
		return KindString
	case time.Time:
		return KindDateTime
	case bool:
		return KindBoolean
	// can't cover because it shouldn't be possible
	default:
		return KindUnknown
	}
}
