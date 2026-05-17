package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	attachURLTemplate = "{+baseurl}/api/now/documents/{provider_id}/attach"
)

// AttachRequestBuilder provides operations to manage the attach endpoint.
type AttachRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewAttachRequestBuilderInternal instantiates a new AttachRequestBuilder.
func NewAttachRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AttachRequestBuilder {
	return &AttachRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, attachURLTemplate, pathParameters),
	}
}

// Post attaches a document using the specified provider.
func (rB *AttachRequestBuilder) Post(ctx context.Context, requestConfiguration *AttachRequestBuilderPostRequestConfiguration) (*newInternal.BaseServiceNowItemResponse[Document], error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, newInternal.ServiceNowItemResponseFromDiscriminatorValue[Document](CreateDocumentFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	return res.(*newInternal.BaseServiceNowItemResponse[Document]), nil
}

// ToPostRequestInformation ...
func (rB *AttachRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AttachRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if data := requestConfiguration.Data; !internal.IsNil(data) {
			kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), newInternal.ContentTypeApplicationJSON, data)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
