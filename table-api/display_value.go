package tableapi

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

// String returns the string representation of the DisplayValue2.
func (e DisplayValue) String() string {
	str, ok := map[DisplayValue]string{
		DisplayValueUnknown: "unknown",
		DisplayValueTrue:    "true",
		DisplayValueFalse:   "false",
		DisplayValueAll:     "all",
	}[e]
	if !ok {
		return View2Unknown.String()
	}
	return str
}
