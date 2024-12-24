package tableapi

// Deprecated: deprecated since v{unreleased}.
//
// Entry ...
type Entry interface {
	Value(string) *TableValue
	Set(string, interface{})
	Keys() []string
	Len() int
}
