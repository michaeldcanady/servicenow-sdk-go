package cdmchangesetapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
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

func NewChangesetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ChangesetsRequestBuilder {
	return &ChangesetsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, changesetsURLTemplate, pathParameters),
	}
}

// Activity returns a ChangesetActivityRequestBuilder.
func (rB *ChangesetsRequestBuilder) Activity() *ChangesetActivityRequestBuilder {
	return NewChangesetActivityRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// CommitStatus returns a CommitStatusRequestBuilder.
func (rB *ChangesetsRequestBuilder) CommitStatus() *CommitStatusRequestBuilder {
	return NewCommitStatusRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ImpactedSharedComponents returns an ImpactedSharedComponentsRequestBuilder.
func (rB *ChangesetsRequestBuilder) ImpactedSharedComponents() *ImpactedSharedComponentsRequestBuilder {
	return NewImpactedSharedComponentsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ImpactedDeployables returns an ImpactedDeployablesRequestBuilder.
func (rB *ChangesetsRequestBuilder) ImpactedDeployables() *ImpactedDeployablesRequestBuilder {
	return NewImpactedDeployablesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ByID returns a ChangesetItemRequestBuilder.
func (rB *ChangesetsRequestBuilder) ByID(id string) *ChangesetItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["changeset_id"] = id
	return NewChangesetItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

func (rB *ChangesetsRequestBuilder) Get(ctx context.Context, config *ChangesetsRequestBuilderGetRequestConfiguration) (ChangesetsResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
		if config.QueryParameters != nil {
			kiotaRequestInfo.AddQueryParameters(config.QueryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateChangesetsResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(ChangesetsResponse), nil
}

func (rB *ChangesetsRequestBuilder) Delete(ctx context.Context, config *ChangesetsRequestBuilderDeleteRequestConfiguration) error {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
		if config.QueryParameters != nil {
			kiotaRequestInfo.AddQueryParameters(config.QueryParameters)
		}
	}
	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, nil)
}

// ChangesetActivityRequestBuilder handles /changesets/activity endpoint.
type ChangesetActivityRequestBuilder struct {
	core.RequestBuilder
}

func NewChangesetActivityRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ChangesetActivityRequestBuilder {
	return &ChangesetActivityRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, changesetActivityURLTemplate, pathParameters),
	}
}

func (rB *ChangesetActivityRequestBuilder) Get(ctx context.Context, config *ChangesetActivityRequestBuilderGetRequestConfiguration) (ChangesetActivityResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
		if config.QueryParameters != nil {
			kiotaRequestInfo.AddQueryParameters(config.QueryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateChangesetActivityResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(ChangesetActivityResponse), nil
}

// CommitStatusRequestBuilder handles /changesets/commit-status endpoint.
type CommitStatusRequestBuilder struct {
	core.RequestBuilder
}

func NewCommitStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CommitStatusRequestBuilder {
	return &CommitStatusRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, commitStatusURLTemplate, pathParameters),
	}
}

func (rB *CommitStatusRequestBuilder) ByID(commitId string) *CommitStatusItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["commit_id"] = commitId
	return NewCommitStatusItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// CommitStatusItemRequestBuilder handles /changesets/commit-status/{commit_id} endpoint.
type CommitStatusItemRequestBuilder struct {
	core.RequestBuilder
}

func NewCommitStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CommitStatusItemRequestBuilder {
	return &CommitStatusItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, commitStatusURLTemplate, pathParameters),
	}
}

func (rB *CommitStatusItemRequestBuilder) Get(ctx context.Context, config *CommitStatusRequestBuilderGetRequestConfiguration) (CommitStatusResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateCommitStatusResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(CommitStatusResponse), nil
}

// ImpactedSharedComponentsRequestBuilder handles /changesets/impacted-shared-components endpoint.
type ImpactedSharedComponentsRequestBuilder struct {
	core.RequestBuilder
}

func NewImpactedSharedComponentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ImpactedSharedComponentsRequestBuilder {
	return &ImpactedSharedComponentsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, impactedSharedComponentsURLTemplate, pathParameters),
	}
}

func (rB *ImpactedSharedComponentsRequestBuilder) Get(ctx context.Context, config *ImpactedSharedComponentsRequestBuilderGetRequestConfiguration) (ImpactedSharedComponentsResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
		if config.QueryParameters != nil {
			kiotaRequestInfo.AddQueryParameters(config.QueryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateImpactedSharedComponentsResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(ImpactedSharedComponentsResponse), nil
}

// ImpactedDeployablesRequestBuilder handles /changesets/impacted-deployables endpoint.
type ImpactedDeployablesRequestBuilder struct {
	core.RequestBuilder
}

func NewImpactedDeployablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ImpactedDeployablesRequestBuilder {
	return &ImpactedDeployablesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, impactedDeployablesURLTemplate, pathParameters),
	}
}

func (rB *ImpactedDeployablesRequestBuilder) Get(ctx context.Context, config *ImpactedDeployablesRequestBuilderGetRequestConfiguration) (ImpactedDeployablesResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
		if config.QueryParameters != nil {
			kiotaRequestInfo.AddQueryParameters(config.QueryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateImpactedDeployablesResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(ImpactedDeployablesResponse), nil
}

// ChangesetItemRequestBuilder handles /changesets/{changeset_id} endpoint.
type ChangesetItemRequestBuilder struct {
	core.RequestBuilder
}

func NewChangesetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ChangesetItemRequestBuilder {
	return &ChangesetItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, changesetItemImpactedDeployablesURLTemplate, pathParameters),
	}
}

// ImpactedDeployables returns an ImpactedDeployablesBySysIdRequestBuilder.
func (rB *ChangesetItemRequestBuilder) ImpactedDeployables() *ImpactedDeployablesBySysIdRequestBuilder {
	return NewImpactedDeployablesBySysIdRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ImpactedDeployablesBySysIdRequestBuilder handles /changesets/{changeset_id}/impacted-deployables endpoint.
type ImpactedDeployablesBySysIdRequestBuilder struct {
	core.RequestBuilder
}

func NewImpactedDeployablesBySysIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ImpactedDeployablesBySysIdRequestBuilder {
	return &ImpactedDeployablesBySysIdRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, changesetItemImpactedDeployablesURLTemplate, pathParameters),
	}
}

func (rB *ImpactedDeployablesBySysIdRequestBuilder) Get(ctx context.Context, config *ImpactedDeployablesBySysIdRequestBuilderGetRequestConfiguration) (ImpactedDeployablesBySysIdResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateImpactedDeployablesBySysIdResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(ImpactedDeployablesBySysIdResponse), nil
}
