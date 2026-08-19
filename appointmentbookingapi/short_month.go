package appointmentbookingapi

import (
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
)

// TODO: may be worth making a "month" enum, and then having a ConvertToShort/ConvertToLong/etc

const (
	shortMonthUnknown = "unknown"
	shortMonthJan     = "Jan"
	shortMonthFeb     = "Feb"
	shortMonthMar     = "Mar"
	shortMonthApr     = "Apr"
	shortMonthMay     = "May"
	shortMonthJun     = "Jun"
	shortMonthJul     = "Jul"
	shortMonthAug     = "Aug"
	shortMonthSep     = "Sep"
	shortMonthOct     = "Oct"
	shortMonthNov     = "Nov"
	shortMonthDec     = "Dec"
)

// ShortMonth specifies the UI shortMonth for which to render the data.
type ShortMonth int32

const (
	// ShortMonthUnknown represents an unknown UI shortMonth.
	ShortMonthUnknown ShortMonth = iota - 1
	ShortMonthJan
	ShortMonthFeb
	ShortMonthMar
	ShortMonthApr
	ShortMonthMay
	ShortMonthJun
	ShortMonthJul
	ShortMonthAug
	ShortMonthSep
	ShortMonthOct
	ShortMonthNov
	ShortMonthDec
)

// ParseShortMonth resolves the wire representation of a month to a [ShortMonth].
// Matching is case-insensitive; the constants themselves are canonical title-case ("Jan").
func ParseShortMonth(s string) (interface{}, error) {
	if month, ok := shortMonthValues[strings.ToLower(s)]; ok {
		return month, nil
	}
	return ShortMonthUnknown, unknownEnumValueError("short month", s)
}

var shortMonthStrings = map[ShortMonth]string{
	ShortMonthUnknown: shortMonthUnknown,
	ShortMonthJan:     shortMonthJan,
	ShortMonthFeb:     shortMonthFeb,
	ShortMonthMar:     shortMonthMar,
	ShortMonthApr:     shortMonthApr,
	ShortMonthMay:     shortMonthMay,
	ShortMonthJun:     shortMonthJun,
	ShortMonthJul:     shortMonthJul,
	ShortMonthAug:     shortMonthAug,
	ShortMonthSep:     shortMonthSep,
	ShortMonthOct:     shortMonthOct,
	ShortMonthNov:     shortMonthNov,
	ShortMonthDec:     shortMonthDec,
}

// shortMonthValues is the lower-cased inverse of [shortMonthStrings], used by [ParseShortMonth].
var shortMonthValues = invertEnumStrings(shortMonthStrings, ShortMonthUnknown)

// String returns the string representation of the ShortMonth.
func (e ShortMonth) String() string {
	return conversion.EnumString(shortMonthStrings, e, shortMonthUnknown)
}
