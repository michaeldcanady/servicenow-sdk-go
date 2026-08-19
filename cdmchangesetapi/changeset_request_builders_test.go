package cdmchangesetapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// errNetwork is a stand-in for a transport-level error returned by the adapter.
var errNetwork = errors.New("network error")

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

func TestChangesetsRequestBuilder_Activity(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewChangesetsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	activityBuilder := builder.Activity()
	require.NotNil(t, activityBuilder)
	assert.Equal(t, changesetActivityURLTemplate, activityBuilder.GetURLTemplate())
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

func TestChangesetsRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewChangesetsRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)

	err = builder.Delete(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
}

func TestChangesetActivityRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewChangesetActivityRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestCommitStatusItemRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewCommitStatusItemRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestImpactedSharedComponentsRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewImpactedSharedComponentsRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestImpactedDeployablesRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewImpactedDeployablesRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestImpactedDeployablesBySysIdRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewImpactedDeployablesBySysIDRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestCommitStatusRequestBuilder_ByID(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewCommitStatusRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	itemBuilder := builder.ByID("commit123")
	require.NotNil(t, itemBuilder)
	assert.Equal(t, "commit123", itemBuilder.GetPathParameters()["commit_id"])
	assert.Equal(t, commitStatusURLTemplate, itemBuilder.GetURLTemplate())
}

// ---------------------------------------------------------------------------
// Happy-path / adapter-error coverage for verb methods.
// ---------------------------------------------------------------------------

func TestChangesetsRequestBuilder_Get_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowCollectionResponse[*ChangesetResult](CreateChangesetResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns snerrors.ErrNilResponse",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantErr: snerrors.ErrNilResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewChangesetsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestChangesetsRequestBuilder_Delete_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(errNetwork)
			},
			wantErr: errNetwork,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewChangesetsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			err := builder.Delete(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			adapter.AssertExpectations(t)
		})
	}
}

func TestChangesetActivityRequestBuilder_Get_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowCollectionResponse[*ChangesetActivityResult](CreateChangesetActivityResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns snerrors.ErrNilResponse",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantErr: snerrors.ErrNilResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewChangesetActivityRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestCommitStatusItemRequestBuilder_Get_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*CommitStatusResult](CreateCommitStatusResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns snerrors.ErrNilResponse",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantErr: snerrors.ErrNilResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewChangesetsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter).CommitStatus().ByID("commit123")

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestImpactedSharedComponentsRequestBuilder_Get_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowCollectionResponse[*ImpactedSharedComponentResult](CreateImpactedSharedComponentResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns snerrors.ErrNilResponse",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantErr: snerrors.ErrNilResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewImpactedSharedComponentsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestImpactedDeployablesRequestBuilder_Get_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowCollectionResponse[*ImpactedDeployableResult](CreateImpactedDeployableResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns snerrors.ErrNilResponse",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantErr: snerrors.ErrNilResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewImpactedDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestImpactedDeployablesBySysIDRequestBuilder_Get_HappyAndError(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
	}{
		{
			name: "happy path",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowCollectionResponse[*ImpactedDeployableBySysIDResult](CreateImpactedDeployableBySysIDResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns snerrors.ErrNilResponse",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantErr: snerrors.ErrNilResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)
			builder := NewChangesetsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter).ByID("chg-123").ImpactedDeployables()

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

// isDefaultErrorMapping reports whether the given argument is a non-nil
// abstractions.ErrorMappings with core.DefaultErrorMapping()'s status-code keys.
func isDefaultErrorMapping(v any) bool {
	mapping, ok := v.(abstractions.ErrorMappings)
	if !ok || mapping == nil {
		return false
	}
	for _, code := range []string{"400", "401", "403", "404", "429", "5XX", "XXX"} {
		if _, ok := mapping[code]; !ok {
			return false
		}
	}
	return len(mapping) == 7
}

// TestChangesetsRequestBuilder_Get_PassesDefaultErrorMapping guards against #565:
// CDM builders previously passed literal nil instead of core.DefaultErrorMapping(),
// so ServiceNow API errors never mapped to a typed core.ServiceNowError.
func TestChangesetsRequestBuilder_Get_PassesDefaultErrorMapping(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.MatchedBy(isDefaultErrorMapping)).
		Return(core.NewBaseServiceNowCollectionResponse[*ChangesetResult](CreateChangesetResultFromDiscriminatorValue), nil)

	builder := NewChangesetsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	_, err := builder.Get(context.Background(), nil)

	require.NoError(t, err)
	adapter.AssertExpectations(t)
}
