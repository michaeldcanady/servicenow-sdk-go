package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	exploreURLTemplate = "{+baseurl}/api/now/documents/explore{?page,limit,query,table_name,folder_sys_id,record_sys_id}"
)

// ExploreRequestBuilder provides operations to manage the explore endpoint.
type ExploreRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewExploreRequestBuilderInternal instantiates a new ExploreRequestBuilder.
func NewExploreRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ExploreRequestBuilder {
	return &ExploreRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, exploreURLTemplate, pathParameters),
	}
}

// Get retrieves folder and document metadata with filters, sorting, and pagination.
func (rB *ExploreRequestBuilder) Get(ctx context.Context, requestConfiguration *ExploreRequestBuilderGetRequestConfiguration) (*newInternal.BaseServiceNowCollectionResponse[Document], error) {
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
func (rB *ExploreRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *ExploreRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParams := requestConfiguration.QueryParameters; !internal.IsNil(queryParams) {
			kiotaRequestInfo.AddQueryParameters(queryParams)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
