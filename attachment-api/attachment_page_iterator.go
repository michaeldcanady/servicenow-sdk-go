package attachmentapi

import (
	"context"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttachmentPageIterator represents an iterator for paginated attachment collections.
type AttachmentPageIterator struct {
	*newInternal.PageIterator[Attachment]
}

// NewAttachmentPageIterator creates a new AttachmentPageIterator instance.
func NewAttachmentPageIterator(
	res newInternal.ServiceNowCollectionResponse[Attachment],
	reqAdapter abstractions.RequestAdapter,
	constructorFunc serialization.ParsableFactory,
	options ...newInternal.Option[*newInternal.PageIterator[Attachment]],
) (*AttachmentPageIterator, error) {
	iterator, err := newInternal.NewPageIterator[Attachment](res, reqAdapter, constructorFunc, options...)
	if err != nil {
		return nil, err
	}

	return &AttachmentPageIterator{
		PageIterator: iterator,
	}, nil
}

// Iterate traverses the pages and invokes the callback for each item.
//
// reverse determines the direction of page traversal.
// callback should return true to continue iteration, or false to stop.
func (i *AttachmentPageIterator) Iterate(ctx context.Context, reverse bool, callback func(Attachment) bool) error {
	return i.PageIterator.Iterate(ctx, reverse, callback)
}

// NextItem returns the next item in the collection, fetching the next page if necessary.
//
// Returns [newInternal.ErrNoMoreItems] if the end of the collection is reached.
func (i *AttachmentPageIterator) NextItem(ctx context.Context) (Attachment, error) {
	return i.PageIterator.NextItem(ctx)
}

// PreviousItem returns the previous item in the collection, fetching the previous page if necessary.
//
// Returns [newInternal.ErrNoMoreItems] if the beginning of the collection is reached.
func (i *AttachmentPageIterator) PreviousItem(ctx context.Context) (Attachment, error) {
	return i.PageIterator.PreviousItem(ctx)
}
