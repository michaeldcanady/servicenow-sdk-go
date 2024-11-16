package attachmentapi

import (
	"context"
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	attachmentItemURLTemplate = "{+baseurl}/api/now/v1/attachment{/sys_id}"
)

// AttachmentItemRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentItemRequestBuilder struct {
	abstractions.BaseRequestBuilder
}

// newRequestBuilder2Internal instantiates a new AttachmentItemRequestBuilder with custom parsable for table entries.
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
	urlParams["request-raw-url"] = rawURL
	return NewAttachmentItemRequestBuilderInternal(urlParams, requestAdapter)
}

func (rB *AttachmentItemRequestBuilder) File() *AttachmentItemFileRequestBuilder {
	if internal.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.BaseRequestBuilder.PathParameters)

	return NewAttachmentItemFileRequestBuilderInternal(pathParameters, rB.BaseRequestBuilder.RequestAdapter)
}

func (rB *AttachmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TableAttachmentItemRequestBuilderGetRequestConfiguration) (Attachmentable, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.toGetRequestInformation(ctx, nil, requestConfiguration)
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

// Delete removes the attachment item.
func (rB *AttachmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TableAttachmentItemRequestBuilderDeleteRequestConfiguration) error {
	if internal.IsNil(rB) {
		return nil
	}

	requestInfo, err := rB.toDeleteRequestInformation(ctx, nil, requestConfiguration)
	if err != nil {
		return err
	}

	// TODO: add error factory
	errorMapping := abstractions.ErrorMappings{}

	return rB.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
}

// toGetRequestInformation converts request configurations to Get request information.
func (rB *AttachmentItemRequestBuilder) toGetRequestInformation(_ context.Context, _ Attachmentable, requestConfiguration *TableAttachmentItemRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) { //nolint:unparam
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

// toDeleteRequestInformation converts request configurations to Delete request information.
func (rB *AttachmentItemRequestBuilder) toDeleteRequestInformation(_ context.Context, _ Attachmentable, requestConfiguration *TableAttachmentItemRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) { //nolint:unparam
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
