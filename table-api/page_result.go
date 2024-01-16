package tableapi

// PageResult is a type that represents a single page of results from a table query.
// It contains a slice of table entries and links to navigate between pages.
type PageResult struct {
	// Result is a slice of pointers to table entries that match the query.
	Result []*TableEntry
	// NextPageLink is the URL of the next page of results, or empty if there is no next page.
	NextPageLink string
	// PreviousPageLink is the URL of the previous page of results, or empty if there is no previous page.
	PreviousPageLink string
	// FirstPageLink is the URL of the first page of results.
	FirstPageLink string
	// LastPageLink is the URL of the last page of results.
	LastPageLink string
}
