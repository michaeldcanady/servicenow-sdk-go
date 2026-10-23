// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"context"
	"slices"
	"strings"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// DefaultAPIKeyParameter is the ServiceNow API key header name.
const DefaultAPIKeyParameter = "x-sn-apikey" //nolint:gosec // Header name, not a credential value.

// APIKeyConfig configures API key authentication.
type APIKeyConfig struct {
	ProviderConfig
	// Key is the API key value.
	Key string
	// ParameterName is the header or query parameter name. An empty value uses
	// DefaultAPIKeyParameter.
	ParameterName string
	// Location is where the key is sent. The default is a header.
	Location AuthLocation
	// Prefix is placed verbatim before the key. An empty prefix sends the key
	// unchanged.
	Prefix string
}

// Validate checks the API key settings.
func (c APIKeyConfig) Validate() error {
	if err := c.ProviderConfig.Validate(); err != nil {
		return err
	}
	if strings.TrimSpace(c.Key) == "" {
		return snerrors.ErrEmptyAPIKey
	}
	if !slices.Contains([]AuthLocation{AuthLocationHeader, AuthLocationQuery}, c.Location) {
		return snerrors.ErrInvalidAuthLocation
	}
	return nil
}

// APIKeyProvider adds an API key to each request.
type APIKeyProvider struct {
	key           string
	parameterName string
	location      AuthLocation
	prefix        string
	validator     *authentication.AllowedHostsValidator
}

// NewAPIKeyProvider creates an API key authentication provider.
func NewAPIKeyProvider(config APIKeyConfig) (*APIKeyProvider, error) {
	if strings.TrimSpace(config.ParameterName) == "" {
		config.ParameterName = DefaultAPIKeyParameter
	}
	if err := config.Validate(); err != nil {
		return nil, err
	}

	validator, err := newAllowedHostsValidator(config.ProviderConfig)
	if err != nil {
		return nil, err
	}

	return &APIKeyProvider{
		key:           config.Key,
		parameterName: config.ParameterName,
		location:      config.Location,
		prefix:        config.Prefix,
		validator:     validator,
	}, nil
}

// Initialize is a no-op for API key authentication.
func (p *APIKeyProvider) Initialize(baseURL string) error {
	if p == nil {
		return snerrors.ErrNilProvider
	}
	return nil
}

// AuthenticateRequest adds the API key as a header or query parameter.
func (p *APIKeyProvider) AuthenticateRequest(ctx context.Context, request *abstractions.RequestInformation, _ map[string]interface{}) error {
	if p == nil {
		return snerrors.ErrNilProvider
	}

	uri, allowed, err := requestURI(ctx, request, p.validator)
	if err != nil || !allowed {
		return err
	}

	value := p.prefix + p.key
	switch p.location {
	case AuthLocationHeader:
		setHeaderIfAbsent(request, p.parameterName, value)
	case AuthLocationQuery:
		query := uri.Query()
		query.Set(p.parameterName, value)
		uri.RawQuery = query.Encode()
		request.SetUri(*uri)
	default:
		return snerrors.ErrInvalidAuthLocation
	}
	return nil
}
