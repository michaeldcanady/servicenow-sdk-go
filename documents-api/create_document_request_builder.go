package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	createDocumentURLTemplate = "{+baseurl}/api/now/documents/createDocument"
)

// CreateDocumentRequestBuilder provides operations to manage the createDocument endpoint.
type CreateDocumentRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewCreateDocumentRequestBuilderInternal instantiates a new CreateDocumentRequestBuilder.
func NewCreateDocumentRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CreateDocumentRequestBuilder {
	return &CreateDocumentRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, createDocumentURLTemplate, pathParameters),
	}
}

// Post creates or links a document from an attachment, DMS repo or cloud.
func (rB *CreateDocumentRequestBuilder) Post(ctx context.Context, requestConfiguration *CreateDocumentRequestBuilderPostRequestConfiguration) (*newInternal.BaseServiceNowItemResponse[Document], error) {
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
func (rB *CreateDocumentRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CreateDocumentRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
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
			err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), newInternal.ContentTypeApplicationJSON, data)
			if err != nil {
				return nil, err
			}
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
