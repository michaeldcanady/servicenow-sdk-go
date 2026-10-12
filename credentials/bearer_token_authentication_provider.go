// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// BearerTokenAuthenticationProvider is a wrapper around Kiota's BaseBearerTokenAuthenticationProvider
// that implements the Preparable interface.
type BearerTokenAuthenticationProvider struct {
	*authentication.BaseBearerTokenAuthenticationProvider
	tokenProvider authentication.AccessTokenProvider
}

// NewBearerTokenAuthenticationProvider creates a new BearerTokenAuthenticationProvider.
func NewBearerTokenAuthenticationProvider(tokenProvider authentication.AccessTokenProvider) *BearerTokenAuthenticationProvider {
	return &BearerTokenAuthenticationProvider{
		BaseBearerTokenAuthenticationProvider: authentication.NewBaseBearerTokenAuthenticationProvider(tokenProvider),
		tokenProvider:                         tokenProvider,
	}
}

// Initialize initializes the underlying token provider if it is Preparable.
func (p *BearerTokenAuthenticationProvider) Initialize(baseURL string) {
	if preparable, ok := p.tokenProvider.(Preparable); ok {
		preparable.Initialize(baseURL)
	}
}

// Revoke revokes the current token through the underlying token provider. It is
// a no-op success when the provider does not support revocation.
func (p *BearerTokenAuthenticationProvider) Revoke(ctx context.Context) error {
	if revoker, ok := p.tokenProvider.(Revoker); ok {
		return revoker.Revoke(ctx)
	}

	return nil
}

// AuthenticateRequest authenticates the provided RequestInformation.
func (p *BearerTokenAuthenticationProvider) AuthenticateRequest(ctx context.Context, request *abstractions.RequestInformation, additionalAuthenticationContext map[string]interface{}) error {
	return p.BaseBearerTokenAuthenticationProvider.AuthenticateRequest(ctx, request, additionalAuthenticationContext)
}
