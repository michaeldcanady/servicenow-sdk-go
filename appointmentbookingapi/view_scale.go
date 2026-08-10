package appointmentbookingapi

import (
	"errors"
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

func ParseViewScale(s string) (interface{}, error) {
	switch strings.ToLower(s) {
	case viewScaleWeek:
		return ViewScaleWeek, nil
	case viewScaleDay:
		return ViewScaleDay, nil
	}
	return ViewScaleUnknown, errors.New("unknown viewScale")
}

var viewScaleStrings = map[ViewScale]string{
	ViewScaleUnknown: viewScaleUnknown,
	ViewScaleDay:     viewScaleDay,
	ViewScaleWeek:    viewScaleWeek,
}

// String returns the string representation of the ViewScale.
func (e ViewScale) String() string {
	return conversion.EnumString(viewScaleStrings, e, viewScaleUnknown)
}
