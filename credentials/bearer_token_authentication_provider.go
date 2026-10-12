// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// BearerOptions contains options for the Bearer token authentication provider.
type BearerOptions struct {
	CommonOptions
}

// BearerTokenAuthenticationProvider is a wrapper around Kiota's BaseBearerTokenAuthenticationProvider
// that implements the Preparable interface.
type BearerTokenAuthenticationProvider struct {
	*authentication.BaseBearerTokenAuthenticationProvider
	tokenProvider         authentication.AccessTokenProvider
	allowedHostsValidator *authentication.AllowedHostsValidator
}

// NewBearerTokenAuthenticationProvider creates a new BearerTokenAuthenticationProvider.
func NewBearerTokenAuthenticationProvider(tokenProvider authentication.AccessTokenProvider, opts BearerOptions) (authentication.AuthenticationProvider, error) {
	return &BearerTokenAuthenticationProvider{
		BaseBearerTokenAuthenticationProvider: authentication.NewBaseBearerTokenAuthenticationProvider(tokenProvider),
		tokenProvider:                         tokenProvider,
		allowedHostsValidator:                 createAllowedHostsValidator(opts.AllowedHosts),
	}, nil
}

// Initialize initializes the underlying token provider if it is Preparable.
func (p *BearerTokenAuthenticationProvider) Initialize(baseURL string) {
	if preparable, ok := p.tokenProvider.(Preparable); ok {
		preparable.Initialize(baseURL)
	}
}

// AuthenticateRequest authenticates the provided RequestInformation.
func (p *BearerTokenAuthenticationProvider) AuthenticateRequest(ctx context.Context, request *abstractions.RequestInformation, additionalAuthenticationContext map[string]interface{}) error {
	// Check allowed hosts if configured
	if p.allowedHostsValidator != nil {
		uri, err := request.GetUri()
		if err != nil {
			return err
		}
		if uri != nil && !p.allowedHostsValidator.IsUrlHostValid(uri) {
			return nil // Host not allowed, skip auth
		}
	}
	return p.BaseBearerTokenAuthenticationProvider.AuthenticateRequest(ctx, request, additionalAuthenticationContext)
}
