package appserviceapi

import (
	"context"
	"errors"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const findServiceURLTemplate = "{+baseurl}/api/now/cmdb/csdm/app_service/find_service{?name,number}"

// FindServiceRequestBuilder provides operations to find an application service.
type FindServiceRequestBuilder struct {
	core.RequestBuilder
}

// NewFindServiceRequestBuilderInternal instantiates a new [FindServiceRequestBuilder].
func NewFindServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FindServiceRequestBuilder {
	return &FindServiceRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, findServiceURLTemplate, pathParameters),
	}
}

// NewFindServiceRequestBuilder instantiates a new [FindServiceRequestBuilder].
func NewFindServiceRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *FindServiceRequestBuilder {
	return NewFindServiceRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Get sends a GET request to find an application service.
func (rB *FindServiceRequestBuilder) Get(ctx context.Context, config *FindServiceRequestConfiguration) (FindServiceResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateFindServiceResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	typedResp, ok := res.(FindServiceResponse)
	if !ok {
		// TODO: standardize error
		return nil, errors.New("unexpected type")
	}

	return typedResp, nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *FindServiceRequestBuilder) ToGetRequestInformation(_ context.Context, config *FindServiceRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	// TODO: check for nil/empty template and path parameters
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	return requestInfo, nil
}
