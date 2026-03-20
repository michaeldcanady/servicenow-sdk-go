package servicenowsdkgo

import (
	"maps"

	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// tableURLTemplate2 the url template for Service-Now batch API
	nowURLTemplate = "{+baseurl}/api/now/"
)

type NowRequestBuilder struct {
	kiota.RequestBuilder
}

func NewNowRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *NowRequestBuilder {
	m := &NowRequestBuilder{
		RequestBuilder: kiota.NewBaseRequestBuilder(requestAdapter, nowURLTemplate, pathParameters),
	}
	return m
}

func NewNowRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *NowRequestBuilder {
	return NewNowRequestBuilderInternal(map[string]string{utils.RawURLKey: rawURL}, requestAdapter)
}

// Table returns a TableRequestBuilder2 associated with the NowRequestBuilder.
func (rB *NowRequestBuilder) Table(tableName string) *tableapi.TableRequestBuilder[*tableapi.TableRecord] {
	pathParameters := maps.Clone(rB.GetPathParameters())

	pathParameters["table"] = tableName
	return tableapi.NewDefaultTableRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
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
func (rB *NowRequestBuilder) Attachment() *attachmentapi.AttachmentRequestBuilder2 {
	return attachmentapi.NewAttachmentRequestBuilder2Internal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Batch returns a BatchRequestBuilder, entrypoint into the batch api.
func (rB *NowRequestBuilder) Batch() *batchapi.BatchRequestBuilder {
	return batchapi.NewBatchRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
