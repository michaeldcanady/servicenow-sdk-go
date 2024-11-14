package servicenowsdkgo

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

type NowRequestBuilder struct {
	context.Context
	core.RequestBuilder
}

// NewNowRequestBuilder creates a new instance of the NowRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created NowRequestBuiabstraction
func NewNowRequestBuilder(ctx context.Context, url string, client *ServiceNowClient) *NowRequestBuilder {
	pathParameters := map[string]string{internal.BasePathParameter: url}
	requestBuilder := core.NewRequestBuilder(client, "{+baseurl}/Now", pathParameters) //nolint:staticcheck
	return &NowRequestBuilder{
		Context:        ctx,
		RequestBuilder: *requestBuilder,
	}
}

// Deprecated: deprecated since v{unreleased}. Use `Table2` instead.
// Table returns a TableRequestBuilder associated with the NowRequestBuilder.
// It accepts a table name as a parameter and constructs the URL for table-related requests.
// The returned TableRequestBuilder can be used to build and execute table-related requests.
func (rB *NowRequestBuilder) Table(tableName string) *tableapi.TableRequestBuilder {
	rB.RequestBuilder.PathParameters["table"] = tableName
	return tableapi.NewTableRequestBuilder(rB.Context, rB.RequestBuilder.Client.(*ServiceNowClient),
		rB.RequestBuilder.PathParameters)
}

var _ intCore.ClientSendableAdapterFunc[*ServiceNowClient] = sendableAdapter

func sendableAdapter(adaptee *ServiceNowClient, ctx context.Context, info intCore.RequestInformation, mapping intCore.ErrorMapping) (*http.Response, error) {
	oldInfo := core.NewRequestInformation()
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
		oldInfo.Method = core.GET
	case "post":
		oldInfo.Method = core.POST
	case "patch":
		oldInfo.Method = core.PATCH
	case "delete":
		oldInfo.Method = core.DELETE
	case "options":
		oldInfo.Method = core.OPTIONS
	case "connect":
		oldInfo.Method = core.CONNECT
	case "put":
		oldInfo.Method = core.PUT
	case "trace":
		oldInfo.Method = core.TRACE
	case "head":
		oldInfo.Method = core.HEAD
	}

	return adaptee.SendWithContext(ctx, oldInfo, mapping.(core.ErrorMapping))
}

// Table returns a TableRequestBuilder associated with the NowRequestBuilder.
// It accepts a table name as a parameter and constructs the URL for table-related requests.
// The returned TableRequestBuilder can be used to build and execute table-related requests.
func (rB *NowRequestBuilder) Table2(tableName string) tableapi.TableRequestBuilder2[*tableapi.TableRecordImpl] {
	rB.RequestBuilder.PathParameters["table"] = tableName

	requestBuilder, _ := tableapi.NewDefaultTableRequestBuilder2(
		intCore.NewClietSendableAdapter(sendableAdapter, rB.RequestBuilder.Client.(*ServiceNowClient)),
		rB.RequestBuilder.PathParameters,
	)
	return requestBuilder
}

// Attachment returns an AttachmentRequestBuilder associated with the NowRequestBuilder.
// It allows you to work with attachments and manage attachments in ServiceNow.
func (rB *NowRequestBuilder) Attachment() *attachmentapi.AttachmentRequestBuilder {
	return attachmentapi.NewAttachmentRequestBuilder(rB.RequestBuilder.Client.(*ServiceNowClient), rB.RequestBuilder.PathParameters)
}
