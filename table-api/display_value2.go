package tableapi

import "strings"

// DisplayValue2 Determines the type of data returned, either the actual values from the database or the display values of the fields.
// Display values are manipulated based on the actual value in the database and user or system settings and preferences.
// If returning display values, the value that is returned is dependent on the field type.
type DisplayValue2 int

const (
	// DisplayValue2Unknown represents an unknown display value type
	DisplayValue2Unknown DisplayValue2 = iota - 1
	// DisplayValue2True represents the display values for all fields
	DisplayValue2True
	// DisplayValue2False represents the actual values from the database
	DisplayValue2False
	// DisplayValue2All represents both actual and display values
	DisplayValue2All
)

// ParseDisplayValue2 converts provided string to a DisplayValue2
func ParseDisplayValue2(str string) DisplayValue2 {
	str = strings.ToLower(str)

	for _, displayValue := range []DisplayValue2{DisplayValue2All, DisplayValue2False, DisplayValue2True, DisplayValue2Unknown} {
		displayValueString := strings.ToLower(displayValue.String())
		if str == displayValueString {
			return displayValue
		}
	}

	return DisplayValue2Unknown
}

// String returns display value as a string
func (dV DisplayValue2) String() string {
	value, ok := map[DisplayValue2]string{
		DisplayValue2Unknown: "unknown",
		DisplayValue2True:    "true",
		DisplayValue2False:   "false",
		DisplayValue2All:     "all",
	}[dV]

	if !ok {
		return DisplayValue2Unknown.String()
	}

	return value
}
