package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// TablePageIterator[T] represents an iterator object that can be used to get subsequent pages of a table collection.
type TablePageIterator[T Entry] struct {
	core.PageIterator[T, TableCollectionResponse3[T]]
}

// NewTablePageIterator[T] creates a new page iterator.
func NewTablePageIterator[T Entry](currentPage TableCollectionResponse3[T], client core.Client) (*TablePageIterator[T], error) {
	var pageIterator core.PageIterator[T, TableCollectionResponse3[T]]

	pageIterator, err := core.NewPageIterator[T, TableCollectionResponse3[T]](currentPage, client)
	if err != nil {
		return nil, err
	}
	pageIterator = core.NewForwardPageIterator(pageIterator)
	pageIterator = core.NewReversePageIterator(pageIterator)

	return &TablePageIterator[T]{
		pageIterator,
	}, nil
}
