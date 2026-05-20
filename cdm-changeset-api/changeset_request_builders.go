package cdmchangesetapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	changesetsURLTemplate        = "{+baseurl}/api/sn_cdm/changesets{?appName,number,state,changesetNumber}"
	changesetActivityURLTemplate = "{+baseurl}/api/sn_cdm/changesets/activity{?changesetNumber,returnFields}"
	commitStatusURLTemplate      = "{+baseurl}/api/sn_cdm/changesets/commit-status/{commit_id}"
)

// ChangesetsRequestBuilder handles /changesets endpoint.
type ChangesetsRequestBuilder struct {
	newInternal.RequestBuilder
}

func NewChangesetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ChangesetsRequestBuilder {
	return &ChangesetsRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, changesetsURLTemplate, pathParameters),
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

func (rB *ChangesetsRequestBuilder) Get(ctx context.Context, config *ChangesetsRequestBuilderGetRequestConfiguration) (ChangesetsResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
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
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateChangesetsResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(ChangesetsResponse), nil
}

func (rB *ChangesetsRequestBuilder) Delete(ctx context.Context, config *ChangesetsRequestBuilderDeleteRequestConfiguration) error {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
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
	newInternal.RequestBuilder
}

func NewChangesetActivityRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ChangesetActivityRequestBuilder {
	return &ChangesetActivityRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, changesetActivityURLTemplate, pathParameters),
	}
}

func (rB *ChangesetActivityRequestBuilder) Get(ctx context.Context, config *ChangesetActivityRequestBuilderGetRequestConfiguration) (ChangesetActivityResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
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
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateChangesetActivityResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(ChangesetActivityResponse), nil
}

// CommitStatusRequestBuilder handles /changesets/commit-status endpoint.
type CommitStatusRequestBuilder struct {
	newInternal.RequestBuilder
}

func NewCommitStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CommitStatusRequestBuilder {
	return &CommitStatusRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, commitStatusURLTemplate, pathParameters),
	}
}

func (rB *CommitStatusRequestBuilder) ByID(commitId string) *CommitStatusItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["commit_id"] = commitId
	return NewCommitStatusItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// CommitStatusItemRequestBuilder handles /changesets/commit-status/{commit_id} endpoint.
type CommitStatusItemRequestBuilder struct {
	newInternal.RequestBuilder
}

func NewCommitStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CommitStatusItemRequestBuilder {
	return &CommitStatusItemRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, commitStatusURLTemplate, pathParameters),
	}
}

func (rB *CommitStatusItemRequestBuilder) Get(ctx context.Context, config *CommitStatusRequestBuilderGetRequestConfiguration) (CommitStatusResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateCommitStatusResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(CommitStatusResponse), nil
}
