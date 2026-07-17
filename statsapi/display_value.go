package statsapi

import "github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"

const (
	displayValueUnknown = "unknown"
	displayValueTrue    = "true"
	displayValueFalse   = "false"
	displayValueAll     = "all"
)

// DisplayValue determines whether the Stats API returns display values, actual database values, or both.
type DisplayValue int

const (
	// DisplayValueUnknown is the zero value. It is unset/omitted from the request rather than sent as "unknown".
	DisplayValueUnknown DisplayValue = iota
	// DisplayValueTrue returns the display values for all fields.
	DisplayValueTrue
	// DisplayValueFalse returns the actual values from the database.
	DisplayValueFalse
	// DisplayValueAll returns both display and actual values for all fields.
	DisplayValueAll
)

var displayValueStrings = map[DisplayValue]string{
	DisplayValueUnknown: displayValueUnknown,
	DisplayValueTrue:    displayValueTrue,
	DisplayValueFalse:   displayValueFalse,
	DisplayValueAll:     displayValueAll,
}

// String returns the string representation of the DisplayValue, as sent in sysparm_display_value.
func (e DisplayValue) String() string {
	return conversion.EnumString(displayValueStrings, e, displayValueUnknown)
}
