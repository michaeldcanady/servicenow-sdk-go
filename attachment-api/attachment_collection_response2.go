package attachmentapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreateAttachmentCollectionResponseFromDiscriminatorValue is a parsable factory for creating an Collection Request for AttachmentCollectionResponse
func CreateAttachmentCollectionResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowCollectionResponse[*Attachment](CreateAttachmentFromDiscriminatorValue), nil
}

type AttachmentCollectionResponse = internal.BaseServiceNowCollectionResponse[*Attachment]
