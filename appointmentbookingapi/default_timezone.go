package appointmentbookingapi

import (
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
)

const (
	defaultTimeZoneUnknown  = "unknown"
	defaultTimeZoneUser     = "user"
	defaultTimeZoneLocation = "location"
)

// DefaultTimeZone
type DefaultTimeZone int8

const (
	// DefaultTimeZoneUnknown
	DefaultTimeZoneUnknown DefaultTimeZone = iota - 1
	// DefaultTimeZoneUser
	DefaultTimeZoneUser
	// DefaultTimeZoneLocation
	DefaultTimeZoneLocation
)

// ParseDefaultTimeZone resolves the wire representation of a default time zone to a [DefaultTimeZone].
// Matching is case-insensitive.
func ParseDefaultTimeZone(s string) (interface{}, error) {
	if zone, ok := defaultTimeZoneValues[strings.ToLower(s)]; ok {
		return zone, nil
	}
	return DefaultTimeZoneUnknown, unknownEnumValueError("default time zone", s)
}

var defaultTimeZoneStrings = map[DefaultTimeZone]string{
	DefaultTimeZoneUnknown:  defaultTimeZoneUnknown,
	DefaultTimeZoneUser:     defaultTimeZoneUser,
	DefaultTimeZoneLocation: defaultTimeZoneLocation,
}

// defaultTimeZoneValues is the lower-cased inverse of [defaultTimeZoneStrings], used by [ParseDefaultTimeZone].
var defaultTimeZoneValues = invertEnumStrings(defaultTimeZoneStrings, DefaultTimeZoneUnknown)

// String returns the string representation of the DefaultTimeZone.
func (e DefaultTimeZone) String() string {
	return conversion.EnumString(defaultTimeZoneStrings, e, defaultTimeZoneUnknown)
}
