package tableapi

// DisplayValue2 determines the type of data returned, either the actual values from the database or the display values of the fields.
type DisplayValue2 int

const (
	// DisplayValue2Unknown represents an unknown display value setting.
	DisplayValue2Unknown DisplayValue2 = iota - 1
	// DisplayValue2True returns the display values for all fields.
	DisplayValue2True
	// DisplayValue2False returns the actual values from the database.
	DisplayValue2False
	// DisplayValue2All returns both display and actual values for all fields.
	DisplayValue2All
)

// String returns the string representation of the DisplayValue2.
func (e DisplayValue2) String() string {
	str, ok := map[DisplayValue2]string{
		DisplayValue2Unknown: "unknown",
		DisplayValue2True:    "true",
		DisplayValue2False:   "false",
		DisplayValue2All:     "all",
	}[e]
	if !ok {
		return View2Unknown.String()
	}
	return str
}
