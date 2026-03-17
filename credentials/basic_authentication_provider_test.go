package credentials

import (
	"context"
	"encoding/base64"
	"fmt"
	"testing"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestBasicAuthenticationProvider_AuthenticateRequest(t *testing.T) {
	tests := []struct {
		name           string
		username       string
		password       string
		ctx            context.Context
		wantErr        bool
		expectedHeader string
	}{
		{
			name:           "Valid Credentials",
			username:       "admin",
			password:       "password",
			ctx:            context.Background(),
			wantErr:        false,
			expectedHeader: fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("admin:password"))),
		},
		{
			name:     "Empty Username",
			username: "",
			password: "password",
			ctx:      context.Background(),
			wantErr:  true,
		},
		{
			name:     "Empty Password",
			username: "admin",
			password: "",
			ctx:      context.Background(),
			wantErr:  true,
		},
		{
			name:     "Cancelled Context",
			username: "admin",
			password: "password",
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			}(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider := NewBasicProvider(tt.username, tt.password)
			request := abstractions.NewRequestInformation()

			err := provider.AuthenticateRequest(tt.ctx, request, nil)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				authHeader := request.Headers.Get("Authorization")
				assert.Equal(t, tt.expectedHeader, authHeader[0])
			}
		})
	}
}

func TestBasicAuthenticationProvider_AuthenticateRequest_NilProvider(t *testing.T) {
	var provider *BasicAuthenticationProvider
	request := abstractions.NewRequestInformation()
	err := provider.AuthenticateRequest(context.Background(), request, nil)
	assert.Error(t, err)
	assert.Equal(t, "provider is nil", err.Error())
}
