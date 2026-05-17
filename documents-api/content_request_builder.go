package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	contentURLTemplate = "{+baseurl}/api/now/documents/{document_sys_id}/content"
)

// ContentRequestBuilder provides operations to manage document content.
type ContentRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewContentRequestBuilderInternal instantiates a new ContentRequestBuilder.
func NewContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ContentRequestBuilder {
	return &ContentRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, contentURLTemplate, pathParameters),
	}
}

// Get fetches and streams the default version attachment for the document.
func (rB *ContentRequestBuilder) Get(ctx context.Context, requestConfiguration *ContentRequestBuilderGetRequestConfiguration) ([]byte, error) {
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

	res, err := rB.GetRequestAdapter().SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	return res.([]byte), nil
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *ContentRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *ContentRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
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
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationOctetStream)

	return kiotaRequestInfo.RequestInformation, nil
}
