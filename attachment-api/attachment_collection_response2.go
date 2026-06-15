package attachmentapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type AttachmentCollectionResponse2 = internal.ServiceNowCollectionResponse[*Attachment]

// CreateAttachment2FromDiscriminatorValue is a parsable factory for creating an Collection Request for Attachment2Model
func CreateAttachmentCollectionResponse2FromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowCollectionResponse[*Attachment](CreateAttachment2FromDiscriminatorValue), nil
}

type AttachmentCollectionResponse2Model = internal.BaseServiceNowCollectionResponse[*Attachment]
