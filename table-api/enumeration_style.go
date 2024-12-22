package tableapi

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
