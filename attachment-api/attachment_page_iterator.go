package attachmentapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// AttachmentPageIterator represents an iterator for paginated attachment collections.
type AttachmentPageIterator struct {
	*newInternal.PageIterator[Attachment2]
}

// NewAttachmentPageIterator creates a new AttachmentPageIterator instance.
func NewAttachmentPageIterator(
	res newInternal.ServiceNowCollectionResponse[Attachment2],
	reqAdapter abstractions.RequestAdapter,
	options ...kiota.Option[*newInternal.PageIterator[Attachment2]],
) (*AttachmentPageIterator, error) {
	iterator, err := newInternal.NewPageIterator[Attachment2](res, reqAdapter, CreateAttachment2FromDiscriminatorValue, options...)
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
func (i *AttachmentPageIterator) Iterate(ctx context.Context, reverse bool, callback func(Attachment2) bool) error {
	return i.PageIterator.Iterate(ctx, reverse, callback)
}

// NextItem returns the next item in the collection, fetching the next page if necessary.
//
// Returns [newInternal.ErrNoMoreItems] if the end of the collection is reached.
func (i *AttachmentPageIterator) NextItem(ctx context.Context) (Attachment2, error) {
	return i.PageIterator.NextItem(ctx)
}

// PreviousItem returns the previous item in the collection, fetching the previous page if necessary.
//
// Returns [newInternal.ErrNoMoreItems] if the beginning of the collection is reached.
func (i *AttachmentPageIterator) PreviousItem(ctx context.Context) (Attachment2, error) {
	return i.PageIterator.PreviousItem(ctx)
}
