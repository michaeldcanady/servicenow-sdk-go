package appointmentbookingapi

import (
	"errors"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
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

func ParseTimeFormat(s string) (interface{}, error) {
	switch strings.ToLower(s) {
	case timeFormat12Hr:
		return TimeFormat12Hr, nil
	case timeFormat24Hr:
		return TimeFormat24Hr, nil
	}
	return TimeFormatUnknown, errors.New("unknown timeFormat")
}

var timeFormatStrings = map[TimeFormat]string{
	TimeFormatUnknown: timeFormatUnknown,
	TimeFormat12Hr:    timeFormat12Hr,
	TimeFormat24Hr:    timeFormat24Hr,
}

// String returns the string representation of the TimeFormat.
func (e TimeFormat) String() string {
	return conversion.EnumString(timeFormatStrings, e, timeFormatUnknown)
}
