package attachmentapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	model "github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// NewAttachmentPageIterator creates a new AttachmentPageIterator instance.
func NewAttachmentPageIterator(
	res model.ServiceNowCollectionResponse[Attachment2],
	reqAdapter abstractions.RequestAdapter,
	options ...kiota.Option[*model.PageIterator[Attachment2]],
) (*model.PageIterator[Attachment2], error) {
	return model.NewPageIterator[Attachment2](res, reqAdapter, CreateAttachment2FromDiscriminatorValue, options...)
}
