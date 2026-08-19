//nolint:dupl // per-verb request-builder methods share the mandatory nil-guard/send boilerplate by convention; each depends on its own outer type, response type, and discriminator factory, so it can't be extracted into a shared helper
package cdmchangesetapi

import (
	"context"
	"fmt"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	changesetsURLTemplate                       = "{+baseurl}/api/sn_cdm/changesets{?appName,number,state,changesetNumber}"
	changesetActivityURLTemplate                = "{+baseurl}/api/sn_cdm/changesets/activity{?changesetNumber,returnFields}"
	commitStatusURLTemplate                     = "{+baseurl}/api/sn_cdm/changesets/commit-status/{commit_id}"
	impactedSharedComponentsURLTemplate         = "{+baseurl}/api/sn_cdm/changesets/impacted-shared-components{?changesetNumber,returnFields}"
	impactedDeployablesURLTemplate              = "{+baseurl}/api/sn_cdm/changesets/impacted-deployables{?changesetNumber,returnFields}"
	changesetItemImpactedDeployablesURLTemplate = "{+baseurl}/api/sn_cdm/changesets/{changeset_id}/impacted-deployables"
)

// ChangesetsRequestBuilder handles /changesets endpoint.
type ChangesetsRequestBuilder struct {
	core.RequestBuilder
}

// NewChangesetsRequestBuilderInternal instantiates a new ChangesetsRequestBuilder.
func NewChangesetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ChangesetsRequestBuilder {
	return &ChangesetsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, changesetsURLTemplate, pathParameters),
	}
}

// Activity returns a ChangesetActivityRequestBuilder.
func (rB *ChangesetsRequestBuilder) Activity() *ChangesetActivityRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewChangesetActivityRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// CommitStatus returns a CommitStatusRequestBuilder.
func (rB *ChangesetsRequestBuilder) CommitStatus() *CommitStatusRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewCommitStatusRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ImpactedSharedComponents returns an ImpactedSharedComponentsRequestBuilder.
func (rB *ChangesetsRequestBuilder) ImpactedSharedComponents() *ImpactedSharedComponentsRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewImpactedSharedComponentsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ImpactedDeployables returns an ImpactedDeployablesRequestBuilder.
func (rB *ChangesetsRequestBuilder) ImpactedDeployables() *ImpactedDeployablesRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewImpactedDeployablesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ByID returns a ChangesetItemRequestBuilder.
func (rB *ChangesetsRequestBuilder) ByID(id string) *ChangesetItemRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["changeset_id"] = id
	return NewChangesetItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Get sends a GET request.
func (rB *ChangesetsRequestBuilder) Get(ctx context.Context, config *ChangesetsRequestBuilderGetRequestConfiguration) (ChangesetsResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}
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
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateChangesetsResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.(ChangesetsResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*ChangesetsResponse)(nil))
	}
	return typedRes, nil
}

// Delete sends a DELETE request.
func (rB *ChangesetsRequestBuilder) Delete(ctx context.Context, config *ChangesetsRequestBuilderDeleteRequestConfiguration) error {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return snerrors.ErrNilRequestAdapter
	}
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
	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, core.DefaultErrorMapping())
}

// ChangesetActivityRequestBuilder handles /changesets/activity endpoint.
type ChangesetActivityRequestBuilder struct {
	core.RequestBuilder
}

// NewChangesetActivityRequestBuilderInternal instantiates a new ChangesetActivityRequestBuilder.
func NewChangesetActivityRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ChangesetActivityRequestBuilder {
	return &ChangesetActivityRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, changesetActivityURLTemplate, pathParameters),
	}
}

// Get sends a GET request.
func (rB *ChangesetActivityRequestBuilder) Get(ctx context.Context, config *ChangesetActivityRequestBuilderGetRequestConfiguration) (ChangesetActivityResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}
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
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateChangesetActivityResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.(ChangesetActivityResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*ChangesetActivityResponse)(nil))
	}
	return typedRes, nil
}

// CommitStatusRequestBuilder handles /changesets/commit-status endpoint.
type CommitStatusRequestBuilder struct {
	core.RequestBuilder
}

// NewCommitStatusRequestBuilderInternal instantiates a new CommitStatusRequestBuilder.
func NewCommitStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CommitStatusRequestBuilder {
	return &CommitStatusRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, commitStatusURLTemplate, pathParameters),
	}
}

// ByID returns the by id request builder.
func (rB *CommitStatusRequestBuilder) ByID(commitID string) *CommitStatusItemRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["commit_id"] = commitID
	return NewCommitStatusItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// CommitStatusItemRequestBuilder handles /changesets/commit-status/{commit_id} endpoint.
type CommitStatusItemRequestBuilder struct {
	core.RequestBuilder
}

// NewCommitStatusItemRequestBuilderInternal instantiates a new CommitStatusItemRequestBuilder.
func NewCommitStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CommitStatusItemRequestBuilder {
	return &CommitStatusItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, commitStatusURLTemplate, pathParameters),
	}
}

// Get sends a GET request.
func (rB *CommitStatusItemRequestBuilder) Get(ctx context.Context, config *CommitStatusRequestBuilderGetRequestConfiguration) (CommitStatusResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateCommitStatusResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.(CommitStatusResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*CommitStatusResponse)(nil))
	}
	return typedRes, nil
}

// ImpactedSharedComponentsRequestBuilder handles /changesets/impacted-shared-components endpoint.
type ImpactedSharedComponentsRequestBuilder struct {
	core.RequestBuilder
}

// NewImpactedSharedComponentsRequestBuilderInternal instantiates a new ImpactedSharedComponentsRequestBuilder.
func NewImpactedSharedComponentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ImpactedSharedComponentsRequestBuilder {
	return &ImpactedSharedComponentsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, impactedSharedComponentsURLTemplate, pathParameters),
	}
}

// Get sends a GET request.
func (rB *ImpactedSharedComponentsRequestBuilder) Get(ctx context.Context, config *ImpactedSharedComponentsRequestBuilderGetRequestConfiguration) (ImpactedSharedComponentsResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}
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
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateImpactedSharedComponentsResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.(ImpactedSharedComponentsResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*ImpactedSharedComponentsResponse)(nil))
	}
	return typedRes, nil
}

// ImpactedDeployablesRequestBuilder handles /changesets/impacted-deployables endpoint.
type ImpactedDeployablesRequestBuilder struct {
	core.RequestBuilder
}

// NewImpactedDeployablesRequestBuilderInternal instantiates a new ImpactedDeployablesRequestBuilder.
func NewImpactedDeployablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ImpactedDeployablesRequestBuilder {
	return &ImpactedDeployablesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, impactedDeployablesURLTemplate, pathParameters),
	}
}

// Get sends a GET request.
func (rB *ImpactedDeployablesRequestBuilder) Get(ctx context.Context, config *ImpactedDeployablesRequestBuilderGetRequestConfiguration) (ImpactedDeployablesResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}
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
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateImpactedDeployablesResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.(ImpactedDeployablesResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*ImpactedDeployablesResponse)(nil))
	}
	return typedRes, nil
}

// ChangesetItemRequestBuilder handles /changesets/{changeset_id} endpoint.
type ChangesetItemRequestBuilder struct {
	core.RequestBuilder
}

// NewChangesetItemRequestBuilderInternal instantiates a new ChangesetItemRequestBuilder.
func NewChangesetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ChangesetItemRequestBuilder {
	return &ChangesetItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, changesetItemImpactedDeployablesURLTemplate, pathParameters),
	}
}

// ImpactedDeployables returns an ImpactedDeployablesBySysIDRequestBuilder.
func (rB *ChangesetItemRequestBuilder) ImpactedDeployables() *ImpactedDeployablesBySysIDRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewImpactedDeployablesBySysIDRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ImpactedDeployablesBySysIDRequestBuilder handles /changesets/{changeset_id}/impacted-deployables endpoint.
type ImpactedDeployablesBySysIDRequestBuilder struct {
	core.RequestBuilder
}

// NewImpactedDeployablesBySysIDRequestBuilderInternal instantiates a new ImpactedDeployablesBySysIDRequestBuilder.
func NewImpactedDeployablesBySysIDRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ImpactedDeployablesBySysIDRequestBuilder {
	return &ImpactedDeployablesBySysIDRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, changesetItemImpactedDeployablesURLTemplate, pathParameters),
	}
}

// Get sends a GET request.
func (rB *ImpactedDeployablesBySysIDRequestBuilder) Get(ctx context.Context, config *ImpactedDeployablesBySysIDRequestBuilderGetRequestConfiguration) (ImpactedDeployablesBySysIDResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateImpactedDeployablesBySysIDResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.(ImpactedDeployablesBySysIDResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*ImpactedDeployablesBySysIDResponse)(nil))
	}
	return typedRes, nil
}
