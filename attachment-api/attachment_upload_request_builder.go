package attachmentapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	intHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// attachmentUploadURLTemplate url template for Service-Now's attachment upload endpoint
	attachmentUploadURLTemplate = "{+baseurl}/api/now/v1/attachment/upload"
)

// AttachmentUploadRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentUploadRequestBuilder struct {
	abstractions.BaseRequestBuilder
}

// NewAPIV1CompatibleAttachmentUploadRequestBuilderInternal converts api v1 compatible elements into api v2 compatible elements
func NewAPIV1CompatibleAttachmentUploadRequestBuilderInternal(
	pathParameters map[string]string,
	client core.Client,
) *AttachmentUploadRequestBuilder {
	reqAdapter, _ := internal.NewServiceNowRequestAdapterBase(core.NewAPIV1ClientAdapter(client))

	return NewAttachmentUploadRequestBuilderInternal(
		pathParameters,
		reqAdapter,
	)
}

// NewAttachmentUploadRequestBuilderInternal instantiates a new AttachmentUploadRequestBuilder with custom parsable for table entries.
func NewAttachmentUploadRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentUploadRequestBuilder {
	m := &AttachmentUploadRequestBuilder{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, attachmentUploadURLTemplate, pathParameters),
	}
	return m
}

// NewAttachmentUploadRequestBuilder instantiates a new AttachmentUploadRequestBuilder with custom parsable for table entries.
func NewAttachmentUploadRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentUploadRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[intCore.RawURLKey] = rawURL
	return NewAttachmentUploadRequestBuilderInternal(urlParams, requestAdapter)
}

// Post Uploads the provided attachment using the provided arguments
func (rB *AttachmentUploadRequestBuilder) Post(ctx context.Context, body abstractions.MultipartBody, requestConfiguration *AttachmentUploadRequestBuilderPostRequestConfiguration) (Fileable, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	contentType, err := body.GetPartValue("Content-Type")
	if err != nil {
		return nil, err
	}
	if contentType == nil {
		return nil, errors.New("Content-Type is required")
	}
	tableName, err := body.GetPartValue("table_name")
	if err != nil {
		return nil, err
	}
	if tableName == nil {
		return nil, errors.New("table_name is required")
	}
	tableSysID, err := body.GetPartValue("table_sys_id")
	if err != nil {
		return nil, err
	}
	if tableSysID == nil {
		return nil, errors.New("table_sys_id is required")
	}
	uploadFile, err := body.GetPartValue("uploadFile")
	if err != nil {
		return nil, err
	}
	if uploadFile == nil {
		return nil, errors.New("uploadFile is required")
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}

	// TODO: add error factory
	errorMapping := abstractions.ErrorMappings{}

	resp, err := rB.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	file, ok := resp.(Fileable)
	if !ok {
		return nil, errors.New("resp is not Fileable")
	}

	return file, nil
}

// ToPostRequestInformation converts request configurations to Get request information.
func (rB *AttachmentUploadRequestBuilder) ToPostRequestInformation(ctx context.Context, body abstractions.MultipartBody, requestConfiguration *AttachmentUploadRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) { //nolint:unparam
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := intHttp.NewRequestInformationWithMethodAndURLTemplateAndPathParameters(abstractions.POST, rB.UrlTemplate, rB.PathParameters)
	if !internal.IsNil(requestConfiguration) {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	if err := requestInfo.SetContentFromParsable(ctx, rB.BaseRequestBuilder.RequestAdapter, "multipart/form-data", body); err != nil {
		return nil, err
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")

	return &requestInfo.RequestInformation, nil
}
