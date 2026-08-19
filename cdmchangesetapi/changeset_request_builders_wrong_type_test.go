package cdmchangesetapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// These tests exercise the comma-ok type-assertion guard: when the adapter
// returns a non-nil value whose concrete type doesn't match the expected
// response type, the verb method must return an error instead of panicking.

func TestChangesetsRequestBuilder_Get_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewChangesetsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Get(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestChangesetActivityRequestBuilder_Get_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewChangesetActivityRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Get(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestCommitStatusItemRequestBuilder_Get_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewCommitStatusItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "commit_id": "commit123"}, adapter)

	resp, err := builder.Get(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestImpactedSharedComponentsRequestBuilder_Get_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewImpactedSharedComponentsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Get(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestImpactedDeployablesRequestBuilder_Get_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewImpactedDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Get(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestImpactedDeployablesBySysIDRequestBuilder_Get_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewImpactedDeployablesBySysIDRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "changeset_id": "changeset123"}, adapter)

	resp, err := builder.Get(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}
