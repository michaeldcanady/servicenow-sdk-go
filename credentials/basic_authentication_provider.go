// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// BasicOptions contains options for the Basic authentication provider.
type BasicOptions struct {
	CommonOptions
}

// BasicAuthenticationProvider is an authentication.AuthenticationProvider implementing HTTP
// Basic authentication with a static username/password pair.
type BasicAuthenticationProvider struct {
	username              string
	password              string
	allowedHostsValidator *authentication.AllowedHostsValidator
}

// NewBasicProvider creates a new BasicAuthenticationProvider with positional
// credentials and shared options.
func NewBasicProvider(username, password string, opts ...BasicOptions) *BasicAuthenticationProvider {
	var o BasicOptions
	if len(opts) > 0 {
		o = opts[0]
	}
	return &BasicAuthenticationProvider{
		username:              username,
		password:              password,
		allowedHostsValidator: createAllowedHostsValidator(o.AllowedHosts),
	}
}

// AuthenticateRequest authenticates the provided RequestInformation.
func (b *BasicAuthenticationProvider) AuthenticateRequest(context context.Context, request *abstractions.RequestInformation, _ map[string]interface{}) error {
	if conversion.IsNil(b) {
		return errors.New("provider is nil")
	}

	if err := context.Err(); err != nil {
		return err
	}

	if strings.TrimSpace(b.username) == "" {
		return snerrors.ErrEmptyUsername
	}

	if strings.TrimSpace(b.password) == "" {
		return snerrors.ErrEmptyPassword
	}

	// Check allowed hosts if configured
	if b.allowedHostsValidator != nil {
		uri, err := request.GetUri()
		if err != nil {
			return err
		}
		if uri != nil && !b.allowedHostsValidator.IsUrlHostValid(uri) {
			return nil // Host not allowed, skip auth
		}
	}

	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", b.username, b.password)))

	request.Headers.Add("Authorization", fmt.Sprintf("Basic %s", auth))

	return nil
}

// createAllowedHostsValidator creates a host validator from the allowed hosts list.
func createAllowedHostsValidator(hosts []string) *authentication.AllowedHostsValidator {
	if len(hosts) > 0 {
		validator, _ := authentication.NewAllowedHostsValidatorErrorCheck(hosts)
		return validator
	}
	return nil
}
