package appointmentbookingapi

import (
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
)

const (
	viewScaleUnknown = "unknown"
	viewScaleDay     = "day"
	viewScaleWeek    = "week"
)

// ViewScale specifies the UI viewScale for which to render the data.
type ViewScale int8

const (
	// ViewScaleUnknown represents an unknown UI viewScale.
	ViewScaleUnknown ViewScale = iota - 1
	// ViewScaleDay
	ViewScaleDay
	// ViewScaleWeek
	ViewScaleWeek
)

// ParseViewScale resolves the wire representation of a view scale to a [ViewScale].
// Matching is case-insensitive.
func ParseViewScale(s string) (interface{}, error) {
	if scale, ok := viewScaleValues[strings.ToLower(s)]; ok {
		return scale, nil
	}
	return ViewScaleUnknown, unknownEnumValueError("view scale", s)
}

var viewScaleStrings = map[ViewScale]string{
	ViewScaleUnknown: viewScaleUnknown,
	ViewScaleDay:     viewScaleDay,
	ViewScaleWeek:    viewScaleWeek,
}

// viewScaleValues is the lower-cased inverse of [viewScaleStrings], used by [ParseViewScale].
var viewScaleValues = invertEnumStrings(viewScaleStrings, ViewScaleUnknown)

// String returns the string representation of the ViewScale.
func (e ViewScale) String() string {
	return conversion.EnumString(viewScaleStrings, e, viewScaleUnknown)
}
