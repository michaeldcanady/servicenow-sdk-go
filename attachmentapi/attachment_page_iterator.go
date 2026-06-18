package attachmentapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// NewAttachmentPageIterator creates a new AttachmentPageIterator instance.
func NewAttachmentPageIterator(
	res core.ServiceNowCollectionResponse[*Attachment],
	reqAdapter abstractions.RequestAdapter,
	options ...internal.Option[*core.PageIterator[*Attachment]],
) (*core.PageIterator[*Attachment], error) {
	return core.NewPageIterator[*Attachment](res, reqAdapter, CreateAttachmentFromDiscriminatorValue, options...)
}
