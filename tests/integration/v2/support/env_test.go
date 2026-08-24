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
