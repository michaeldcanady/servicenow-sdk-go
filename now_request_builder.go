package servicenowsdkgo

import (
	"maps"

	accountapi "github.com/michaeldcanady/servicenow-sdk-go/v2/accountapi"
	actsubapi "github.com/michaeldcanady/servicenow-sdk-go/v2/activitysubscriptionsapi"
	aggregationapi "github.com/michaeldcanady/servicenow-sdk-go/v2/aggregationapi"
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/v2/attachmentapi"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/v2/batchapi"
	cmdbinstanceapi "github.com/michaeldcanady/servicenow-sdk-go/v2/cmdbinstanceapi"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	documentsapi "github.com/michaeldcanady/servicenow-sdk-go/v2/documentsapi"
	internal "github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/v2/tableapi"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	nowURLTemplate = "{+baseurl}/api/now"
	tablePathKey   = "table"
)

// NowRequestBuilder provides operations to manage the "now" API namespace.
type NowRequestBuilder struct {
	core.RequestBuilder
}

// NewServiceNowRequestBuilderInternal instantiates a new [NowRequestBuilder] from raw path
// parameters. It is exported so sibling packages can construct a NowRequestBuilder while
// chaining through this SDK's fluent builder tree; consumers should generally use
// [NewServiceNowRequestBuilder] instead.
func NewServiceNowRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *NowRequestBuilder {
	return &NowRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, nowURLTemplate, pathParameters),
	}
}

// NewServiceNowRequestBuilder instantiates a new [NowRequestBuilder] with the provided base URL
// and request adapter.
func NewServiceNowRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *NowRequestBuilder {
	return NewServiceNowRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Table returns a [tableapi.TableRequestBuilder] for the specified table name.
func (rB *NowRequestBuilder) Table(tableName string) *tableapi.TableRequestBuilder[*tableapi.TableRecord] {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters[tablePathKey] = tableName
	return tableapi.NewDefaultTableRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Stats returns a [aggregationapi.StatsRequestBuilder] for the specified table.
func (rB *NowRequestBuilder) Stats(tableName string) *aggregationapi.StatsRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters[tablePathKey] = tableName
	return aggregationapi.NewStatsRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Attachment returns an [attachmentapi.AttachmentRequestBuilder] associated with the [NowRequestBuilder].
func (rB *NowRequestBuilder) Attachment() *attachmentapi.AttachmentRequestBuilder {
	return attachmentapi.NewAttachmentRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Batch returns a [batchapi.BatchRequestBuilder] associated with the [NowRequestBuilder].
func (rB *NowRequestBuilder) Batch() *batchapi.BatchRequestBuilder {
	return batchapi.NewBatchRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Documents returns a [documentsapi.DocumentsRequestBuilder] associated with the [NowRequestBuilder].
func (rB *NowRequestBuilder) Documents() *documentsapi.DocumentsRequestBuilder {
	return documentsapi.NewDocumentsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Cmdb returns a [cmdbinstanceapi.CmdbRequestBuilder] associated with the [NowRequestBuilder].
func (rB *NowRequestBuilder) Cmdb() *cmdbinstanceapi.CmdbRequestBuilder {
	return cmdbinstanceapi.NewCmdbRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Account returns an [accountapi.AccountRequestBuilder] associated with the [NowRequestBuilder].
func (rB *NowRequestBuilder) Account() *accountapi.AccountRequestBuilder {
	return accountapi.NewAccountRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ActSub returns an [actsubapi.ActSubRequestBuilder] associated with the [NowRequestBuilder].
func (rB *NowRequestBuilder) ActSub() *actsubapi.ActSubRequestBuilder {
	return actsubapi.NewActSubRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
