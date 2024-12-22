package core

// Deprecated: deprecated since v{unreleased}.
// PageResult represents a single page of results from a table.
type PageResult[T any] struct {
	Result           []*T
	NextPageLink     string
	PreviousPageLink string
	FirstPageLink    string
	LastPageLink     string
}
