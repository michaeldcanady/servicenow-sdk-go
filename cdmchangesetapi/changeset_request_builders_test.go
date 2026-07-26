package cdmchangesetapi

import (
	"context"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	requestInfo.AddQueryParameters(*config.QueryParameters)

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
	requestInfo.AddQueryParameters(*config.QueryParameters)

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
	requestInfo.AddQueryParameters(*config.QueryParameters)

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
	requestInfo.AddQueryParameters(*config.QueryParameters)

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

func TestChangesetsRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*ChangesetsRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)

			err = builder.Delete(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		})
	}
}

func TestChangesetActivityRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*ChangesetActivityRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestCommitStatusItemRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*CommitStatusItemRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestImpactedSharedComponentsRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*ImpactedSharedComponentsRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestImpactedDeployablesRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*ImpactedDeployablesRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestImpactedDeployablesBySysIdRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*ImpactedDeployablesBySysIDRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}
