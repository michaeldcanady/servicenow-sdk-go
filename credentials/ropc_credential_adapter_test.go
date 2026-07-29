package credentials

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestROPCCredential_GetAuthentication(t *testing.T) {
	tests := []struct {
		name     string
		client   *ropcTestClient
		wantAuth string
		wantErr  string
	}{
		{
			name:    "token acquisition error propagates",
			client:  &ropcTestClient{acquireErr: errors.New("acquire failed")},
			wantErr: "initial token acquisition failed",
		},
		{
			name:     "successful acquisition returns bearer header",
			client:   &ropcTestClient{acquireToken: &AccessToken{AccessToken: "ropc-token", ExpiresAt: time.Now().Add(time.Hour)}},
			wantAuth: "Bearer ropc-token",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			credential, err := NewROPCCredential(test.client, "user", "pass", nil)
			require.NoError(t, err)

			auth, err := credential.GetAuthentication()

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				assert.Empty(t, auth)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.wantAuth, auth)
		})
	}
}
