package tableapi

type pageResult[T any] struct {
	Result           []T
	NextPageLink     string
	PreviousPageLink string
	FirstPageLink    string
	LastPageLink     string
}

// PageResult represents a single page of results from a table.
type PageResult pageResult[*TableEntry]
