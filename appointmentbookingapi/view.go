package appointmentbookingapi

import (
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
)

const (
	viewUnknown  = "unknown"
	viewPlatform = "platform"
	viewPortal   = "portal"
)

// View specifies the UI view for which to render the data.
type View int32

const (
	// ViewUnknown represents an unknown UI view.
	ViewUnknown View = iota - 1
	// ViewPlatform
	ViewPlatform
	// ViewPortal
	ViewPortal
)

// ParseView resolves the wire representation of a view to a [View].
// Matching is case-insensitive.
func ParseView(s string) (interface{}, error) {
	if view, ok := viewValues[strings.ToLower(s)]; ok {
		return view, nil
	}
	return ViewUnknown, unknownEnumValueError("view", s)
}

var viewStrings = map[View]string{
	ViewUnknown:  viewUnknown,
	ViewPlatform: viewPlatform,
	ViewPortal:   viewPortal,
}

// viewValues is the lower-cased inverse of [viewStrings], used by [ParseView].
var viewValues = invertEnumStrings(viewStrings, ViewUnknown)

// String returns the string representation of the View.
func (e View) String() string {
	return conversion.EnumString(viewStrings, e, viewUnknown)
}
