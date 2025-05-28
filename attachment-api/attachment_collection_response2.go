package attachmentapi

import (
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type AttachmentCollectionResponse2 = newInternal.ServiceNowCollectionResponse[Attachment2]

// CreateAttachment2FromDiscriminatorValue is a parsable factory for creating an Attachment2Model
func CreateAttachmentCollectionResponse2FromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowCollectionResponse[Attachment2](CreateAttachment2FromDiscriminatorValue), nil
}

type AttachmentCollectionResponse2Model = newInternal.BaseServiceNowCollectionResponse[Attachment2]
