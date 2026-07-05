package tableapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// NewTablePageIterator creates a new TablePageIterator instance.
func NewTablePageIterator[T model.ServiceNowItem](
	res core.ServiceNowCollectionResponse[T],
	reqAdapter abstractions.RequestAdapter,
	constructorFunc serialization.ParsableFactory,
	options ...internal.Option[*core.PageIterator[T]],
) (*core.PageIterator[T], error) {
	return core.NewPageIterator(res, reqAdapter, constructorFunc, options...)
}

// NewDefaultTablePageIterator creates a new TablePageIterator instance for TableRecord.
func NewDefaultTablePageIterator(
	res core.ServiceNowCollectionResponse[*TableRecord],
	reqAdapter abstractions.RequestAdapter,
	options ...internal.Option[*core.PageIterator[*TableRecord]],
) (*core.PageIterator[*TableRecord], error) {
	return NewTablePageIterator(res, reqAdapter, CreateTableRecordFromDiscriminatorValue, options...)
}
