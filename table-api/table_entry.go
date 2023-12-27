package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/internal"

func NewTableEntry() TableEntry {
	return TableEntry{}
}

type TableEntry2 interface {
	Value(string) *TableValue
	Set(string, interface{})
	Keys() []string
	Len() int
}

// Deprecated: deprecated since v{version}.
//
// TableEntry represents a single Service-Now Table Entry.
type TableEntry = internal.TableEntry
