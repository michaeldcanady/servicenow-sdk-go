package core

// PageResult represents a single page of results from a table.
type PageResult[T any] struct {
	Result []*T
	// NextPageLink the link to the next page.
	NextPageLink string
	// PreviousPageLink the link to the previous page.
	PreviousPageLink string
	// FirstPageLink the link to the first page.
	FirstPageLink string
	// LastPageLink the link to the last page.
	LastPageLink string
}
