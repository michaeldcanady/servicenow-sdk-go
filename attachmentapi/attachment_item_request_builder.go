package attachmentapi

import (
	"context"
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// attachmentItemURLTemplate the url template for Service-Now's attachment item endpoint
	attachmentItemURLTemplate = "{+baseurl}/api/now/v1/attachment{/sys_id}"
)

// AttachmentItemRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentItemRequestBuilder struct {
	core.RequestBuilder
}

// newAttachmentItemRequestBuilderInternal instantiates a new AttachmentItemRequestBuilder with the provided requestBuilder
func newAttachmentItemRequestBuilderInternal(requestBuilder core.RequestBuilder) *AttachmentItemRequestBuilder {
	m := &AttachmentItemRequestBuilder{
		requestBuilder,
	}
	return m
}

// NewAttachmentItemRequestBuilderInternal instantiates a new AttachmentItemRequestBuilder with custom parsable for table entries.
func NewAttachmentItemRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentItemRequestBuilder {
	return newAttachmentItemRequestBuilderInternal(
		core.NewBaseRequestBuilder(requestAdapter, attachmentItemURLTemplate, pathParameters),
	)
}

// NewAttachmentItemRequestBuilder instantiates a new AttachmentItemRequestBuilder with custom parsable for table entries.
func NewAttachmentItemRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[internal.RawURLKey] = rawURL
	return NewAttachmentItemRequestBuilderInternal(urlParams, requestAdapter)
}

// File provides way to manage Service-Now's Attachment Item File endpoint
func (rB *AttachmentItemRequestBuilder) File() *AttachmentItemFileRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewAttachmentItemFileRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Get returns an Attachment using the provided arguments
func (rB *AttachmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AttachmentItemRequestBuilderGetRequestConfiguration) (core.ServiceNowItemResponse[*Attachment], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[*Attachment](CreateAttachmentFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}

	typedRes, ok := res.(core.ServiceNowItemResponse[*Attachment])
	if !ok {
		return nil, errors.New("res is not ServiceNowItemResponse[*Attachment]")
	}

	return typedRes, nil
}

// Delete removes the attachment item using the provided arguments
func (rB *AttachmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AttachmentItemRequestBuilderDeleteRequestConfiguration) error {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return snerrors.ErrNilRequestBuilder
	}

	requestInfo, err := rB.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}

	errorMapping := core.DefaultErrorMapping()
	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
}

// ToGetRequestInformation converts request configurations to Post request information.
func (rB *AttachmentItemRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *AttachmentItemRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, requestConfiguration)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}

// ToDeleteRequestInformation converts request configurations to Delete request information.
func (rB *AttachmentItemRequestBuilder) ToDeleteRequestInformation(_ context.Context, requestConfiguration *AttachmentItemRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, requestConfiguration)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
