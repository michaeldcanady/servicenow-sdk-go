package tableapi

// Deprecated: deprecated since v{version}. Use `PageResult2[T]` instead.
//
// PageResult represents a single page of results from a table.
type PageResult = PageResult2[TableEntry]

// PageResult2[T] represents a single page of results from a table.
type PageResult2[T Entry] struct {
	Result           []*T
	NextPageLink     string
	PreviousPageLink string
	FirstPageLink    string
	LastPageLink     string
}
