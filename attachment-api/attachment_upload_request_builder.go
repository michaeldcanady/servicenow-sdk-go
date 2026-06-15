package attachmentapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// attachmentUploadURLTemplate url template for Service-Now's attachment upload endpoint
	attachmentUploadURLTemplate = "{+baseurl}/api/now/v1/attachment/upload"
)

// AttachmentUploadRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentUploadRequestBuilder struct {
	internal.RequestBuilder
}

// newAttachmentUploadRequestBuilderInternal instantiates a new AttachmentUploadRequestBuilder with the provided requestBuilder
func newAttachmentUploadRequestBuilderInternal(requestBuilder internal.RequestBuilder) *AttachmentUploadRequestBuilder {
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
		internal.NewBaseRequestBuilder(requestAdapter, attachmentUploadURLTemplate, pathParameters),
	)
}

// NewAttachmentUploadRequestBuilder instantiates a new AttachmentUploadRequestBuilder with custom parsable for table entries.
func NewAttachmentUploadRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentUploadRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[internal.RawURLKey] = rawURL
	return NewAttachmentUploadRequestBuilderInternal(urlParams, requestAdapter)
}

// Post Uploads the provided attachment using the provided arguments
func (rB *AttachmentUploadRequestBuilder) Post(ctx context.Context, body abstractions.MultipartBody, requestConfiguration *AttachmentUploadRequestBuilderPostRequestConfiguration) (*FileModel, error) {
	if conversion.IsNil(rB) {
		return nil, nil
	}

	if conversion.IsNil(body) {
		return nil, errors.New("body is nil")
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
		"XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
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
	if conversion.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}
	return kiotaRequestInfo.RequestInformation, nil
}
