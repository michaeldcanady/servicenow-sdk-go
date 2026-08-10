package appointmentbookingapi

import (
	"errors"
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
)

func ParseShortWeekday(s string) (interface{}, error) {
	switch strings.ToLower(s) {
	case shortWeekdayMon:
		return shortWeekdayMon, nil
	case shortWeekdayTue:
		return ShortWeekdayTue, nil
	case shortWeekdayWed:
		return ShortWeekdayWed, nil
	case shortWeekdayThu:
		return ShortWeekdayThu, nil
	case shortWeekdayFri:
		return ShortWeekdayFri, nil
	}
	return ShortWeekdayUnknown, errors.New("unknown short month")
}

var shortWeekdayStrings = map[ShortWeekday]string{
	ShortWeekdayUnknown: shortWeekdayUnknown,
	ShortWeekdayMon:     shortWeekdayMon,
	ShortWeekdayTue:     shortWeekdayTue,
	ShortWeekdayWed:     shortWeekdayWed,
	ShortWeekdayThu:     shortWeekdayThu,
	ShortWeekdayFri:     shortWeekdayFri,
}

// String returns the string representation of the ShortWeekday.
func (e ShortWeekday) String() string {
	return conversion.EnumString(shortWeekdayStrings, e, shortWeekdayUnknown)
}
