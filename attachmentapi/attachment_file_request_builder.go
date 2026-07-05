package attachmentapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// attachmentFileURLTemplate the url template for Service-Now's attachment file endpoint
	attachmentFileURLTemplate = "{+baseurl}/api/now/v1/attachment/file{?encryption_context,file_name,table_name,table_sys_id}"
)

// AttachmentFileRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentFileRequestBuilder struct {
	core.RequestBuilder
}

// newAttachmentFileRequestBuilderInternal instantiates a new AttachmentFileRequestBuilder with the provided requestBuilder
func newAttachmentFileRequestBuilderInternal(requestBuilder core.RequestBuilder) *AttachmentFileRequestBuilder {
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
		core.NewBaseRequestBuilder(requestAdapter, attachmentFileURLTemplate, pathParameters),
	)
}

// NewAttachmentFileRequestBuilder instantiates a new AttachmentFileRequestBuilder with custom parsable for table entries.
func NewAttachmentFileRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentFileRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[internal.RawURLKey] = rawURL
	return NewAttachmentFileRequestBuilderInternal(urlParams, requestAdapter)
}

// Post uploads provided content to Service-Now using provided parameters
func (rB *AttachmentFileRequestBuilder) Post(ctx context.Context, media *Media, requestConfiguration *AttachmentFileRequestBuilderPostRequestConfiguration) (core.ServiceNowItemResponse[*File], error) {
	if conversion.IsNil(rB) {
		return nil, nil
	}

	if conversion.IsNil(requestConfiguration) {
		return nil, errors.New("requestConfiguration is nil")
	}

	if requestConfiguration.QueryParameters == nil {
		return nil, errors.New("requestConfiguration.QueryParameters is nil")
	}

	if requestConfiguration.QueryParameters.TableSysID == nil || *requestConfiguration.QueryParameters.TableSysID == "" {
		return nil, errors.New("requestConfiguration.QueryParameters.TableSysID is nil or empty")
	}

	if requestConfiguration.QueryParameters.TableName == nil || *requestConfiguration.QueryParameters.TableName == "" {
		return nil, errors.New("requestConfiguration.QueryParameters.TableName is nil or empty")
	}

	if requestConfiguration.QueryParameters.FileName == nil || *requestConfiguration.QueryParameters.FileName == "" {
		return nil, errors.New("requestConfiguration.QueryParameters.FileName is nil or empty")
	}

	if conversion.IsNil(media) {
		return nil, errors.New("media is nil")
	}

	if media.contentType == "" {
		return nil, errors.New("media.contentType is nil or empty")
	}

	if len(media.data) == 0 {
		return nil, errors.New("media.data is nil or empty")
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, media, requestConfiguration)
	// unable to test since nothing returns an error
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	requestAdapter := rB.GetRequestAdapter()
	if conversion.IsNil(requestAdapter) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	resp, err := requestAdapter.Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[*File](CreateFileFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(resp) {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := resp.(core.ServiceNowItemResponse[*File])
	if !ok {
		return nil, errors.New("resp is not ServiceNowItemResponse[*File]")
	}

	return typedResp, nil
}

// ToPostRequestInformation converts request configurations to Post request information.
func (rB *AttachmentFileRequestBuilder) ToPostRequestInformation(_ context.Context, media *Media, requestConfiguration *AttachmentFileRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	requestAdapter := rB.GetRequestAdapter()
	if conversion.IsNil(requestAdapter) {
		return nil, errors.New("requestAdapter is nil")
	}

	kiotaRequestInfo.SetStreamContentAndContentType(media.GetData(), media.GetContentType())

	return kiotaRequestInfo.RequestInformation, nil
}
