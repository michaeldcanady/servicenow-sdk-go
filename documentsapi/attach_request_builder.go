package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	attachURLTemplate = "{+baseurl}/api/now/v1/documents/{provider_id}/attach"
)

// AttachRequestBuilder provides operations to manage the attach endpoint.
type AttachRequestBuilder struct {
	internal.RequestBuilder
}

// NewAttachRequestBuilderInternal instantiates a new AttachRequestBuilder.
func NewAttachRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AttachRequestBuilder {
	return &AttachRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, attachURLTemplate, pathParameters),
	}
}

// Post attaches a document using the specified provider.
func (rB *AttachRequestBuilder) Post(ctx context.Context, requestConfiguration *AttachRequestBuilderPostRequestConfiguration) (*internal.BaseServiceNowItemResponse[Document], error) {
	if conversion.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := internal.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowItemResponseFromDiscriminatorValue[Document](CreateDocumentFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(*internal.BaseServiceNowItemResponse[Document]), nil
}

// ToPostRequestInformation ...
func (rB *AttachRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AttachRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if data := requestConfiguration.Data; !conversion.IsNil(data) {
			err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, data)
			if err != nil {
				return nil, err
			}
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
