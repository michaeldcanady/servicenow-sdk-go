package documentsapi

import (
	"context"
	"fmt"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
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
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
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
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := res.(*core.BaseServiceNowCollectionResponse[Document])
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*core.BaseServiceNowCollectionResponse[Document])(nil))
	}

	return typedResp, nil
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *ExploreRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *ExploreRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(requestConfiguration) {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*requestConfiguration.QueryParameters)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
