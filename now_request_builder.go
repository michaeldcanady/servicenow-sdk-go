package servicenowsdkgo

import (
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

const (
	// tableURLTemplate2 the url template for Service-Now batch API
	tableURLTemplate2 = "{+baseurl}/api/now/v1/table/{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}"
)

type NowRequestBuilder struct {
	newInternal.BaseRequestBuilder
}

func NewNowRequestBuilder(url string, client core.Client2) *NowRequestBuilder {
	pathParameters := map[string]string{"baseurl": url}
	requestBuilder := core.NewRequestBuilder2(client, nowURLTemplate, pathParameters)
	return &NowRequestBuilder{
		*requestBuilder,
	}
}

// Table returns a TableRequestBuilder2 associated with the NowRequestBuilder.
func (rB *NowRequestBuilder) Table(tableName string) *tableapi.TableRequestBuilder2[*tableapi.TableRecord] {
	pathParameters := make(map[string]string)
	for k, v := range rB.PathParameters {
		pathParameters[k] = v
	}
	pathParameters["table"] = tableName
	return tableapi.NewDefaultTableRequestBuilder2Internal(pathParameters, rB.Client.(*ServiceNowClient).RequestAdapter)
}

// Attachment returns an AttachmentRequestBuilder associated with the NowRequestBuilder.
// It allows you to work with attachments and manage attachments in ServiceNow.
func (rB *NowRequestBuilder) Attachment() *attachmentapi.AttachmentRequestBuilder2 {
	return attachmentapi.NewAttachmentRequestBuilder2Internal(rB.PathParameters, rB.Client.(*ServiceNowClient).RequestAdapter)
}

// Batch returns a BatchRequestBuilder, entrypoint into the batch api.
func (rB *NowRequestBuilder) Batch() *batchapi.BatchRequestBuilder {
	return batchapi.NewBatchRequestBuilderInternal(rB.PathParameters, rB.Client.(*ServiceNowClient).RequestAdapter)
}
