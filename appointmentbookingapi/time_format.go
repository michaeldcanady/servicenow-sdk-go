package appointmentbookingapi

import (
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
)

const (
	timeFormatUnknown = "unknown"
	timeFormat12Hr    = "12hr"
	timeFormat24Hr    = "24hr"
)

// TimeFormat specifies the UI timeFormat for which to render the data.
type TimeFormat int8

const (
	// TimeFormatUnknown represents an unknown UI timeFormat.
	TimeFormatUnknown TimeFormat = iota - 1
	// TimeFormat12Hr
	TimeFormat12Hr
	// TimeFormat24Hr
	TimeFormat24Hr
)

// ParseTimeFormat resolves the wire representation of a time format to a [TimeFormat].
// Matching is case-insensitive.
func ParseTimeFormat(s string) (interface{}, error) {
	if format, ok := timeFormatValues[strings.ToLower(s)]; ok {
		return format, nil
	}
	return TimeFormatUnknown, unknownEnumValueError("time format", s)
}

var timeFormatStrings = map[TimeFormat]string{
	TimeFormatUnknown: timeFormatUnknown,
	TimeFormat12Hr:    timeFormat12Hr,
	TimeFormat24Hr:    timeFormat24Hr,
}

// timeFormatValues is the lower-cased inverse of [timeFormatStrings], used by [ParseTimeFormat].
var timeFormatValues = invertEnumStrings(timeFormatStrings, TimeFormatUnknown)

// String returns the string representation of the TimeFormat.
func (e TimeFormat) String() string {
	return conversion.EnumString(timeFormatStrings, e, timeFormatUnknown)
}
