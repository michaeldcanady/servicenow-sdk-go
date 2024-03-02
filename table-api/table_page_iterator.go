package tableapi

import (
	"net/http"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

// TablePageIterato2[T] is an iterator over pages of table entries.
type TablePageIterator[T Entry] struct {
	// pageIterator is the core page iterator that this table page iterator wraps.
	pageIterator core.PageIterator2[T]
}

func constructTableCollection[T Entry](response *http.Response) (core.CollectionResponse[T], error) {
	resp := &TableCollectionResponse2[T]{}

	err := internal.ParseResponse(response, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NewTablePageIterator[T] creates a new TablePageIterator2 instance.
func NewTablePageIterator[T Entry](collection *TableCollectionResponse2[T], client core.Client) (*TablePageIterator[T], error) {
	pageIterator, err := core.NewPageIterator2[T](collection, client, constructTableCollection[T])
	if err != nil {
		return nil, err
	}

	return &TablePageIterator[T]{
		pageIterator: (*pageIterator),
	}, nil
}
