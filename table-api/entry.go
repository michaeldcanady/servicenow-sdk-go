package tableapi

// Deprecated: deprecated since v{unreleased}. Please use [model.ServiceNowItem] or [TableRecord]
type Entry interface {
	Value(string) *TableValue
	Set(string, interface{})
	Keys() []string
	Len() int
}
