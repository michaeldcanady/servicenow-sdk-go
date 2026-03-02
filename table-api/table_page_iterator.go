package tableapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// NewTablePageIterator creates a new TablePageIterator instance.
func NewTablePageIterator[T model.ServiceNowItem](
	res newInternal.ServiceNowCollectionResponse[T],
	reqAdapter abstractions.RequestAdapter,
	constructorFunc serialization.ParsableFactory,
	options ...newInternal.Option[*newInternal.PageIterator[T]],
) (*newInternal.PageIterator[T], error) {
	return newInternal.NewPageIterator[T](res, reqAdapter, constructorFunc, options...)
}

// NewDefaultTablePageIterator creates a new TablePageIterator instance for TableRecord.
func NewDefaultTablePageIterator(
	res newInternal.ServiceNowCollectionResponse[*TableRecord],
	reqAdapter abstractions.RequestAdapter,
	options ...newInternal.Option[*newInternal.PageIterator[*TableRecord]],
) (*newInternal.PageIterator[*TableRecord], error) {
	return NewTablePageIterator(res, reqAdapter, CreateTableRecordFromDiscriminatorValue, options...)
}
