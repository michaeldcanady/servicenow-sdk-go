package attachmentapi

import (
	"context"
	"errors"
	"maps"

	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	model "github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// attachmentItemURLTemplate the url template for Service-Now's attachment item endpoint
	attachmentItemURLTemplate = "{+baseurl}/api/now/v1/attachment{/sys_id}"
)

// AttachmentItemRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentItemRequestBuilder struct {
	kiota.RequestBuilder
}

// newAttachmentItemRequestBuilderInternal instantiates a new AttachmentItemRequestBuilder with the provided requestBuilder
func newAttachmentItemRequestBuilderInternal(requestBuilder kiota.RequestBuilder) *AttachmentItemRequestBuilder {
	m := &AttachmentItemRequestBuilder{
		requestBuilder,
	}
	return m
}

// NewAttachmentItemRequestBuilderInternal instantiates a new AttachmentItemRequestBuilder with custom parsable for table entries.
func NewAttachmentItemRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentItemRequestBuilder {
	return newAttachmentItemRequestBuilderInternal(
		kiota.NewBaseRequestBuilder(requestAdapter, attachmentItemURLTemplate, pathParameters),
	)
}

func NewAttachmentItemRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[utils.RawURLKey] = rawURL
	return NewAttachmentItemRequestBuilderInternal(urlParams, requestAdapter)
}

func (rB *AttachmentItemRequestBuilder) isNil() bool {
	return utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder)
}

// NewAttachmentItemRequestBuilder instantiates a new AttachmentItemRequestBuilder with custom parsable for table entries.
func NewAttachmentItemRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[utils.RawURLKey] = rawURL
	return NewAttachmentItemRequestBuilderInternal(urlParams, requestAdapter)
}

// File provides way to manage Service-Now's Attachment Item File endpoint
func (rB *AttachmentItemRequestBuilder) File() *AttachmentItemFileRequestBuilder {
	if rB.isNil() {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())

	return NewAttachmentItemFileRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Get returns an Attachment using the provided arguments
func (rB *AttachmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AttachmentItemRequestBuilderGetRequestConfiguration) (newInternal.ServiceNowItemResponse[Attachment2], error) {
	if rB.isNil() {
		return nil, nil
	}

	requestInfo := rB.ToGetRequestInformation(ctx, requestConfiguration)

	errorMapping := abstractions.ErrorMappings{
		"XXX": model.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, newInternal.ServiceNowItemResponseFromDiscriminatorValue[Attachment2](CreateAttachment2FromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if utils.IsNil(res) {
		return nil, errors.New("response is nil")
	}

	typedRes, ok := res.(newInternal.ServiceNowItemResponse[Attachment2])
	if !ok {
		return nil, errors.New("res is not ServiceNowItemResponse[Attachment2]")
	}

	return typedRes, nil
}

// Delete removes the attachment item using the provided arguments
func (rB *AttachmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AttachmentItemRequestBuilderDeleteRequestConfiguration) error {
	if rB.isNil() {
		return nil
	}

	requestInfo := rB.ToDeleteRequestInformation(ctx, requestConfiguration)

	errorMapping := abstractions.ErrorMappings{
		"XXX": model.CreateServiceNowErrorFromDiscriminatorValue,
	}

	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
}

// ToGetRequestInformation converts request configurations to Post request information.
func (rB *AttachmentItemRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *AttachmentItemRequestBuilderGetRequestConfiguration) *abstractions.RequestInformation {
	if rB.isNil() {
		return nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &kiota.KiotaRequestInformation{RequestInformation: requestInfo}
	if !utils.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !utils.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !utils.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	requestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), utils.ContentTypeApplicationJSON)
	return kiotaRequestInfo.RequestInformation
}

// ToDeleteRequestInformation converts request configurations to Delete request information.
func (rB *AttachmentItemRequestBuilder) ToDeleteRequestInformation(_ context.Context, requestConfiguration *AttachmentItemRequestBuilderDeleteRequestConfiguration) *abstractions.RequestInformation {
	if rB.isNil() {
		return nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &kiota.KiotaRequestInformation{RequestInformation: requestInfo}
	if !utils.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !utils.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !utils.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	requestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), utils.ContentTypeApplicationJSON)
	return kiotaRequestInfo.RequestInformation
}
