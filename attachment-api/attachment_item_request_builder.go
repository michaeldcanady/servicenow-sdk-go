package attachmentapi

import (
	"context"
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	intHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// attachmentItemURLTemplate the url template for Service-Now's attachment item endpoint
	attachmentItemURLTemplate = "{+baseurl}/api/now/v1/attachment{/sys_id}"
)

// AttachmentItemRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentItemRequestBuilder struct {
	abstractions.BaseRequestBuilder
}

// NewAPIV1CompatibleAttachmentItemRequestBuilderInternal converts api v1 compatible elements into api v2 compatible elements
func NewAPIV1CompatibleAttachmentItemRequestBuilderInternal(
	pathParameters map[string]string,
	client core.Client, //nolint: staticcheck
) *AttachmentItemRequestBuilder {
	reqAdapter, _ := internal.NewServiceNowRequestAdapterBase(core.NewAPIV1ClientAdapter(client))

	return NewAttachmentItemRequestBuilderInternal(
		pathParameters,
		reqAdapter,
	)
}

// NewAttachmentItemRequestBuilderInternal instantiates a new AttachmentItemRequestBuilder with custom parsable for table entries.
func NewAttachmentItemRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentItemRequestBuilder {
	m := &AttachmentItemRequestBuilder{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, attachmentItemURLTemplate, pathParameters),
	}
	return m
}

// NewAttachmentItemRequestBuilder instantiates a new AttachmentItemRequestBuilder with custom parsable for table entries.
func NewAttachmentItemRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[intCore.RawURLKey] = rawURL
	return NewAttachmentItemRequestBuilderInternal(urlParams, requestAdapter)
}

// File provides way to manage Service-Now's Attachment Item File endpoint
func (rB *AttachmentItemRequestBuilder) File() *AttachmentItemFileRequestBuilder {
	if internal.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.BaseRequestBuilder.PathParameters)

	return NewAttachmentItemFileRequestBuilderInternal(pathParameters, rB.BaseRequestBuilder.RequestAdapter)
}

// Get returns an Attachment using the provided arguments
func (rB *AttachmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TableAttachmentItemRequestBuilderGetRequestConfiguration) (Attachmentable, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, nil, requestConfiguration)
	if err != nil {
		return nil, err
	}

	// TODO: add error factory
	errorMapping := abstractions.ErrorMappings{}

	res, err := rB.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAttachmentFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	typedRes, ok := interface{}(res).(Attachmentable)
	if !ok {
		return nil, errors.New("res is not Attachmentable")
	}

	return typedRes, nil
}

// Delete removes the attachment item using the provided arguments
func (rB *AttachmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TableAttachmentItemRequestBuilderDeleteRequestConfiguration) error {
	if internal.IsNil(rB) {
		return nil
	}

	requestInfo, err := rB.ToDeleteRequestInformation(ctx, nil, requestConfiguration)
	if err != nil {
		return err
	}

	// TODO: add error factory
	errorMapping := abstractions.ErrorMappings{}

	return rB.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *AttachmentItemRequestBuilder) ToGetRequestInformation(_ context.Context, _ Attachmentable, requestConfiguration *TableAttachmentItemRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) { //nolint:unparam
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := intHttp.NewRequestInformationWithMethodAndURLTemplateAndPathParameters(abstractions.GET, rB.UrlTemplate, rB.PathParameters)
	if !internal.IsNil(requestConfiguration) {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.AddAll(requestConfiguration.Headers)
	requestInfo.AddRequestOptions(requestConfiguration.Options)
	requestInfo.Headers.TryAdd("Accept", "application/json")

	return &requestInfo.RequestInformation, nil
}

// ToDeleteRequestInformation converts request configurations to Delete request information.
func (rB *AttachmentItemRequestBuilder) ToDeleteRequestInformation(_ context.Context, _ Attachmentable, requestConfiguration *TableAttachmentItemRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) { //nolint:unparam
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := intHttp.NewRequestInformationWithMethodAndURLTemplateAndPathParameters(abstractions.DELETE, rB.UrlTemplate, rB.PathParameters)
	if !internal.IsNil(requestConfiguration) {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.AddAll(requestConfiguration.Headers)
	requestInfo.AddRequestOptions(requestConfiguration.Options)
	requestInfo.Headers.TryAdd("Accept", "application/json")

	return &requestInfo.RequestInformation, nil
}
