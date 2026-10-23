// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"context"
	"strings"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// BearerConfig configures a fixed Bearer token.
type BearerConfig struct {
	ProviderConfig
	// Token is the fixed access token.
	Token string
}

// Validate checks the Bearer token settings.
func (c BearerConfig) Validate() error {
	if err := c.ProviderConfig.Validate(); err != nil {
		return err
	}
	if strings.TrimSpace(c.Token) == "" {
		return snerrors.ErrEmptyToken
	}
	return nil
}

// BearerProvider adds a fixed Bearer token to each request.
type BearerProvider struct {
	token     string
	validator *authentication.AllowedHostsValidator
}

// NewBearerProvider creates a provider for a fixed Bearer token.
func NewBearerProvider(config BearerConfig) (*BearerProvider, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	validator, err := newAllowedHostsValidator(config.ProviderConfig)
	if err != nil {
		return nil, err
	}

	return &BearerProvider{
		token:     config.Token,
		validator: validator,
	}, nil
}

// Initialize is a no-op for a fixed Bearer token.
func (p *BearerProvider) Initialize(baseURL string) error {
	if p == nil {
		return snerrors.ErrNilProvider
	}
	return nil
}

// AuthenticateRequest adds the fixed Bearer token.
func (p *BearerProvider) AuthenticateRequest(ctx context.Context, request *abstractions.RequestInformation, _ map[string]interface{}) error {
	if p == nil {
		return snerrors.ErrNilProvider
	}

	_, allowed, err := requestURI(ctx, request, p.validator)
	if err != nil || !allowed {
		return err
	}

	setHeaderIfAbsent(request, "Authorization", "Bearer "+p.token)
	return nil
}
