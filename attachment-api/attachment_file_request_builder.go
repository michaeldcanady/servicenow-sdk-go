package attachmentapi

import (
	"context"
	"errors"

	internalErrors "github.com/michaeldcanady/servicenow-sdk-go/internal/errors"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// attachmentFileURLTemplate the url template for Service-Now's attachment file endpoint
	attachmentFileURLTemplate = "{+baseurl}/api/now/v1/attachment/file{?encryption_context,file_name,table_name,table_sys_id}"
)

// AttachmentFileRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentFileRequestBuilder struct {
	kiota.RequestBuilder
}

// newAttachmentFileRequestBuilderInternal instantiates a new AttachmentFileRequestBuilder with the provided requestBuilder
func newAttachmentFileRequestBuilderInternal(requestBuilder kiota.RequestBuilder) *AttachmentFileRequestBuilder {
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
		kiota.NewBaseRequestBuilder(requestAdapter, attachmentFileURLTemplate, pathParameters),
	)
}

// NewAttachmentFileRequestBuilder instantiates a new AttachmentFileRequestBuilder with custom parsable for table entries.
func NewAttachmentFileRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentFileRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[utils.RawURLKey] = rawURL
	return NewAttachmentFileRequestBuilderInternal(urlParams, requestAdapter)
}

// Post uploads provided content to Service-Now using provided parameters
func (rB *AttachmentFileRequestBuilder) Post(ctx context.Context, media *Media, requestConfiguration *AttachmentFileRequestBuilderPostRequestConfiguration) (*FileModel, error) {
	if utils.IsNil(rB) {
		return nil, nil
	}

	if utils.IsNil(requestConfiguration) || utils.IsNil(requestConfiguration.QueryParameters) {
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

	if utils.IsNil(media) {
		return nil, errors.New("media is nil")
	}

	if media.contentType == "" {
		return nil, errors.New("contentType can't be empty")
	}

	if len(media.data) == 0 {
		return nil, errors.New("data is empty")
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, media, requestConfiguration)
	// unable to test since nothing returns an error
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": internalErrors.CreateServiceNowErrorFromDiscriminatorValue,
	}

	requestAdapter := rB.GetRequestAdapter()
	if utils.IsNil(requestAdapter) {
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
func (rB *AttachmentFileRequestBuilder) ToPostRequestInformation(ctx context.Context, media *Media, requestConfiguration *AttachmentFileRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
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
		if parameters := requestConfiguration.QueryParameters; !utils.IsNil(parameters) {
			kiotaRequestInfo.AddQueryParameters(parameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), utils.ContentTypeApplicationJSON)

	requestAdapter := rB.GetRequestAdapter()
	if utils.IsNil(requestAdapter) {
		return nil, errors.New("requestAdapter is nil")
	}

	if err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), media.GetContentType(), media); err != nil {
		return nil, err
	}

	return kiotaRequestInfo.RequestInformation, nil
}
