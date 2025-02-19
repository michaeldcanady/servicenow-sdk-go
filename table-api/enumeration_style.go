package tableapi

import "strings"

// enumerationStyle represents enumeration styles
type enumerationStyle int64

const (
	// enumerationStyleUnknown
	enumerationStyleUnknown enumerationStyle = iota - 1
	// enumerationStyleAll
	enumerationStyleAll
	// enumerationStyleOnlyChanged
	enumerationStyleOnlyChanged
	// enumerationStyleOnlyChangedToNil
	enumerationStyleOnlyChangedToNil
)

func parseEnumerationStyle(str string) enumerationStyle { //nolint:unused
	str = strings.ToLower(str)

	for _, displayValue := range []enumerationStyle{enumerationStyleUnknown, enumerationStyleAll, enumerationStyleOnlyChanged, enumerationStyleOnlyChangedToNil} {
		displayValueString := strings.ToLower(displayValue.String())
		if str == displayValueString {
			return displayValue
		}
	}

	return enumerationStyleUnknown
}

// String returns string representation
func (eS enumerationStyle) String() string {
	style, ok := map[enumerationStyle]string{
		enumerationStyleUnknown:          "unknown",
		enumerationStyleAll:              "all",
		enumerationStyleOnlyChanged:      "onlychanged",
		enumerationStyleOnlyChangedToNil: "onlychangedtonil",
	}[eS]
	if !ok {
		return enumerationStyleUnknown.String()
	}
	return style
}
