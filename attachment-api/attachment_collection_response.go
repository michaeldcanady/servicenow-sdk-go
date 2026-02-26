package attachmentapi

import (
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type AttachmentCollectionResponse = newInternal.ServiceNowCollectionResponse[Attachment]

type AttachmentCollectionResponseModel = newInternal.BaseServiceNowCollectionResponse[Attachment]

// CreateAttachmentCollectionResponseFromDiscriminatorValue is a parsable factory for creating an Collection Request for Attachment
func CreateAttachmentCollectionResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowCollectionResponse[Attachment](CreateAttachmentFromDiscriminatorValue), nil
}
