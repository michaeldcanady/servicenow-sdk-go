package attachmentapi

import (
	"context"
	"errors"

	model "github.com/michaeldcanady/servicenow-sdk-go/internal/errors"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// attachmentUploadURLTemplate url template for Service-Now's attachment upload endpoint
	attachmentUploadURLTemplate = "{+baseurl}/api/now/v1/attachment/upload"
)

// AttachmentUploadRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentUploadRequestBuilder struct {
	kiota.RequestBuilder
}

// newAttachmentUploadRequestBuilderInternal instantiates a new AttachmentUploadRequestBuilder with the provided requestBuilder
func newAttachmentUploadRequestBuilderInternal(requestBuilder kiota.RequestBuilder) *AttachmentUploadRequestBuilder {
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
		kiota.NewBaseRequestBuilder(requestAdapter, attachmentUploadURLTemplate, pathParameters),
	)
}

// NewAttachmentUploadRequestBuilder instantiates a new AttachmentUploadRequestBuilder with custom parsable for table entries.
func NewAttachmentUploadRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentUploadRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[utils.RawURLKey] = rawURL
	return NewAttachmentUploadRequestBuilderInternal(urlParams, requestAdapter)
}

// Post Uploads the provided attachment using the provided arguments
func (rB *AttachmentUploadRequestBuilder) Post(ctx context.Context, body abstractions.MultipartBody, requestConfiguration *AttachmentUploadRequestBuilderPostRequestConfiguration) (*FileModel, error) {
	if utils.IsNil(rB) {
		return nil, nil
	}

	if utils.IsNil(body) {
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
		"XXX": model.CreateServiceNowErrorFromDiscriminatorValue,
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
	if utils.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &kiota.KiotaRequestInformation{RequestInformation: requestInfo}
	if !utils.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !utils.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !utils.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), utils.ContentTypeApplicationJSON)

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), utils.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}
	return kiotaRequestInfo.RequestInformation, nil
}
