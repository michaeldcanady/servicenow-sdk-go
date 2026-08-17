package cdmapplicationsapi

import (
	"context"
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const deployablesURLTemplate = "{+baseurl}/api/sn_cdm/applications/deployables{?appName,name}"

// DeployablesRequestBuilder provides operations to manage deployables.
type DeployablesRequestBuilder struct {
	core.RequestBuilder
}

// NewDeployablesRequestBuilderInternal instantiates a new [DeployablesRequestBuilder].
func NewDeployablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *DeployablesRequestBuilder {
	return &DeployablesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, deployablesURLTemplate, pathParameters),
	}
}

// NewDeployablesRequestBuilder instantiates a new [DeployablesRequestBuilder].
func NewDeployablesRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *DeployablesRequestBuilder {
	return NewDeployablesRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Delete deletes a deployable.
func (rB *DeployablesRequestBuilder) Delete(ctx context.Context, config *DeployablesRequestBuilderDeleteRequestConfiguration) error {
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

// Put updates deployables.
func (rB *DeployablesRequestBuilder) Put(ctx context.Context, config *DeployablesRequestBuilderPutRequestConfiguration) (DeployableUpdateResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	requestInfo, err := rB.ToPutRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateDeployableUpdateResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, nil
	}
	typedRes, ok := res.(DeployableUpdateResponse)
	if !ok {
		// TODO: standardize error
		return nil, errors.New("unexpected type")
	}
	return typedRes, nil
}

func (rB *DeployablesRequestBuilder) ToPutRequestInformation(ctx context.Context, config *DeployablesRequestBuilderPutRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
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

func (rB *DeployablesRequestBuilder) ToDeleteRequestInformation(ctx context.Context, config *DeployablesRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
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

// Exports returns an [ExportsRequestBuilder].
func (rB *DeployablesRequestBuilder) Exports() *ExportsRequestBuilder {
	return NewExportsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
