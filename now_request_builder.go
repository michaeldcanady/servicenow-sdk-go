package servicenowsdkgo

import (
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

type NowRequestBuilder struct {
	core.RequestBuilder
}

// NewNowRequestBuilder creates a new instance of the NowRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created NowRequestBuilder
func NewNowRequestBuilder(url string, client core.Client) *NowRequestBuilder {
	pathParameters := map[string]string{"baseurl": url}
	requestBuilder := core.NewRequestBuilder(client, "{+baseurl}/Now", pathParameters)
	return &NowRequestBuilder{
		*requestBuilder,
	}
}

// Table returns a TableRequestBuilder associated with the NowRequestBuilder.
// It accepts a table name as a parameter and constructs the URL for table-related requests.
// The returned TableRequestBuilder can be used to build and execute table-related requests.
func (rB *NowRequestBuilder) Table(tableName string) *tableapi.TableRequestBuilder {
	rB.PathParameters["table"] = tableName
	return tableapi.NewTableRequestBuilder(rB.Client, rB.PathParameters)
}

// Deprecated: deprecated since v{unreleased}.
//
// Attachment returns an AttachmentRequestBuilder associated with the NowRequestBuilder.
// It allows you to work with attachments and manage attachments in ServiceNow.
func (rB *NowRequestBuilder) Attachment() *attachmentapi.AttachmentRequestBuilder {
	return attachmentapi.NewAttachmentRequestBuilder(rB.Client, rB.PathParameters)
}

// Attachment returns an AttachmentRequestBuilder associated with the NowRequestBuilder.
// It allows you to work with attachments and manage attachments in ServiceNow.
func (rB *NowRequestBuilder) Attachment2() *attachmentapi.AttachmentRequestBuilder2 {
	return attachmentapi.NewAttachmentRequestBuilder2Internal(rB.PathParameters, rB.Client.(*ServiceNowClient).requestAdapter)
}

// Batch returns a BatchRequestBuilder, entrypoint into the batch api.
func (rB *NowRequestBuilder) Batch() *batchapi.BatchRequestBuilder {
	return batchapi.NewBatchRequestBuilderInternal(rB.PathParameters, rB.Client.(*ServiceNowClient).requestAdapter)
}
