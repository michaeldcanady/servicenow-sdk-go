package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	exploreURLTemplate = "{+baseurl}/api/now/v1/documents/explore{?page,limit,query,table_name,folder_sys_id,record_sys_id}"
)

// ExploreRequestBuilder provides operations to manage the explore endpoint.
type ExploreRequestBuilder struct {
	core.RequestBuilder
}

// NewExploreRequestBuilderInternal instantiates a new ExploreRequestBuilder.
func NewExploreRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ExploreRequestBuilder {
	return &ExploreRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, exploreURLTemplate, pathParameters),
	}
}

// Get retrieves folder and document metadata with filters, sorting, and pagination.
func (rB *ExploreRequestBuilder) Get(ctx context.Context, requestConfiguration *ExploreRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowCollectionResponse[Document], error) {
	if conversion.IsNil(rB) {
		return nil, nil
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
func (rB *ExploreRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *ExploreRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		kiotaRequestInfo.Headers.AddAll(requestConfiguration.Headers)
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
		kiotaRequestInfo.AddQueryParameters(requestConfiguration.QueryParameters)
	}
	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return kiotaRequestInfo.RequestInformation, nil
}
