package appointmentbookingapi

import (
	"errors"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
)

// TODO: may be worth making a "month" enum, and then having a ConvertToShort/ConvertToLong/etc

const (
	shortMonthUnknown = "unknown"
	shortMonthJan     = "jan"
	shortMonthFeb     = "Feb"
	shortMonthMar     = "Mar"
	shortMonthApr     = "Apr"
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
	ShortMonthJun
	ShortMonthJul
	ShortMonthAug
	ShortMonthSep
	ShortMonthOct
	ShortMonthNov
	ShortMonthDec
)

func ParseShortMonth(s string) (interface{}, error) {
	switch strings.ToLower(s) {
	case shortMonthJan:
		return ShortMonthJan, nil
	case shortMonthFeb:
		return ShortMonthFeb, nil
	case shortMonthMar:
		return ShortMonthMar, nil
	case shortMonthApr:
		return ShortMonthApr, nil
	case shortMonthJun:
		return ShortMonthJun, nil
	case shortMonthJul:
		return ShortMonthJul, nil
	case shortMonthAug:
		return ShortMonthAug, nil
	case shortMonthSep:
		return ShortMonthSep, nil
	case shortMonthOct:
		return ShortMonthOct, nil
	case shortMonthNov:
		return ShortMonthNov, nil
	case shortMonthDec:
		return ShortMonthDec, nil
	}
	return ShortMonthUnknown, errors.New("unknown short month")
}

var shortMonthStrings = map[ShortMonth]string{
	ShortMonthUnknown: shortMonthUnknown,
	ShortMonthJan:     shortMonthJan,
	ShortMonthFeb:     shortMonthFeb,
	ShortMonthMar:     shortMonthMar,
	ShortMonthApr:     shortMonthApr,
	ShortMonthJun:     shortMonthJun,
	ShortMonthJul:     shortMonthJul,
	ShortMonthAug:     shortMonthAug,
	ShortMonthSep:     shortMonthSep,
	ShortMonthOct:     shortMonthOct,
	ShortMonthNov:     shortMonthNov,
	ShortMonthDec:     shortMonthDec,
}

// String returns the string representation of the ShortMonth.
func (e ShortMonth) String() string {
	return conversion.EnumString(shortMonthStrings, e, shortMonthUnknown)
}
