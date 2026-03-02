package attachmentapi

import (
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// NewAttachmentPageIterator creates a new AttachmentPageIterator instance.
func NewAttachmentPageIterator(
	res newInternal.ServiceNowCollectionResponse[Attachment2],
	reqAdapter abstractions.RequestAdapter,
	options ...newInternal.Option[*newInternal.PageIterator[Attachment2]],
) (*newInternal.PageIterator[Attachment2], error) {
	return newInternal.NewPageIterator[Attachment2](res, reqAdapter, CreateAttachment2FromDiscriminatorValue, options...)
}
