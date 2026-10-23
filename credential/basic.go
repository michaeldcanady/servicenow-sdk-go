// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"context"
	"encoding/base64"
	"strings"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// BasicConfig configures HTTP Basic authentication.
type BasicConfig struct {
	ProviderConfig
	// Username is the ServiceNow user name.
	Username string
	// Password is the ServiceNow password.
	Password string
}

// Validate checks the Basic authentication settings.
func (c BasicConfig) Validate() error {
	if err := c.ProviderConfig.Validate(); err != nil {
		return err
	}
	if strings.TrimSpace(c.Username) == "" {
		return snerrors.ErrEmptyUsername
	}
	if strings.TrimSpace(c.Password) == "" {
		return snerrors.ErrEmptyPassword
	}
	return nil
}

// BasicProvider adds an HTTP Basic authorization header to each request.
type BasicProvider struct {
	username  string
	password  string
	validator *authentication.AllowedHostsValidator
}

// NewBasicProvider creates a Basic authentication provider.
func NewBasicProvider(config BasicConfig) (*BasicProvider, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	validator, err := newAllowedHostsValidator(config.ProviderConfig)
	if err != nil {
		return nil, err
	}

	return &BasicProvider{
		username:  config.Username,
		password:  config.Password,
		validator: validator,
	}, nil
}

// Initialize is a no-op for Basic authentication.
func (p *BasicProvider) Initialize(baseURL string) error {
	if p == nil {
		return snerrors.ErrNilProvider
	}
	return nil
}

// AuthenticateRequest adds the HTTP Basic authorization header.
func (p *BasicProvider) AuthenticateRequest(ctx context.Context, request *abstractions.RequestInformation, _ map[string]interface{}) error {
	if p == nil {
		return snerrors.ErrNilProvider
	}

	_, allowed, err := requestURI(ctx, request, p.validator)
	if err != nil || !allowed {
		return err
	}

	credentials := base64.StdEncoding.EncodeToString([]byte(p.username + ":" + p.password))
	setHeaderIfAbsent(request, "Authorization", "Basic "+credentials)
	return nil
}
