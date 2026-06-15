package servicenowsdkgo

import (
	"maps"

	accountapi "github.com/michaeldcanady/servicenow-sdk-go/v2/account-api"
	actsubapi "github.com/michaeldcanady/servicenow-sdk-go/v2/actsub-api"
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/v2/attachment-api"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/v2/batch-api"
	cmdbinstanceapi "github.com/michaeldcanady/servicenow-sdk-go/v2/cmdb-instance-api"
	documentsapi "github.com/michaeldcanady/servicenow-sdk-go/v2/documents-api"
	internal "github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/v2/table-api"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const nowURLTemplate = "{+baseurl}/api/now"

type NowRequestBuilder struct {
	internal.RequestBuilder
}

func NewServiceNowRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *NowRequestBuilder {
	return &NowRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, nowURLTemplate, pathParameters),
	}
}

func NewServiceNowRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *NowRequestBuilder {
	return NewServiceNowRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

func (rB *NowRequestBuilder) Table(tableName string) *tableapi.TableRequestBuilder[*tableapi.TableRecord] {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["table"] = tableName
	return tableapi.NewDefaultTableRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

func (rB *NowRequestBuilder) Attachment() *attachmentapi.AttachmentRequestBuilder {
	return attachmentapi.NewAttachmentRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

func (rB *NowRequestBuilder) Batch() *batchapi.BatchRequestBuilder {
	return batchapi.NewBatchRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Documents returns a DocumentsRequestBuilder associated with the NowRequestBuilder.
func (rB *NowRequestBuilder) Documents() *documentsapi.DocumentsRequestBuilder {
	return documentsapi.NewDocumentsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Cmdb returns a CmdbRequestBuilder associated with the NowRequestBuilder.
func (rB *NowRequestBuilder) Cmdb() *cmdbinstanceapi.CmdbRequestBuilder {
	return cmdbinstanceapi.NewCmdbRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Account returns an AccountRequestBuilder associated with the NowRequestBuilder.
func (rB *NowRequestBuilder) Account() *accountapi.AccountRequestBuilder {
	return accountapi.NewAccountRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ActSub returns an ActSubRequestBuilder associated with the NowRequestBuilder.
func (rB *NowRequestBuilder) ActSub() *actsubapi.ActSubRequestBuilder {
	return actsubapi.NewActSubRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
