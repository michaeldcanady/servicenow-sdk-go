// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithInstance(t *testing.T) {
	tests := []struct {
		name        string
		instance    string
		wantBaseURL string
	}{
		{
			name:        "builds https url from instance name",
			instance:    "dev12345",
			wantBaseURL: "https://dev12345.service-now.com",
		},
		{
			name:        "empty instance still builds a url",
			instance:    "",
			wantBaseURL: "https://.service-now.com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := &AuthConfig{}
			WithInstance(test.instance)(config)

			assert.Equal(t, test.wantBaseURL, config.baseURL)
		})
	}
}

func TestWithAllowedHosts(t *testing.T) {
	tests := []struct {
		name  string
		hosts []string
	}{
		{
			name:  "sets a single host",
			hosts: []string{"dev12345.service-now.com"},
		},
		{
			name:  "sets multiple hosts",
			hosts: []string{"a.example.com", "b.example.com"},
		},
		{
			name:  "no hosts clears to empty slice",
			hosts: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := &AuthConfig{}
			WithAllowedHosts(test.hosts...)(config)

			assert.Equal(t, test.hosts, config.allowedHosts)
		})
	}
}

func TestWithTokenStore(t *testing.T) {
	tests := []struct {
		name  string
		store TokenStore
	}{
		{
			name:  "sets a non-nil token store",
			store: &optionsFakeTokenStore{},
		},
		{
			name:  "nil token store is accepted",
			store: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := &AuthConfig{}
			WithTokenStore(test.store)(config)

			assert.Equal(t, test.store, config.tokenStore)
		})
	}
}

func TestWithHTTPClient(t *testing.T) {
	customClient := &http.Client{}

	tests := []struct {
		name   string
		client *http.Client
	}{
		{
			name:   "sets a custom http client",
			client: customClient,
		},
		{
			name:   "nil http client is accepted",
			client: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := &AuthConfig{}
			WithHTTPClient(test.client)(config)

			assert.Equal(t, test.client, config.httpClient)
		})
	}
}

func TestWithScopes(t *testing.T) {
	tests := []struct {
		name   string
		scopes []string
	}{
		{
			name:   "sets a single scope",
			scopes: []string{"read"},
		},
		{
			name:   "sets multiple scopes",
			scopes: []string{"read", "write"},
		},
		{
			name:   "no scopes clears to empty slice",
			scopes: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := &AuthConfig{}
			WithScopes(test.scopes...)(config)

			assert.Equal(t, test.scopes, config.scopes)
		})
	}
}

// TODO: should be a Mock token store
type optionsFakeTokenStore struct{}

func (f *optionsFakeTokenStore) Save(_ context.Context, _ string, _ *AccessToken) error {
	return nil
}

func (f *optionsFakeTokenStore) Load(_ context.Context, _ string) (*AccessToken, error) {
	return nil, nil
}

func (f *optionsFakeTokenStore) Delete(_ context.Context, _ string) error {
	return nil
}
