package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	versionStateURLTemplate = "{+baseurl}/api/now/v1/documents/versionstate/{version_sys_id}"
)

// VersionStateRequestBuilder provides operations to manage the versionstate endpoint.
type VersionStateRequestBuilder struct {
	internal.RequestBuilder
}

// NewVersionStateRequestBuilderInternal instantiates a new VersionStateRequestBuilder.
func NewVersionStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *VersionStateRequestBuilder {
	return &VersionStateRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, versionStateURLTemplate, pathParameters),
	}
}

// Get retrieves the state of the specified document version.
func (rB *VersionStateRequestBuilder) Get(ctx context.Context, requestConfiguration *VersionStateRequestBuilderGetRequestConfiguration) (*internal.BaseServiceNowItemResponse[Document], error) {
	if conversion.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
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

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *VersionStateRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *VersionStateRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
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
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
