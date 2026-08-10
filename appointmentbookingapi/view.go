package appointmentbookingapi

import (
	"errors"
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

func ParseView(s string) (interface{}, error) {
	switch strings.ToLower(s) {
	case viewPortal:
		return ViewPortal, nil
	case viewPlatform:
		return ViewPlatform, nil
	}
	return ViewUnknown, errors.New("unknown view")
}

var viewStrings = map[View]string{
	ViewUnknown:  viewUnknown,
	ViewPlatform: viewPlatform,
	ViewPortal:   viewPortal,
}

// String returns the string representation of the View.
func (e View) String() string {
	return conversion.EnumString(viewStrings, e, viewUnknown)
}
