package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// TablePageIterator represents an iterator object that can be used to get subsequent pages of a table collection.
type TablePageIterator[T Entry] struct {
	*core.PageIterator[T, TableCollectionResponse3[T]]
}

func NewTablePageIterator[T Entry](currentPage TableCollectionResponse3[T], client core.Client) (*TablePageIterator[T], error) {
	pageIterator, err := core.NewPageIterator[T, TableCollectionResponse3[T]](currentPage, client)
	if err != nil {
		return nil, err
	}

	return &TablePageIterator[T]{
		PageIterator: pageIterator,
	}, nil
}
