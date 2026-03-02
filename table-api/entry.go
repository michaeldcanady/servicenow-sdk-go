package tableapi

// Deprecated: deprecated since v1.9.0. Please use [model.ServiceNowItem] or [TableRecord]
type Entry interface {
	Value(string) *TableValue
	Set(string, interface{})
	Keys() []string
	Len() int
}
