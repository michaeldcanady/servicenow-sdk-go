package tableapi

const (
	displayValueUnknown = "unknown"
	displayValueTrue    = "true"
	displayValueFalse   = "false"
	displayValueAll     = "all"
)

// DisplayValue determines the type of data returned, either the actual values from the database or the display values of the fields.
type DisplayValue int

const (
	// DisplayValueUnknown represents an unknown display value setting.
	DisplayValueUnknown DisplayValue = iota - 1
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

// String returns the string representation of the DisplayValue.
func (e DisplayValue) String() string {
	str, ok := displayValueStrings[e]
	if !ok {
		return DisplayValueUnknown.String()
	}
	return str
}
