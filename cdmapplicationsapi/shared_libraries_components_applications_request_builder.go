package cdmapplicationsapi

import (
	"context"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const sharedLibrariesComponentsApplicationsURLTemplate = "{+baseurl}/api/sn_cdm/applications/shared_libraries/components/applications{?appName,sharedComponentName,name}"

// SharedLibrariesComponentsApplicationsRequestBuilder provides operations to access shared library component applications.
type SharedLibrariesComponentsApplicationsRequestBuilder struct {
	core.RequestBuilder
}

// NewSharedLibrariesComponentsApplicationsRequestBuilderInternal instantiates a new [SharedLibrariesComponentsApplicationsRequestBuilder].
func NewSharedLibrariesComponentsApplicationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SharedLibrariesComponentsApplicationsRequestBuilder {
	return &SharedLibrariesComponentsApplicationsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, sharedLibrariesComponentsApplicationsURLTemplate, pathParameters),
	}
}

// NewSharedLibrariesComponentsApplicationsRequestBuilder instantiates a new [SharedLibrariesComponentsApplicationsRequestBuilder].
func NewSharedLibrariesComponentsApplicationsRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *SharedLibrariesComponentsApplicationsRequestBuilder {
	return NewSharedLibrariesComponentsApplicationsRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Get gets the collection of shared library component applications.
func (rB *SharedLibrariesComponentsApplicationsRequestBuilder) Get(ctx context.Context, config *SharedLibrariesComponentsApplicationsRequestBuilderGetRequestConfiguration) (SharedLibrariesComponentsApplicationsResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateSharedLibrariesComponentsApplicationsResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.(SharedLibrariesComponentsApplicationsResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*SharedLibrariesComponentsApplicationsResponse)(nil))
	}
	return typedRes, nil
}

// ToGetRequestInformation builds the request information for the Get method.
func (rB *SharedLibrariesComponentsApplicationsRequestBuilder) ToGetRequestInformation(_ context.Context, config *SharedLibrariesComponentsApplicationsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
		if config.QueryParameters != nil {
			requestInfo.AddQueryParameters(*config.QueryParameters)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	return requestInfo, nil
}
