package attachmentapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreateAttachmentCollectionResponseFromDiscriminatorValue is a parsable factory for creating an Collection Request for AttachmentCollectionResponse
func CreateAttachmentCollectionResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*Attachment](CreateAttachmentFromDiscriminatorValue), nil
}

type AttachmentCollectionResponse = core.BaseServiceNowCollectionResponse[*Attachment]
