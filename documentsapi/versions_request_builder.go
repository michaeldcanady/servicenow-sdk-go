package documentsapi

import (
	"context"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	versionsURLTemplate = "{+baseurl}/api/now/v1/documents/versions/{document_sys_id}"
)

// VersionsRequestBuilder provides operations to manage document versions.
type VersionsRequestBuilder struct {
	core.RequestBuilder
}

// NewVersionsRequestBuilderInternal instantiates a new VersionsRequestBuilder.
func NewVersionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *VersionsRequestBuilder {
	return &VersionsRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, versionsURLTemplate, pathParameters),
	}
}

// Get retrieves the versions of the specified document.
func (rB *VersionsRequestBuilder) Get(ctx context.Context, requestConfiguration *VersionsRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowCollectionResponse[Document], error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowCollectionResponseFromDiscriminatorValue[Document](CreateDocumentFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(*core.BaseServiceNowCollectionResponse[Document]), nil
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *VersionsRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *VersionsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		kiotaRequestInfo.Headers.AddAll(requestConfiguration.Headers)
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return kiotaRequestInfo.RequestInformation, nil
}
