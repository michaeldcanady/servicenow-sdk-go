package servicenowsdkgo

import (
	"maps"

	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

const (
	// tableURLTemplate2 the url template for Service-Now batch API
	nowURLTemplate = "{+baseurl}/api/now/"
)

type NowRequestBuilder struct {
	kiota.BaseRequestBuilder
}

// Table returns a TableRequestBuilder2 associated with the NowRequestBuilder.
func (rB *NowRequestBuilder) Table(tableName string) *tableapi.TableRequestBuilder[*tableapi.TableRecord] {
	pathParameters := maps.Clone(rB.PathParameters)

	pathParameters["table"] = tableName
	return tableapi.NewDefaultTableRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Attachment returns an AttachmentRequestBuilder associated with the NowRequestBuilder.
// It allows you to work with attachments and manage attachments in ServiceNow.
func (rB *NowRequestBuilder) Attachment() *attachmentapi.AttachmentRequestBuilder2 {
	return attachmentapi.NewAttachmentRequestBuilder2Internal(maps.Clone(rB.PathParameters), rB.GetRequestAdapter())
}

// Batch returns a BatchRequestBuilder, entrypoint into the batch api.
func (rB *NowRequestBuilder) Batch() *batchapi.BatchRequestBuilder {
	return batchapi.NewBatchRequestBuilderInternal(maps.Clone(rB.PathParameters), rB.GetRequestAdapter())
}
