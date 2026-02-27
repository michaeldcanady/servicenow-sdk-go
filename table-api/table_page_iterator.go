package tableapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// TablePageIterator represents an iterator for paginated table collections.
type TablePageIterator[T model.ServiceNowItem] struct {
	*newInternal.PageIterator[T]
}

// DefaultTablePageIterator represents a TablePageIterator for the default TableRecord type.
type DefaultTablePageIterator = TablePageIterator[*TableRecord]

// NewTablePageIterator creates a new TablePageIterator instance.
func NewTablePageIterator[T model.ServiceNowItem](
	res newInternal.ServiceNowCollectionResponse[T],
	reqAdapter abstractions.RequestAdapter,
	constructorFunc serialization.ParsableFactory,
	options ...kiota.Option[*newInternal.PageIterator[T]],
) (*TablePageIterator[T], error) {
	iterator, err := newInternal.NewPageIterator[T](res, reqAdapter, constructorFunc, options...)
	if err != nil {
		return nil, err
	}

	return &TablePageIterator[T]{
		PageIterator: iterator,
	}, nil
}

// NewDefaultTablePageIterator creates a new TablePageIterator instance for TableRecord.
func NewDefaultTablePageIterator(
	res newInternal.ServiceNowCollectionResponse[*TableRecord],
	reqAdapter abstractions.RequestAdapter,
	options ...kiota.Option[*newInternal.PageIterator[*TableRecord]],
) (*DefaultTablePageIterator, error) {
	return NewTablePageIterator(res, reqAdapter, CreateTableRecordFromDiscriminatorValue, options...)
}

// Iterate traverses the pages and invokes the callback for each item.
//
// reverse determines the direction of page traversal.
// callback should return true to continue iteration, or false to stop.
func (i *TablePageIterator[T]) Iterate(ctx context.Context, reverse bool, callback func(T) bool) error {
	return i.PageIterator.Iterate(ctx, reverse, callback)
}

// NextItem returns the next item in the collection, fetching the next page if necessary.
//
// Returns [newInternal.ErrNoMoreItems] if the end of the collection is reached.
func (i *TablePageIterator[T]) NextItem(ctx context.Context) (T, error) {
	return i.PageIterator.NextItem(ctx)
}

// PreviousItem returns the previous item in the collection, fetching the previous page if necessary.
//
// Returns [newInternal.ErrNoMoreItems] if the beginning of the collection is reached.
func (i *TablePageIterator[T]) PreviousItem(ctx context.Context) (T, error) {
	return i.PageIterator.PreviousItem(ctx)
}
