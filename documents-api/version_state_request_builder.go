package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	versionStateURLTemplate = "{+baseurl}/api/now/documents/versionstate/{version_sys_id}"
)

// VersionStateRequestBuilder provides operations to manage the versionstate endpoint.
type VersionStateRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewVersionStateRequestBuilderInternal instantiates a new VersionStateRequestBuilder.
func NewVersionStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *VersionStateRequestBuilder {
	return &VersionStateRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, versionStateURLTemplate, pathParameters),
	}
}

// Get retrieves the state of the specified document version.
func (rB *VersionStateRequestBuilder) Get(ctx context.Context, requestConfiguration *VersionStateRequestBuilderGetRequestConfiguration) (*newInternal.BaseServiceNowItemResponse[Document], error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, newInternal.ServiceNowItemResponseFromDiscriminatorValue[Document](CreateDocumentFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	return res.(*newInternal.BaseServiceNowItemResponse[Document]), nil
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *VersionStateRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *VersionStateRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
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
