package servicenowsdkgo

import (
	"maps"

	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// nowURLTemplate the url template for Service-Now batch API
	nowURLTemplate = "{+baseurl}/api/now/"
)

type NowRequestBuilder struct {
	newInternal.RequestBuilder
}

func NewNowRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *NowRequestBuilder {
	m := &NowRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, nowURLTemplate, pathParameters),
	}
	return m
}

func NewNowRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *NowRequestBuilder {
	return NewNowRequestBuilderInternal(map[string]string{newInternal.RawURLKey: rawURL}, requestAdapter)
}

// Table returns a TableRequestBuilder associated with the NowRequestBuilder.
func (rB *NowRequestBuilder) Table(tableName string) *tableapi.TableRequestBuilder[*tableapi.TableRecord] {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["table"] = tableName
	return tableapi.NewDefaultTableRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Attachment returns an AttachmentRequestBuilder associated with the NowRequestBuilder.
// It allows you to work with attachments and manage attachments in ServiceNow.
func (rB *NowRequestBuilder) Attachment() *attachmentapi.AttachmentRequestBuilder2 {
	pathParameters := maps.Clone(rB.GetPathParameters())

	return attachmentapi.NewAttachmentRequestBuilder2Internal(pathParameters, rB.GetRequestAdapter())
}

// Batch returns a BatchRequestBuilder, entrypoint into the batch api.
func (rB *NowRequestBuilder) Batch() *batchapi.BatchRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())

	return batchapi.NewBatchRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
