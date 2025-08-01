package attachmentapi

import (
	"context"
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// attachmentItemURLTemplate the url template for Service-Now's attachment item endpoint
	attachmentItemURLTemplate = "{+baseurl}/api/now/v1/attachment{/sys_id}"
)

// AttachmentItemRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentItemRequestBuilder struct {
	newInternal.RequestBuilder
}

// newAttachmentItemRequestBuilderInternal instantiates a new AttachmentItemRequestBuilder with the provided requestBuilder
func newAttachmentItemRequestBuilderInternal(requestBuilder newInternal.RequestBuilder) *AttachmentItemRequestBuilder {
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
		newInternal.NewBaseRequestBuilder(requestAdapter, attachmentItemURLTemplate, pathParameters),
	)
}

func (rB *AttachmentItemRequestBuilder) isNil() bool {
	return internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder)
}

// NewAttachmentItemRequestBuilder instantiates a new AttachmentItemRequestBuilder with custom parsable for table entries.
func NewAttachmentItemRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
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
func (rB *AttachmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AttachmentItemRequestBuilderGetRequestConfiguration) (*Attachment2Model, error) {
	if rB.isNil() {
		return nil, nil
	}

	requestInfo := rB.ToGetRequestInformation(ctx, requestConfiguration)

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAttachment2FromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, errors.New("response is nil")
	}

	typedRes, ok := res.(*Attachment2Model)
	if !ok {
		return nil, errors.New("res is not *Attachment2Model")
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
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
}

// ToGetRequestInformation converts request configurations to Post request information.
func (rB *AttachmentItemRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *AttachmentItemRequestBuilderGetRequestConfiguration) *abstractions.RequestInformation {
	if rB.isNil() {
		return nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	requestInfo.Headers.TryAdd(newInternal.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)
	return kiotaRequestInfo.RequestInformation
}

// ToDeleteRequestInformation converts request configurations to Delete request information.
func (rB *AttachmentItemRequestBuilder) ToDeleteRequestInformation(_ context.Context, requestConfiguration *AttachmentItemRequestBuilderDeleteRequestConfiguration) *abstractions.RequestInformation {
	if rB.isNil() {
		return nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	requestInfo.Headers.TryAdd(newInternal.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)
	return kiotaRequestInfo.RequestInformation
}
