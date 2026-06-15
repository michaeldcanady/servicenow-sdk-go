package tableapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// NewTablePageIterator creates a new TablePageIterator instance.
func NewTablePageIterator[T model.ServiceNowItem](
	res internal.ServiceNowCollectionResponse[T],
	reqAdapter abstractions.RequestAdapter,
	constructorFunc serialization.ParsableFactory,
	options ...internal.Option[*internal.PageIterator[T]],
) (*internal.PageIterator[T], error) {
	return internal.NewPageIterator[T](res, reqAdapter, constructorFunc, options...)
}

// NewDefaultTablePageIterator creates a new TablePageIterator instance for TableRecord.
func NewDefaultTablePageIterator(
	res internal.ServiceNowCollectionResponse[*TableRecord],
	reqAdapter abstractions.RequestAdapter,
	options ...internal.Option[*internal.PageIterator[*TableRecord]],
) (*internal.PageIterator[*TableRecord], error) {
	return NewTablePageIterator(res, reqAdapter, CreateTableRecordFromDiscriminatorValue, options...)
}
