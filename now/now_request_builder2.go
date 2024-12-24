package now

import (
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// nowDefaultURLTemplate is the unversioned url template for Service-Now's APIs
	nowDefaultURLTemplate = "{+baseurl}/api/now"
	// nowVersionedURLTemplate is the versioned url template for Service-Now's APIs
	nowVersionedURLTemplate = "{+baseurl}/api/now/{version}" // nolint: unused
)

// NowRequestBuilder2 provides fluent entrypoint in Service-Now's APIs
type NowRequestBuilder2 struct {
	abstractions.BaseRequestBuilder
}

// NewAPIV1CompatibleNowRequestBuilder2Internal converts api v1 compatible elements into api v2 compatible elements
func NewAPIV1CompatibleNowRequestBuilder2Internal(
	pathParameters map[string]string,
	client core.Client, //nolint: staticcheck
) *NowRequestBuilder2 {
	reqAdapter, _ := internal.NewServiceNowRequestAdapterBase(core.NewAPIV1ClientAdapter(client))

	return NewNowRequestBuilder2Internal(
		pathParameters,
		reqAdapter,
	)
}

// NewNowRequestBuilder2Internal instantiates a new NowRequestBuilder2
func NewNowRequestBuilder2Internal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *NowRequestBuilder2 {
	m := &NowRequestBuilder2{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, nowDefaultURLTemplate, pathParameters),
	}
	return m
}

// NewNowRequestBuilder2 instantiates a new NowRequestBuilder2
func NewNowRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *NowRequestBuilder2 {
	urlParams := make(map[string]string)
	urlParams[intCore.RawURLKey] = rawURL
	return NewNowRequestBuilder2Internal(urlParams, requestAdapter)
}

// Table provides way to manage Service-Now table entries
func (rB *NowRequestBuilder2) Table(table string) tableapi.TableRequestBuilder2 {
	rB.BaseRequestBuilder.PathParameters["table"] = table
	return *tableapi.NewDefaultTableRequestBuilder2Internal(rB.BaseRequestBuilder.PathParameters, rB.BaseRequestBuilder.RequestAdapter)
}

// Attachment provides way to manage Service-Now attachments
func (rB *NowRequestBuilder2) Attachment() attachmentapi.AttachmentRequestBuilder2 {
	return *attachmentapi.NewAttachmentRequestBuilder2Internal(rB.BaseRequestBuilder.PathParameters, rB.BaseRequestBuilder.RequestAdapter)
}

// Batch providers way to manage Service-Now batch requests
func (rB *NowRequestBuilder2) Batch() batchapi.BatchRequestBuilder2 {
	return *batchapi.NewBatchRequestBuilder2Internal(rB.BaseRequestBuilder.PathParameters, rB.BaseRequestBuilder.RequestAdapter)
}
