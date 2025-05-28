package attachmentapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// attachmentFileURLTemplate the url template for Service-Now's attachment file endpoint
	attachmentFileURLTemplate = "{+baseurl}/api/now/v1/attachment/file{?encryption_context,file_name,table_name,table_sys_id}"
)

// AttachmentFileRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentFileRequestBuilder struct {
	newInternal.RequestBuilder
}

// newAttachmentFileRequestBuilderInternal instantiates a new AttachmentFileRequestBuilder with the provided requestBuilder
func newAttachmentFileRequestBuilderInternal(requestBuilder newInternal.RequestBuilder) *AttachmentFileRequestBuilder {
	m := &AttachmentFileRequestBuilder{
		requestBuilder,
	}
	return m
}

// NewAttachmentFileRequestBuilderInternal instantiates a new AttachmentFileRequestBuilder with custom parsable for table entries.
func NewAttachmentFileRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentFileRequestBuilder {
	return newAttachmentFileRequestBuilderInternal(
		newInternal.NewBaseRequestBuilder(requestAdapter, attachmentFileURLTemplate, pathParameters),
	)
}

// NewAttachmentFileRequestBuilder instantiates a new AttachmentFileRequestBuilder with custom parsable for table entries.
func NewAttachmentFileRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentFileRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
	return NewAttachmentFileRequestBuilderInternal(urlParams, requestAdapter)
}

// Post uploads provided content to Service-Now using provided parameters
func (rB *AttachmentFileRequestBuilder) Post(ctx context.Context, contentType string, data []byte, requestConfiguration *AttachmentFileRequestBuilderPostRequestConfiguration) (*FileModel, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	if internal.IsNil(requestConfiguration) || internal.IsNil(requestConfiguration.QueryParameters) {
		return nil, errors.New("requestConfiguration or requestConfiguration.QueryParameters can't be empty")
	}

	if requestConfiguration.QueryParameters.TableSysID == nil || *requestConfiguration.QueryParameters.TableSysID == "" {
		return nil, errors.New("requestConfiguration.QueryParameters.TableSysID can't be empty")
	}

	if requestConfiguration.QueryParameters.TableName == nil || *requestConfiguration.QueryParameters.TableName == "" {
		return nil, errors.New("requestConfiguration.QueryParameters.TableName can't be empty")
	}

	if requestConfiguration.QueryParameters.FileName == nil || *requestConfiguration.QueryParameters.FileName == "" {
		return nil, errors.New("requestConfiguration.QueryParameters.FileName can't be empty")
	}

	if contentType == "" {
		return nil, errors.New("contentType can't be empty")
	}

	if data == nil || len(data) == 0 {
		return nil, errors.New("data is empty")
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, contentType, data, requestConfiguration)
	// unable to test since nothing returns an error
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	requestAdapter := rB.GetRequestAdapter()
	if internal.IsNil(requestAdapter) {
		return nil, errors.New("requestAdapter is nil")
	}

	resp, err := requestAdapter.Send(ctx, requestInfo, CreateFileFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	typedResp, ok := resp.(*FileModel)
	if !ok {
		return nil, errors.New("resp is not *FileModel")
	}

	return typedResp, nil
}

// ToPostRequestInformation converts request configurations to Post request information.
func (rB *AttachmentFileRequestBuilder) ToPostRequestInformation(_ context.Context, contentType string, body []byte, requestConfiguration *AttachmentFileRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if parameters := requestConfiguration.QueryParameters; !internal.IsNil(parameters) {
			kiotaRequestInfo.AddQueryParameters(parameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(newInternal.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	kiotaRequestInfo.SetStreamContentAndContentType(body, contentType)
	return kiotaRequestInfo.RequestInformation, nil
}
