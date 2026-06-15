package attachmentapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// NewAttachmentPageIterator creates a new AttachmentPageIterator instance.
func NewAttachmentPageIterator(
	res internal.ServiceNowCollectionResponse[*Attachment],
	reqAdapter abstractions.RequestAdapter,
	options ...internal.Option[*internal.PageIterator[*Attachment]],
) (*internal.PageIterator[*Attachment], error) {
	return internal.NewPageIterator[*Attachment](res, reqAdapter, CreateAttachment2FromDiscriminatorValue, options...)
}
