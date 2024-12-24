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
	// attachmentFileURLTemplate the url template for Service-Now's attachment file endpoint
	attachmentFileURLTemplate = "{+baseurl}/api/now/v1/attachment/file{?encryption_context,file_name,table_name,table_sys_id}"
)

// AttachmentFileRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentFileRequestBuilder struct {
	abstractions.BaseRequestBuilder
}

// NewAPIV1CompatibleAttachmentFileRequestBuilderInternal converts api v1 compatible elements into api v2 compatible elements
func NewAPIV1CompatibleAttachmentFileRequestBuilderInternal(
	pathParameters map[string]string,
	client core.Client, //nolint: staticcheck
) *AttachmentFileRequestBuilder {
	reqAdapter, _ := internal.NewServiceNowRequestAdapterBase(core.NewAPIV1ClientAdapter(client))

	return NewAttachmentFileRequestBuilderInternal(
		pathParameters,
		reqAdapter,
	)
}

// NewAttachmentFileRequestBuilderInternal instantiates a new AttachmentFileRequestBuilder with custom parsable for table entries.
func NewAttachmentFileRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentFileRequestBuilder {
	m := &AttachmentFileRequestBuilder{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, attachmentFileURLTemplate, pathParameters),
	}
	return m
}

// NewAttachmentFileRequestBuilder instantiates a new AttachmentFileRequestBuilder with custom parsable for table entries.
func NewAttachmentFileRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentFileRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[intCore.RawURLKey] = rawURL
	return NewAttachmentFileRequestBuilderInternal(urlParams, requestAdapter)
}

// Post uploads provided content to Service-Now using provided parameters
func (rB *AttachmentFileRequestBuilder) Post(ctx context.Context, contentType string, data []byte, requestConfiguration *TableAttachmentFileRequestBuilderPostRequestConfiguration) (Fileable, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	if internal.IsNil(requestConfiguration) || internal.IsNil(requestConfiguration.QueryParameters) {
		return nil, errors.New("requestConfiguration or requestConfiguration.QueryParameters can't be empty")
	}

	if requestConfiguration.QueryParameters.TableSysID == nil || *requestConfiguration.QueryParameters.TableSysID == "" {
		return nil, errors.New("requestConfiguration.QueryParameters.TableSysId can't be empty")
	}

	if requestConfiguration.QueryParameters.TableName == nil || *requestConfiguration.QueryParameters.TableName == "" {
		return nil, errors.New("requestConfiguration.QueryParameters.TableName can't be empty")
	}

	if requestConfiguration.QueryParameters.FileName == nil || *requestConfiguration.QueryParameters.FileName == "" {
		return nil, errors.New("requestConfiguration.QueryParameters.FileName can't be empty")
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, contentType, data, requestConfiguration)
	if err != nil {
		return nil, err
	}

	// TODO: add error factory
	errorMapping := abstractions.ErrorMappings{}

	resp, err := rB.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	typedResp, ok := resp.(Fileable)
	if !ok {
		return nil, errors.New("resp is not Fileable")
	}

	return typedResp, nil
}

// ToPostRequestInformation converts request configurations to Post request information.
func (rB *AttachmentFileRequestBuilder) ToPostRequestInformation(_ context.Context, contentType string, data []byte, requestConfiguration *TableAttachmentFileRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) { //nolint:unparam
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := intHttp.NewRequestInformationWithMethodAndURLTemplateAndPathParameters(abstractions.POST, rB.UrlTemplate, rB.PathParameters)
	if !internal.IsNil(requestConfiguration) {
		if params := requestConfiguration.QueryParameters; !internal.IsNil(params) {
			requestInfo.AddQueryParameters(params)
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.AddAll(requestConfiguration.Headers)
	requestInfo.AddRequestOptions(requestConfiguration.Options)
	requestInfo.SetStreamContentAndContentType(data, contentType)
	requestInfo.Headers.TryAdd("Accept", "application/json")

	return &requestInfo.RequestInformation, nil
}
