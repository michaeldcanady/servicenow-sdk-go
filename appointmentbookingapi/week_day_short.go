package appointmentbookingapi

import (
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
)

// TODO: may be worth making a "month" enum, and then having a ConvertToShort/ConvertToLong/etc

const (
	shortWeekdayUnknown = "unknown"
	shortWeekdayMon     = "Mon"
	shortWeekdayTue     = "Tue"
	shortWeekdayWed     = "Wed"
	shortWeekdayThu     = "Thu"
	shortWeekdayFri     = "Fri"
	shortWeekdaySat     = "Sat"
	shortWeekdaySun     = "Sun"
)

// ShortWeekday specifies the UI shortWeekday for which to render the data.
type ShortWeekday int32

const (
	// ShortWeekdayUnknown represents an unknown UI shortWeekday.
	ShortWeekdayUnknown ShortWeekday = iota - 1
	ShortWeekdayMon
	ShortWeekdayTue
	ShortWeekdayWed
	ShortWeekdayThu
	ShortWeekdayFri
	ShortWeekdaySat
	ShortWeekdaySun
)

// ParseShortWeekday resolves the wire representation of a weekday to a [ShortWeekday].
// Matching is case-insensitive; the constants themselves are canonical title-case ("Mon").
func ParseShortWeekday(s string) (interface{}, error) {
	if weekday, ok := shortWeekdayValues[strings.ToLower(s)]; ok {
		return weekday, nil
	}
	return ShortWeekdayUnknown, unknownEnumValueError("short weekday", s)
}

var shortWeekdayStrings = map[ShortWeekday]string{
	ShortWeekdayUnknown: shortWeekdayUnknown,
	ShortWeekdayMon:     shortWeekdayMon,
	ShortWeekdayTue:     shortWeekdayTue,
	ShortWeekdayWed:     shortWeekdayWed,
	ShortWeekdayThu:     shortWeekdayThu,
	ShortWeekdayFri:     shortWeekdayFri,
	ShortWeekdaySat:     shortWeekdaySat,
	ShortWeekdaySun:     shortWeekdaySun,
}

// shortWeekdayValues is the lower-cased inverse of [shortWeekdayStrings], used by [ParseShortWeekday].
var shortWeekdayValues = invertEnumStrings(shortWeekdayStrings, ShortWeekdayUnknown)

// String returns the string representation of the ShortWeekday.
func (e ShortWeekday) String() string {
	return conversion.EnumString(shortWeekdayStrings, e, shortWeekdayUnknown)
}
