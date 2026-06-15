package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	exploreURLTemplate = "{+baseurl}/api/now/v1/documents/explore{?page,limit,query,table_name,folder_sys_id,record_sys_id}"
)

// ExploreRequestBuilder provides operations to manage the explore endpoint.
type ExploreRequestBuilder struct {
	internal.RequestBuilder
}

// NewExploreRequestBuilderInternal instantiates a new ExploreRequestBuilder.
func NewExploreRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ExploreRequestBuilder {
	return &ExploreRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, exploreURLTemplate, pathParameters),
	}
}

// Get retrieves folder and document metadata with filters, sorting, and pagination.
func (rB *ExploreRequestBuilder) Get(ctx context.Context, requestConfiguration *ExploreRequestBuilderGetRequestConfiguration) (*internal.BaseServiceNowCollectionResponse[Document], error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowCollectionResponseFromDiscriminatorValue[Document](CreateDocumentFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(*internal.BaseServiceNowCollectionResponse[Document]), nil
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *ExploreRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *ExploreRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParams := requestConfiguration.QueryParameters; !conversion.IsNil(queryParams) {
			kiotaRequestInfo.AddQueryParameters(queryParams)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
