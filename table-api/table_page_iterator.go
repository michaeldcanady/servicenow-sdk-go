package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// TablePageIterator[T] is an iterator over pages of table entries.
type TablePageIterator[T Entry] struct {
	// pageIterator is the core page iterator that this table page iterator wraps.
	pageIterator core.PageIterator[T, *TableCollectionResponse2[T]]
}

// NewTablePageIterator creates a new TablePageIterator instance.
// It takes the current page of results and a client, and returns a pointer to the new TablePageIterator.
// If there is an error while creating the core page iterator, it returns the error.
func NewTablePageIterator[T Entry](currentPage *TableCollectionResponse2[T], client core.Client) (*TablePageIterator[T], error) {
	pageIterator, err := core.NewPageIterator[T, *TableCollectionResponse2[T]](currentPage, client)
	if err != nil {
		return nil, err
	}

	return &TablePageIterator[T]{
		(*pageIterator),
	}, nil
}
