package servicenowsdkgo

import (
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	policyapi "github.com/michaeldcanady/servicenow-sdk-go/policy-api"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func NewNowRequestBuilder2(url string, client core.Client2) *NowRequestBuilder {
	pathParameters := map[string]string{"baseurl": url}
	requestBuilder := core.NewRequestBuilder2(client, nowURLTemplate, pathParameters)
	return &NowRequestBuilder{
		*requestBuilder,
	}
}

// Policy returns a PolicyRequestBuilder associated with the NowRequestBuilder.
func (rB *NowRequestBuilder) Policy() *policyapi.PolicyRequestBuilder {
	return policyapi.NewPolicyRequestBuilderInternal(rB.PathParameters, rB.Client.(*ServiceNowClient).RequestAdapter)
}

// Deprecated: deprecated since v1.9.0. Please use [NowRequestBuilder.TableV2]
func (rB *NowRequestBuilder) Table2(tableName string) *tableapi.TableRequestBuilder {
	rB.PathParameters["table"] = tableName
	return tableapi.New2TableRequestBuilder(rB.Client2, rB.PathParameters)
}

// TableV2 returns a TableRequestBuilder2 associated with the NowRequestBuilder.
func (rB *NowRequestBuilder) TableV2(tableName string) *tableapi.TableRequestBuilder2[*tableapi.TableRecord] {
	pathParameters := make(map[string]string)
	for k, v := range rB.PathParameters {
		pathParameters[k] = v
	}
	pathParameters["table"] = tableName
	return tableapi.NewDefaultTableRequestBuilder2Internal(pathParameters, rB.Client.(*ServiceNowClient).RequestAdapter)
}

// Attachment returns an AttachmentRequestBuilder associated with the NowRequestBuilder.
// It allows you to work with attachments and manage attachments in ServiceNow.
func (rB *NowRequestBuilder) Attachment2() *attachmentapi.AttachmentRequestBuilder2 {
	return attachmentapi.NewAttachmentRequestBuilder2Internal(rB.PathParameters, rB.Client.(*ServiceNowClient).RequestAdapter)
}

// Batch returns a BatchRequestBuilder, entrypoint into the batch api.
func (rB *NowRequestBuilder) Batch() *batchapi.BatchRequestBuilder {
	return batchapi.NewBatchRequestBuilderInternal(rB.PathParameters, rB.Client.(*ServiceNowClient).RequestAdapter)
}
