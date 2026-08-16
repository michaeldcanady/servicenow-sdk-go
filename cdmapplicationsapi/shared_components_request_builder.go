package cdmapplicationsapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const sharedComponentsURLTemplate = "{+baseurl}/api/sn_cdm/applications/shared_components{?appName,sharedComponentName,name}"

// SharedComponentsRequestBuilder provides operations to manage shared components.
type SharedComponentsRequestBuilder struct {
	core.RequestBuilder
}

// NewSharedComponentsRequestBuilderInternal instantiates a new [SharedComponentsRequestBuilder].
func NewSharedComponentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SharedComponentsRequestBuilder {
	return &SharedComponentsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, sharedComponentsURLTemplate, pathParameters),
	}
}

// NewSharedComponentsRequestBuilder instantiates a new [SharedComponentsRequestBuilder].
func NewSharedComponentsRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *SharedComponentsRequestBuilder {
	return NewSharedComponentsRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Delete deletes shared component references.
func (rB *SharedComponentsRequestBuilder) Delete(ctx context.Context, config *SharedComponentsRequestBuilderDeleteRequestConfiguration) error {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return snerrors.ErrNilRequestAdapter
	}

	requestInfo, err := rB.ToDeleteRequestInformation(ctx, config)
	if err != nil {
		return err
	}

	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, core.DefaultErrorMapping())
}

// Put updates shared components.
func (rB *SharedComponentsRequestBuilder) Put(ctx context.Context, body *SharedComponentUpdateRequest, config *SharedComponentsRequestBuilderPutRequestConfiguration) (SharedComponentUpdateResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	requestInfo, err := rB.ToPutRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateSharedComponentUpdateResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, nil
	}
	typedRes, ok := res.(SharedComponentUpdateResponse)
	if !ok {
		// TODO: standardize error
		return nil, errors.New("unexpected type")
	}
	return typedRes, nil
}

// ToDeleteRequestInformation builds the request information for the Delete method.
func (rB *SharedComponentsRequestBuilder) ToDeleteRequestInformation(_ context.Context, config *SharedComponentsRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
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
	return requestInfo, nil
}

// ToPutRequestInformation builds the request information for the Put method.
func (rB *SharedComponentsRequestBuilder) ToPutRequestInformation(ctx context.Context, body *SharedComponentUpdateRequest, config *SharedComponentsRequestBuilderPutRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}
	return requestInfo, nil
}
