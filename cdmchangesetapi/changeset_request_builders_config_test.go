package cdmchangesetapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const configTestBaseURL = "https://example.service-now.com"

func configTestPathParameters() map[string]string {
	return map[string]string{"baseurl": configTestBaseURL}
}

// configTestHeaders returns a populated header set, to drive the config.Headers branch.
func configTestHeaders() *abstractions.RequestHeaders {
	headers := abstractions.NewRequestHeaders()
	headers.Add("X-Test", "value")

	return headers
}

// configTestOptions returns a request option ready to be stored on a RequestInformation, to
// drive the config.Options branch.
func configTestOptions() []abstractions.RequestOption {
	option := mocking.NewMockRequestOption()
	option.On("GetKey").Return(abstractions.RequestOptionKey{Key: "test-option"})

	return []abstractions.RequestOption{option}
}

// newConfigTestAdapter returns an adapter that answers Send and SendNoContent successfully, so
// every verb reaches the end of its happy path.
func newConfigTestAdapter(response any) *mocking.MockRequestAdapter {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(response, nil)
	adapter.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	return adapter
}

// TestRequestBuilders_ConfigurationIsApplied covers the inline request-configuration handling in
// every verb of this package. These builders apply Headers, Options and QueryParameters by hand
// rather than through abstractions.ConfigureRequestInformation, so each branch is only reached
// when a config carrying that field is supplied — the existing happy-path tests all pass nil.
func TestRequestBuilders_ConfigurationIsApplied(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name     string
		response func() any
		call     func(adapter *mocking.MockRequestAdapter) error
	}{
		{
			name: "Changesets Get",
			response: func() any {
				return core.NewBaseServiceNowCollectionResponse[*ChangesetResult](CreateChangesetResultFromDiscriminatorValue)
			},
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewChangesetsRequestBuilderInternal(configTestPathParameters(), adapter).
					Get(ctx, &ChangesetsRequestBuilderGetRequestConfiguration{
						Headers:         configTestHeaders(),
						Options:         configTestOptions(),
						QueryParameters: &ChangesetsRequestBuilderGetQueryParameters{},
					})

				return err
			},
		},
		{
			name:     "Changesets Delete",
			response: func() any { return nil },
			call: func(adapter *mocking.MockRequestAdapter) error {
				return NewChangesetsRequestBuilderInternal(configTestPathParameters(), adapter).
					Delete(ctx, &ChangesetsRequestBuilderDeleteRequestConfiguration{
						Headers:         configTestHeaders(),
						Options:         configTestOptions(),
						QueryParameters: &ChangesetsRequestBuilderDeleteQueryParameters{},
					})
			},
		},
		{
			name: "ChangesetActivity Get",
			response: func() any {
				return core.NewBaseServiceNowCollectionResponse[*ChangesetActivityResult](CreateChangesetActivityResultFromDiscriminatorValue)
			},
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewChangesetActivityRequestBuilderInternal(configTestPathParameters(), adapter).
					Get(ctx, &ChangesetActivityRequestBuilderGetRequestConfiguration{
						Headers:         configTestHeaders(),
						Options:         configTestOptions(),
						QueryParameters: &ChangesetActivityRequestBuilderGetQueryParameters{},
					})

				return err
			},
		},
		{
			name: "CommitStatusItem Get",
			response: func() any {
				return core.NewBaseServiceNowItemResponse[*CommitStatusResult](CreateCommitStatusResultFromDiscriminatorValue)
			},
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewCommitStatusItemRequestBuilderInternal(configTestPathParameters(), adapter).
					Get(ctx, &CommitStatusRequestBuilderGetRequestConfiguration{
						Headers: configTestHeaders(),
						Options: configTestOptions(),
					})

				return err
			},
		},
		{
			name: "ImpactedSharedComponents Get",
			response: func() any {
				return core.NewBaseServiceNowCollectionResponse[*ImpactedSharedComponentResult](CreateImpactedSharedComponentResultFromDiscriminatorValue)
			},
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewImpactedSharedComponentsRequestBuilderInternal(configTestPathParameters(), adapter).
					Get(ctx, &ImpactedSharedComponentsRequestBuilderGetRequestConfiguration{
						Headers:         configTestHeaders(),
						Options:         configTestOptions(),
						QueryParameters: &ImpactedSharedComponentsRequestBuilderGetQueryParameters{},
					})

				return err
			},
		},
		{
			name: "ImpactedDeployables Get",
			response: func() any {
				return core.NewBaseServiceNowCollectionResponse[*ImpactedDeployableResult](CreateImpactedDeployableResultFromDiscriminatorValue)
			},
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewImpactedDeployablesRequestBuilderInternal(configTestPathParameters(), adapter).
					Get(ctx, &ImpactedDeployablesRequestBuilderGetRequestConfiguration{
						Headers:         configTestHeaders(),
						Options:         configTestOptions(),
						QueryParameters: &ImpactedDeployablesRequestBuilderGetQueryParameters{},
					})

				return err
			},
		},
		{
			name: "ImpactedDeployablesBySysID Get",
			response: func() any {
				return core.NewBaseServiceNowCollectionResponse[*ImpactedDeployableBySysIDResult](CreateImpactedDeployableBySysIDResultFromDiscriminatorValue)
			},
			call: func(adapter *mocking.MockRequestAdapter) error {
				_, err := NewImpactedDeployablesBySysIDRequestBuilderInternal(configTestPathParameters(), adapter).
					Get(ctx, &ImpactedDeployablesBySysIDRequestBuilderGetRequestConfiguration{
						Headers: configTestHeaders(),
						Options: configTestOptions(),
					})

				return err
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.NoError(t, test.call(newConfigTestAdapter(test.response())))
		})
	}
}
