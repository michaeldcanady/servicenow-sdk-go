// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package support

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// credentialEnvKeys are pinned in every case so ambient SN_*/SNOW_* variables
// (e.g. real secrets set in the e2e-nightly job) cannot leak into assertions.
var credentialEnvKeys = []string{
	"SN_CLIENT_ID", "SN_CLIENT_SECRET",
	"SNOW_CLIENT_ID", "SNOW_CLIENT_SECRET",
}

func TestClientCredentials(t *testing.T) {
	tests := []struct {
		name       string
		env        map[string]string
		wantID     string
		wantSecret string
	}{
		{
			name:       "defaults when nothing set",
			env:        map[string]string{},
			wantID:     "mock-client-id",
			wantSecret: "mock-client-secret",
		},
		{
			name: "SN_ prefixed values take precedence",
			env: map[string]string{
				"SN_CLIENT_ID":       "live-id",
				"SN_CLIENT_SECRET":   "live-secret",
				"SNOW_CLIENT_ID":     "snow-id",
				"SNOW_CLIENT_SECRET": "snow-secret",
			},
			wantID:     "live-id",
			wantSecret: "live-secret",
		},
		{
			name: "SNOW_ aliases used when SN_ unset",
			env: map[string]string{
				"SNOW_CLIENT_ID":     "snow-id",
				"SNOW_CLIENT_SECRET": "snow-secret",
			},
			wantID:     "snow-id",
			wantSecret: "snow-secret",
		},
		{
			name: "empty strings fall through to next candidate",
			env: map[string]string{
				"SN_CLIENT_ID":       "",
				"SN_CLIENT_SECRET":   "",
				"SNOW_CLIENT_ID":     "snow-id",
				"SNOW_CLIENT_SECRET": "snow-secret",
			},
			wantID:     "snow-id",
			wantSecret: "snow-secret",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, k := range credentialEnvKeys {
				t.Setenv(k, tt.env[k]) // empty string == unset for FirstEnv
			}
			gotID, gotSecret := ClientCredentials()
			assert.Equal(t, tt.wantID, gotID)
			assert.Equal(t, tt.wantSecret, gotSecret)
		})
	}
}

// modeEnvKeys are cleared so ambient SN_*/SNOW_* values cannot leak into the
// mode assertions below.
var modeEnvKeys = []string{
	"SN_OFFLINE",
	"SN_USERNAME", "SN_PASSWORD",
	"SNOW_USERNAME", "SNOW_PASSWORD",
}

func TestRunMode(t *testing.T) {
	tests := []struct {
		name        string
		env         map[string]string
		wantE2E     bool
		wantOffline bool
	}{
		{
			name:        "nothing set defaults to mocked",
			env:         map[string]string{},
			wantE2E:     false,
			wantOffline: true,
		},
		{
			name:        "SN_OFFLINE unset with credentials stays mocked",
			env:         map[string]string{"SN_USERNAME": "live_user", "SN_PASSWORD": "live_pass"},
			wantE2E:     false,
			wantOffline: true,
		},
		{
			name:        "explicit SN_OFFLINE=true is mocked",
			env:         map[string]string{"SN_OFFLINE": "true", "SN_USERNAME": "live_user"},
			wantE2E:     false,
			wantOffline: true,
		},
		{
			name:        "explicit SN_OFFLINE=false is live",
			env:         map[string]string{"SN_OFFLINE": "false", "SN_USERNAME": "live_user"},
			wantE2E:     true,
			wantOffline: false,
		},
		{
			name:        "unrecognized SN_OFFLINE value stays mocked",
			env:         map[string]string{"SN_OFFLINE": "yes", "SN_USERNAME": "live_user"},
			wantE2E:     false,
			wantOffline: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, k := range modeEnvKeys {
				t.Setenv(k, tt.env[k]) // empty string == unset
			}
			assert.Equal(t, tt.wantE2E, IsE2E())
			assert.Equal(t, tt.wantOffline, IsOffline())
			// IsOffline must always be the exact negation of IsE2E.
			assert.Equal(t, !IsE2E(), IsOffline())
		})
	}
}

func TestRequireCredentials(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		wantErr bool
	}{
		{
			name:    "mocked mode requires nothing",
			env:     map[string]string{},
			wantErr: false,
		},
		{
			name:    "mocked mode with partial credentials still requires nothing",
			env:     map[string]string{"SN_USERNAME": "live_user"},
			wantErr: false,
		},
		{
			name:    "live mode without credentials errors",
			env:     map[string]string{"SN_OFFLINE": "false"},
			wantErr: true,
		},
		{
			name:    "live mode with username but no password errors",
			env:     map[string]string{"SN_OFFLINE": "false", "SN_USERNAME": "live_user"},
			wantErr: true,
		},
		{
			name:    "live mode with full credentials passes",
			env:     map[string]string{"SN_OFFLINE": "false", "SN_USERNAME": "live_user", "SN_PASSWORD": "live_pass"},
			wantErr: false,
		},
		{
			name:    "live mode accepts SNOW_ aliases",
			env:     map[string]string{"SN_OFFLINE": "false", "SNOW_USERNAME": "live_user", "SNOW_PASSWORD": "live_pass"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, k := range modeEnvKeys {
				t.Setenv(k, tt.env[k])
			}
			err := RequireCredentials()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
