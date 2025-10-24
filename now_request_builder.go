package servicenowsdkgo

import (
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func NewNowRequestBuilder2(url string, client core.Client2) *NowRequestBuilder {
	pathParameters := map[string]string{"baseurl": url}
	requestBuilder := core.NewRequestBuilder2(client, nowURLTemplate, pathParameters)
	return &NowRequestBuilder{
		*requestBuilder,
	}
}

func (rB *NowRequestBuilder) Table2(tableName string) *tableapi.TableRequestBuilder {
	rB.PathParameters["table"] = tableName
	return tableapi.NewTableRequestBuilder2(rB.Client2, rB.PathParameters)
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
