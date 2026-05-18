package servicenowsdkgo

import (
	"maps"

	accountapi "github.com/michaeldcanady/servicenow-sdk-go/account-api"
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
	cmdbinstanceapi "github.com/michaeldcanady/servicenow-sdk-go/cmdb-instance-api"
	documentsapi "github.com/michaeldcanady/servicenow-sdk-go/documents-api"
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const nowURLTemplate2 = "{+baseurl}/api/now"

type NowRequestBuilder2 struct {
	internal.RequestBuilder
}

func NewServiceNowRequestBuilder3Internal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *NowRequestBuilder2 {
	return &NowRequestBuilder2{
		internal.NewBaseRequestBuilder(requestAdapter, nowURLTemplate2, pathParameters),
	}
}

func NewServiceNowRequestBuilder3(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *NowRequestBuilder2 {
	return NewServiceNowRequestBuilder3Internal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

func (rB *NowRequestBuilder2) Table(tableName string) *tableapi.TableRequestBuilder2[*tableapi.TableRecord] {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["table"] = tableName
	return tableapi.NewDefaultTableRequestBuilder2Internal(pathParameters, rB.GetRequestAdapter())
}

func (rB *NowRequestBuilder2) Attachment() *attachmentapi.AttachmentRequestBuilder2 {
	return attachmentapi.NewAttachmentRequestBuilder2Internal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

func (rB *NowRequestBuilder2) Batch() *batchapi.BatchRequestBuilder {
	return batchapi.NewBatchRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Documents returns a DocumentsRequestBuilder2 associated with the NowRequestBuilder.
func (rB *NowRequestBuilder2) Documents() *documentsapi.DocumentsRequestBuilder2 {
	return documentsapi.NewDocumentsRequestBuilder2Internal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Cmdb returns a CmdbRequestBuilder associated with the NowRequestBuilder.
func (rB *NowRequestBuilder2) Cmdb() *cmdbinstanceapi.CmdbRequestBuilder {
	return cmdbinstanceapi.NewCmdbRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Account returns an AccountRequestBuilder associated with the NowRequestBuilder.
func (rB *NowRequestBuilder2) Account() *accountapi.AccountRequestBuilder {
	return accountapi.NewAccountRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
