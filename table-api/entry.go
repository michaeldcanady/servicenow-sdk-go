package tableapi

type Entry interface {
	Value(string) *TableValue
	Set(string, interface{})
	Keys() []string
	Len() int
}
