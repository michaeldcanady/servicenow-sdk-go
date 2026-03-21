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

const nowURLTemplate2 = "{+baseurl}/api/now"

type NowRequestBuilder struct {
	kiota.RequestBuilder
}

func NewNowRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *NowRequestBuilder {
	return &NowRequestBuilder{
		kiota.NewBaseRequestBuilder(requestAdapter, nowURLTemplate2, pathParameters),
	}
}

func NewServiceNowRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *NowRequestBuilder {
	return NewNowRequestBuilderInternal(map[string]string{utils.RawURLKey: rawURL}, requestAdapter)
}

func (rB *NowRequestBuilder) Table(tableName string) *tableapi.TableRequestBuilder[*tableapi.TableRecord] {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["table"] = tableName
	return tableapi.NewDefaultTableRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

func (rB *NowRequestBuilder) Attachment() *attachmentapi.AttachmentRequestBuilder {
	return attachmentapi.NewAttachmentRequestBuilder2Internal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

func (rB *NowRequestBuilder) Batch() *batchapi.BatchRequestBuilder {
	return batchapi.NewBatchRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
