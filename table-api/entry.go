package tableapi

func NewTableEntry() TableEntry {
	return TableEntry{}
}

type Entry interface {
	Value(string) *TableValue
	Set(string, interface{})
	Keys() []string
	Len() int
}
