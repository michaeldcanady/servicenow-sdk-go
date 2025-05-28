package attachmentapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// attachmentUploadURLTemplate url template for Service-Now's attachment upload endpoint
	attachmentUploadURLTemplate = "{+baseurl}/api/now/v1/attachment/upload"
)

// AttachmentUploadRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentUploadRequestBuilder struct {
	newInternal.RequestBuilder
}

// newAttachmentUploadRequestBuilderInternal instantiates a new AttachmentUploadRequestBuilder with the provided requestBuilder
func newAttachmentUploadRequestBuilderInternal(requestBuilder newInternal.RequestBuilder) *AttachmentUploadRequestBuilder {
	m := &AttachmentUploadRequestBuilder{
		requestBuilder,
	}
	return m
}

// NewAttachmentUploadRequestBuilderInternal instantiates a new AttachmentUploadRequestBuilder with custom parsable for table entries.
func NewAttachmentUploadRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentUploadRequestBuilder {
	return newAttachmentUploadRequestBuilderInternal(
		newInternal.NewBaseRequestBuilder(requestAdapter, attachmentFileURLTemplate, pathParameters),
	)
}

// NewAttachmentUploadRequestBuilder instantiates a new AttachmentUploadRequestBuilder with custom parsable for table entries.
func NewAttachmentUploadRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentUploadRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
	return NewAttachmentUploadRequestBuilderInternal(urlParams, requestAdapter)
}

// Post Uploads the provided attachment using the provided arguments
func (rB *AttachmentUploadRequestBuilder) Post(ctx context.Context, body abstractions.MultipartBody, requestConfiguration *AttachmentUploadRequestBuilderPostRequestConfiguration) (*FileModel, error) {
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

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateFileFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	file, ok := resp.(*FileModel)
	if !ok {
		return nil, errors.New("resp is not Fileable")
	}

	return file, nil
}

// ToPostRequestInformation converts request configurations to Post request information.
func (rB *AttachmentUploadRequestBuilder) ToPostRequestInformation(ctx context.Context, body abstractions.MultipartBody, requestConfiguration *AttachmentUploadRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	kiotaRequestInfo.Headers.TryAdd(newInternal.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), "application/json", body)
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return kiotaRequestInfo.RequestInformation, nil
}
