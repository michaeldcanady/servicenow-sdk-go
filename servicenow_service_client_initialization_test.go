// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package servicenowsdkgo

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/credential"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type errorInitializingProvider struct {
	baseURL string
	err     error
}

func (p *errorInitializingProvider) AuthenticateRequest(context.Context, *abstractions.RequestInformation, map[string]interface{}) error {
	return nil
}

func (p *errorInitializingProvider) Initialize(baseURL string) error {
	p.baseURL = baseURL
	return p.err
}

type legacyInitializingProvider struct {
	baseURL string
}

func (p *legacyInitializingProvider) AuthenticateRequest(context.Context, *abstractions.RequestInformation, map[string]interface{}) error {
	return nil
}

func (p *legacyInitializingProvider) Initialize(baseURL string) {
	p.baseURL = baseURL
}

func newInitializationAdapter() *mocking.MockRequestAdapter {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("EnableBackingStore", mock.Anything).Return()
	adapter.On("SetBaseUrl", mock.Anything).Return()
	return adapter
}

func TestClientInitializesErrorReturningProvider(t *testing.T) {
	provider := &errorInitializingProvider{}
	adapter := newInitializationAdapter()

	client, err := NewServiceNowServiceClient(
		WithURL("https://example.com"),
		WithAuthenticationProvider(provider),
		WithRequestAdapter(adapter),
	)
	require.NoError(t, err)
	require.NotNil(t, client)
	assert.Equal(t, "https://example.com", provider.baseURL)
	adapter.AssertExpectations(t)
}

func TestClientInitializesLegacyProvider(t *testing.T) {
	provider := &legacyInitializingProvider{}
	adapter := newInitializationAdapter()

	client, err := NewServiceNowServiceClient(
		WithURL("https://example.com"),
		WithAuthenticationProvider(provider),
		WithRequestAdapter(adapter),
	)
	require.NoError(t, err)
	require.NotNil(t, client)
	assert.Equal(t, "https://example.com", provider.baseURL)
	adapter.AssertExpectations(t)
}

func TestClientPropagatesInitializationError(t *testing.T) {
	expected := errors.New("initialization failed")
	provider := &errorInitializingProvider{err: expected}
	adapter := newInitializationAdapter()

	client, err := NewServiceNowServiceClient(
		WithURL("https://example.com"),
		WithAuthenticationProvider(provider),
		WithRequestAdapter(adapter),
	)
	assert.Nil(t, client)
	require.ErrorIs(t, err, expected)
	adapter.AssertExpectations(t)
}

func TestClientDoesNotInitializeWithoutABaseURL(t *testing.T) {
	provider := &errorInitializingProvider{}
	adapter := newInitializationAdapter()

	client, err := NewServiceNowServiceClient(
		WithAuthenticationProvider(provider),
		WithRequestAdapter(adapter),
	)
	assert.Nil(t, client)
	require.ErrorIs(t, err, snerrors.ErrEmptyURI)
	assert.Empty(t, provider.baseURL)
}

func TestClientAcceptsDirectCredentialProvider(t *testing.T) {
	provider, err := credential.NewBasicProvider(credential.BasicConfig{
		Username: "user",
		Password: "password",
	})
	require.NoError(t, err)
	adapter := newInitializationAdapter()

	client, err := NewServiceNowServiceClient(
		WithURL("https://example.com"),
		WithAuthenticationProvider(provider),
		WithRequestAdapter(adapter),
	)
	require.NoError(t, err)
	require.NotNil(t, client)
	adapter.AssertExpectations(t)
}
