package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	contentURLTemplate = "{+baseurl}/api/now/v1/documents/{document_sys_id}/content"
)

// ContentRequestBuilder provides operations to manage document content.
type ContentRequestBuilder struct {
	internal.RequestBuilder
}

// NewContentRequestBuilderInternal instantiates a new ContentRequestBuilder.
func NewContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ContentRequestBuilder {
	return &ContentRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, contentURLTemplate, pathParameters),
	}
}

// Get fetches and streams the default version attachment for the document.
func (rB *ContentRequestBuilder) Get(ctx context.Context, requestConfiguration *ContentRequestBuilderGetRequestConfiguration) ([]byte, error) {
	if conversion.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.([]byte), nil
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *ContentRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *ContentRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationOctetStream)

	return kiotaRequestInfo.RequestInformation, nil
}
