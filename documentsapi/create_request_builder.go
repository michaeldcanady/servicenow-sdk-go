package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	createURLTemplate = "{+baseurl}/api/now/v1/documents/create"
)

// CreateRequestBuilder provides operations to manage the create endpoint.
type CreateRequestBuilder struct {
	internal.RequestBuilder
}

// NewCreateRequestBuilderInternal instantiates a new CreateRequestBuilder.
func NewCreateRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, createURLTemplate, pathParameters),
	}
}

// Post creates a new document.
func (rB *CreateRequestBuilder) Post(ctx context.Context, requestConfiguration *CreateRequestBuilderPostRequestConfiguration) (*internal.BaseServiceNowItemResponse[Document], error) {
	if conversion.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"400": internal.CreateBadRequestErrorFromDiscriminatorValue,
		"401": internal.CreateUnauthorizedErrorFromDiscriminatorValue,
		"403": internal.CreateForbiddenErrorFromDiscriminatorValue,
		"404": internal.CreateNotFoundErrorFromDiscriminatorValue,
		"429": internal.CreateTooManyRequestsErrorFromDiscriminatorValue,
		"5XX": internal.CreateServerErrorFromDiscriminatorValue,
		"XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
	}

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
func (rB *CreateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CreateRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
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
