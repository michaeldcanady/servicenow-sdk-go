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
