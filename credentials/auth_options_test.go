// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonOptions(t *testing.T) {
	opts := CommonOptions{
		AllowedHosts: []string{"dev12345.service-now.com"},
		HTTPClient:   &http.Client{},
		TokenStore:   &optionsFakeTokenStore{},
		Scopes:       []string{"read", "write"},
	}

	assert.Equal(t, []string{"dev12345.service-now.com"}, opts.AllowedHosts)
	assert.NotNil(t, opts.HTTPClient)
	assert.NotNil(t, opts.TokenStore)
	assert.Equal(t, []string{"read", "write"}, opts.Scopes)
}

func TestClientCredentialsOptions(t *testing.T) {
	opts := ClientCredentialsOptions{
		CommonOptions: CommonOptions{
			AllowedHosts: []string{"dev12345.service-now.com"},
			Scopes:       []string{"read"},
		},
	}

	assert.Equal(t, []string{"dev12345.service-now.com"}, opts.AllowedHosts)
	assert.Equal(t, []string{"read"}, opts.Scopes)
}

func TestJWTOptions(t *testing.T) {
	opts := JWTOptions{
		CommonOptions: CommonOptions{
			AllowedHosts: []string{"dev12345.service-now.com"},
		},
		TokenProvider: nil,
	}

	assert.Equal(t, []string{"dev12345.service-now.com"}, opts.AllowedHosts)
	assert.Nil(t, opts.TokenProvider)
}

func TestAuthorizationCodeOptions(t *testing.T) {
	opts := AuthorizationCodeOptions{
		CommonOptions: CommonOptions{
			AllowedHosts: []string{"dev12345.service-now.com"},
		},
		Port:           5001,
		StateGenerator: func() string { return "state" },
		URLOpener:      func(string) error { return nil },
		ServerFactory:  nil,
	}

	assert.Equal(t, []string{"dev12345.service-now.com"}, opts.AllowedHosts)
	assert.Equal(t, 5001, opts.Port)
	assert.NotNil(t, opts.StateGenerator)
	assert.NotNil(t, opts.URLOpener)
	assert.Nil(t, opts.ServerFactory)
}

func TestROPCOptions(t *testing.T) {
	opts := ROPCOptions{
		CommonOptions: CommonOptions{
			AllowedHosts: []string{"dev12345.service-now.com"},
		},
	}

	assert.Equal(t, []string{"dev12345.service-now.com"}, opts.AllowedHosts)
}

// TODO: should be a Mock token store
type optionsFakeTokenStore struct{}

func (f *optionsFakeTokenStore) Save(_ context.Context, _ TokenKey, _ *AccessToken) error {
	return nil
}

func (f *optionsFakeTokenStore) Load(_ context.Context, _ TokenKey) (*AccessToken, error) {
	return nil, nil
}

func (f *optionsFakeTokenStore) Delete(_ context.Context, _ TokenKey) error {
	return nil
}
