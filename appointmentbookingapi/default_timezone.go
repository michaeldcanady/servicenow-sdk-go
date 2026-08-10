package appointmentbookingapi

import (
	"errors"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
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

func ParseDefaultTimeZone(s string) (interface{}, error) {
	switch strings.ToLower(s) {
	case defaultTimeZoneLocation:
		return DefaultTimeZoneLocation, nil
	case defaultTimeZoneUser:
		return DefaultTimeZoneUser, nil
	}
	return DefaultTimeZoneUnknown, errors.New("unknown view")
}

var defaultTimeZoneStrings = map[DefaultTimeZone]string{
	DefaultTimeZoneUnknown:  viewUnknown,
	DefaultTimeZoneUser:     defaultTimeZoneUser,
	DefaultTimeZoneLocation: defaultTimeZoneLocation,
}

// String returns the string representation of the View.
func (e DefaultTimeZone) String() string {
	return conversion.EnumString(defaultTimeZoneStrings, e, viewUnknown)
}
