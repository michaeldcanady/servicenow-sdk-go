package cdmchangesetapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestChangesetsRequestBuilder_Get(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewChangesetsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	appName := "test_app"
	config := &ChangesetsRequestBuilderGetRequestConfiguration{
		QueryParameters: &ChangesetsRequestBuilderGetQueryParameters{
			AppName: &appName,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.GetURLTemplate(), builder.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	kiotaRequestInfo.AddQueryParameters(config.QueryParameters)

	uri, _ := requestInfo.GetUri()
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/changesets?appName=test_app", uri.String())
}

func TestChangesetActivityRequestBuilder_Get(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewChangesetActivityRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	changesetNumber := "Chset-1"
	config := &ChangesetActivityRequestBuilderGetRequestConfiguration{
		QueryParameters: &ChangesetActivityRequestBuilderGetQueryParameters{
			ChangesetNumber: &changesetNumber,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.GetURLTemplate(), builder.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	kiotaRequestInfo.AddQueryParameters(config.QueryParameters)

	uri, _ := requestInfo.GetUri()
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/changesets/activity?changesetNumber=Chset-1", uri.String())
}

func TestCommitStatusItemRequestBuilder_Get(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewCommitStatusItemRequestBuilderInternal(map[string]string{
		"baseurl":   "https://example.service-now.com",
		"commit_id": "commit123",
	}, adapter)

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.GetURLTemplate(), builder.GetPathParameters())

	uri, _ := requestInfo.GetUri()
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/changesets/commit-status/commit123", uri.String())
}

func TestChangesetsRequestBuilder_ImpactedSharedComponents(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewChangesetsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	changesetNumber := "Chset-10"
	config := &ImpactedSharedComponentsRequestBuilderGetRequestConfiguration{
		QueryParameters: &ImpactedSharedComponentsRequestBuilderGetQueryParameters{
			ChangesetNumber: &changesetNumber,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.ImpactedSharedComponents().GetURLTemplate(), builder.ImpactedSharedComponents().GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	kiotaRequestInfo.AddQueryParameters(config.QueryParameters)

	uri, _ := requestInfo.GetUri()
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/changesets/impacted-shared-components?changesetNumber=Chset-10", uri.String())
}

func TestChangesetsRequestBuilder_ImpactedDeployables(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewChangesetsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	changesetNumber := "Chset-10"
	config := &ImpactedDeployablesRequestBuilderGetRequestConfiguration{
		QueryParameters: &ImpactedDeployablesRequestBuilderGetQueryParameters{
			ChangesetNumber: &changesetNumber,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.ImpactedDeployables().GetURLTemplate(), builder.ImpactedDeployables().GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	kiotaRequestInfo.AddQueryParameters(config.QueryParameters)

	uri, _ := requestInfo.GetUri()
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/changesets/impacted-deployables?changesetNumber=Chset-10", uri.String())
}

func TestChangesetItemRequestBuilder_ImpactedDeployables(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewChangesetsRequestBuilderInternal(map[string]string{
		"baseurl": "https://example.service-now.com",
	}, adapter)

	changesetID := "sys123"
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.ByID(changesetID).ImpactedDeployables().GetURLTemplate(), builder.ByID(changesetID).ImpactedDeployables().GetPathParameters())

	uri, _ := requestInfo.GetUri()
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/changesets/sys123/impacted-deployables", uri.String())
}
