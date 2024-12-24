package servicenowsdkgo

import (
	"context"
	"maps"
	"net/http"
	"net/url"
	"strings"

	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

// Deprecated: deprecated since v{unreleased}.
//
// NowRequestBuilder ...
type NowRequestBuilder struct {
	core.RequestBuilder
}

// Deprecated: deprecated since v{unreleased}.
//
// NewNowRequestBuilder creates a new instance of the NowRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created NowRequestBuilder
func NewNowRequestBuilder(url string, client *ServiceNowClient) *NowRequestBuilder {
	pathParameters := map[string]string{internal.BasePathParameter: url}
	requestBuilder := core.NewRequestBuilder(client, "{+baseurl}/Now", pathParameters) //nolint:staticcheck
	return &NowRequestBuilder{
		*requestBuilder,
	}
}

// Deprecated: deprecated since v{unreleased}. Use `Table2` instead.
//
// Table returns a TableRequestBuilder associated with the NowRequestBuilder.
// It accepts a table name as a parameter and constructs the URL for table-related requests.
// The returned TableRequestBuilder can be used to build and execute table-related requests.
func (rB *NowRequestBuilder) Table(tableName string) *tableapi.TableRequestBuilder { //nolint: staticcheck
	rB.RequestBuilder.PathParameters["table"] = tableName
	return tableapi.NewTableRequestBuilder(rB.RequestBuilder.Client.(*ServiceNowClient), rB.RequestBuilder.PathParameters)
}

// Table2 provides way to manage Service-Now table entries
func (rB *NowRequestBuilder) Table2(tableName string) *tableapi.TableRequestBuilder2 {
	pathParameters := maps.Clone(rB.RequestBuilder.PathParameters)
	pathParameters["table"] = tableName
	return tableapi.NewAPIV1CompatibleDefaultTableRequestBuilder2Internal(pathParameters, rB.RequestBuilder.Client)
}

var _ intCore.ClientSendableAdapterFunc[*ServiceNowClient] = sendableAdapter

func sendableAdapter(adaptee *ServiceNowClient, ctx context.Context, info intCore.RequestInformation, mapping intCore.ErrorMapping) (*http.Response, error) { //nolint: staticcheck
	oldInfo := core.NewRequestInformation() //nolint: staticcheck
	oldHeaders := http.Header{}

	info.GetHeaders().Iterate(func(s1 string, s2 []string) bool {
		for _, val := range s2 {
			oldHeaders.Add(s1, val)
		}
		return true
	})

	oldInfo.Headers = oldHeaders
	oldInfo.Content = info.GetContent()
	strURL, err := info.Url()
	if err != nil {
		return nil, err
	}
	realURL, err := url.Parse(strURL)
	if err != nil {
		return nil, err
	}
	oldInfo.SetUri(realURL)

	switch strings.ToLower(info.GetMethod()) {
	case "get":
		oldInfo.Method = core.GET //nolint: staticcheck
	case "post":
		oldInfo.Method = core.POST //nolint: staticcheck
	case "patch":
		oldInfo.Method = core.PATCH //nolint: staticcheck
	case "delete":
		oldInfo.Method = core.DELETE //nolint: staticcheck
	case "options":
		oldInfo.Method = core.OPTIONS //nolint: staticcheck
	case "connect":
		oldInfo.Method = core.CONNECT //nolint: staticcheck
	case "put":
		oldInfo.Method = core.PUT //nolint: staticcheck
	case "trace":
		oldInfo.Method = core.TRACE //nolint: staticcheck
	case "head":
		oldInfo.Method = core.HEAD //nolint: staticcheck
	}

	return adaptee.SendWithContext(ctx, oldInfo, mapping.(core.ErrorMapping)) //nolint: staticcheck
}

// Deprecated: deprecated since v{unreleased}. Use `Attachment2` instead.
//
// Attachment returns an AttachmentRequestBuilder associated with the NowRequestBuilder.
// It allows you to work with attachments and manage attachments in ServiceNow.
func (rB *NowRequestBuilder) Attachment() *attachmentapi.AttachmentRequestBuilder { //nolint: staticcheck
	return attachmentapi.NewAttachmentRequestBuilder(rB.RequestBuilder.Client.(*ServiceNowClient), rB.RequestBuilder.PathParameters) //nolint: staticcheck
}

// Attachment2 provides way to manage Service-Now attachments
func (rB *NowRequestBuilder) Attachment2() *attachmentapi.AttachmentRequestBuilder2 {
	pathParameters := maps.Clone(rB.RequestBuilder.PathParameters)
	return attachmentapi.NewAPIV1CompatibleAttachmentRequestBuilder2Internal(pathParameters, rB.RequestBuilder.Client)
}

// Batch providers way to manage Service-Now batch requests
func (rB *NowRequestBuilder) Batch() batchapi.BatchRequestBuilder2 {
	pathParameters := maps.Clone(rB.RequestBuilder.PathParameters)
	return *batchapi.NewAPIV1CompatibleBatchRequestBuilder2Internal(pathParameters, rB.RequestBuilder.Client)
}
