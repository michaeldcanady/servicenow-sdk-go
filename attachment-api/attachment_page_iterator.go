package attachmentapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// NewAttachmentPageIterator creates a new AttachmentPageIterator instance.
func NewAttachmentPageIterator(
	res internal.ServiceNowCollectionResponse[Attachment2],
	reqAdapter abstractions.RequestAdapter,
	options ...internal.Option[*internal.PageIterator[Attachment2]],
) (*internal.PageIterator[Attachment2], error) {
	return internal.NewPageIterator[Attachment2](res, reqAdapter, CreateAttachment2FromDiscriminatorValue, options...)
}
