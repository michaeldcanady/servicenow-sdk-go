package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	versionsURLTemplate = "{+baseurl}/api/now/v1/documents/versions/{document_sys_id}"
)

// VersionsRequestBuilder provides operations to manage document versions.
type VersionsRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewVersionsRequestBuilderInternal instantiates a new VersionsRequestBuilder.
func NewVersionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *VersionsRequestBuilder {
	return &VersionsRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, versionsURLTemplate, pathParameters),
	}
}

// Get retrieves the versions of the specified document.
func (rB *VersionsRequestBuilder) Get(ctx context.Context, requestConfiguration *VersionsRequestBuilderGetRequestConfiguration) (*newInternal.BaseServiceNowCollectionResponse[Document], error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, newInternal.ServiceNowCollectionResponseFromDiscriminatorValue[Document](CreateDocumentFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	return res.(*newInternal.BaseServiceNowCollectionResponse[Document]), nil
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *VersionsRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *VersionsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
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
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
