package tableapi

// Deprecated: deprecated since v{unreleased}.
//
// PageResult represents a single page of results from a table.
type PageResult struct {
	Result           []*TableEntry
	NextPageLink     string
	PreviousPageLink string
	FirstPageLink    string
	LastPageLink     string
}
