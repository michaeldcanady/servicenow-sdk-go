// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package support

import (
	"context"
	"net/url"

	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// StaticTokenProvider returns a fixed token. Used in bearer token and JWT tests.
type StaticTokenProvider struct {
	Token string
	Err   error
}

func (p *StaticTokenProvider) GetAuthorizationToken(_ context.Context, _ *url.URL, _ map[string]interface{}) (string, error) {
	if p.Err != nil {
		return "", p.Err
	}
	return p.Token, nil
}

func (p *StaticTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}

// FailingTokenProvider always returns an error. Used to test JWT provider error handling.
type FailingTokenProvider struct {
	Err error
}

func (p *FailingTokenProvider) GetAuthorizationToken(_ context.Context, _ *url.URL, _ map[string]interface{}) (string, error) {
	return "", p.Err
}

func (p *FailingTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}
