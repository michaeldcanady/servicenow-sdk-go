package attachmentapi

import (
	"context"
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

const (
	// sysIDKey sys id path parameter key
	sysIDKey = "sys_id"
	// attachmentURLTemplate url template for the Service-Now Attachment endpoint
	attachmentURLTemplate = "{+baseurl}/api/now/v1/attachment{?sysparm_limit,sysparm_offset,sysparm_query}"
	// rawURLKey raw url path parameter key
	rawURLKey = "request-raw-url"
)

// AttachmentRequestBuilder2 provides operations to manage Service-Now attachments.
type AttachmentRequestBuilder2 struct {
	abstractions.BaseRequestBuilder
}

// NewAPIV1CompatibleAttachmentRequestBuilder2Internal converts api v1 compatible elements into api v2 compatible elements
func NewAPIV1CompatibleAttachmentRequestBuilder2Internal(
	pathParameters map[string]string,
	client core.Client,
) *AttachmentRequestBuilder2 {
	reqAdapter, _ := internal.NewServiceNowRequestAdapterBase(core.NewAPIV1ClientAdapter(client))

	return NewAttachmentRequestBuilder2Internal(
		pathParameters,
		reqAdapter,
	)
}

// NewAttachmentRequestBuilder2Internal instantiates a new AttachmentRequestBuilder2 with custom parsable for table entries.
func NewAttachmentRequestBuilder2Internal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentRequestBuilder2 {
	m := &AttachmentRequestBuilder2{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, attachmentURLTemplate, pathParameters),
	}
	return m
}

// NewAttachmentRequestBuilder2 instantiates a new AttachmentRequestBuilder2 with custom parsable for table entries.
func NewAttachmentRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentRequestBuilder2 {
	urlParams := make(map[string]string)
	urlParams[rawURLKey] = rawURL
	return NewAttachmentRequestBuilder2Internal(urlParams, requestAdapter)
}

// ByID provides the way to manage attachment item with provided sys id
func (rB *AttachmentRequestBuilder2) ByID(sysID string) *AttachmentItemRequestBuilder {
	if internal.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.BaseRequestBuilder.PathParameters)
	pathParameters[sysIDKey] = sysID

	return NewAttachmentItemRequestBuilderInternal(pathParameters, rB.BaseRequestBuilder.RequestAdapter)
}

// File provides the way to access Service-Now's attachment file API
func (rB *AttachmentRequestBuilder2) File() *AttachmentFileRequestBuilder {
	if internal.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.BaseRequestBuilder.PathParameters)

	return NewAttachmentFileRequestBuilderInternal(pathParameters, rB.BaseRequestBuilder.RequestAdapter)
}

// Upload provides the way to access Service-Now's attachment upload API
func (rB *AttachmentRequestBuilder2) Upload() *AttachmentUploadRequestBuilder {
	if internal.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.BaseRequestBuilder.PathParameters)

	return NewAttachmentUploadRequestBuilderInternal(pathParameters, rB.BaseRequestBuilder.RequestAdapter)
}

// Get returns AttachmentCollectionResponse using provided arguments
func (rB *AttachmentRequestBuilder2) Get(ctx context.Context, requestConfiguration *TableAttachmentRequestBuilder2GetRequestConfiguration) (AttachmentCollectionResponse2, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	if internal.IsNil(requestConfiguration) {
		requestConfiguration = &TableAttachmentRequestBuilder2GetRequestConfiguration{}
	}

	opts := nethttplibrary.NewHeadersInspectionOptions()
	opts.InspectResponseHeaders = true

	requestConfiguration.Options = append(requestConfiguration.Options, opts)

	requestInfo, err := rB.ToGetRequestInformation(ctx, nil, requestConfiguration)
	if err != nil {
		return nil, err
	}

	// TODO: add error factory
	errorMapping := abstractions.ErrorMappings{}

	res, err := rB.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAttachmentCollectionResponse2FromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	snRes, ok := res.(AttachmentCollectionResponse2)
	if !ok {
		return nil, errors.New("res is not AttachmentCollectionResponse2")
	}

	if err := parseNavLinkHeaders(opts.ResponseHeaders.Get("Link"), snRes); err != nil {
		return nil, err
	}

	return snRes, nil
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *AttachmentRequestBuilder2) ToGetRequestInformation(_ context.Context, _ Attachmentable, requestConfiguration *TableAttachmentRequestBuilder2GetRequestConfiguration) (*abstractions.RequestInformation, error) { //nolint:unparam
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := intHttp.NewRequestInformationWithMethodAndURLTemplateAndPathParameters(abstractions.GET, rB.UrlTemplate, rB.PathParameters)
	if !internal.IsNil(requestConfiguration) {
		if params := requestConfiguration.QueryParameters; !internal.IsNil(params) {
			requestInfo.AddQueryParameters(params)
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.AddAll(requestConfiguration.Headers)
	requestInfo.AddRequestOptions(requestConfiguration.Options)
	requestInfo.Headers.TryAdd("Accept", "application/json")

	return &requestInfo.RequestInformation, nil
}
