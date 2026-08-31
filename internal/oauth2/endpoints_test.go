// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package oauth2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEndpoints_Validate(t *testing.T) {
	tests := []struct {
		name        string
		endpoints   *Endpoints
		grantType   string
		expectedErr string
	}{
		{
			name:        "nil receiver",
			endpoints:   nil,
			grantType:   GrantTypeClientCreds,
			expectedErr: "endpoints are not set",
		},
		{
			name:        "missing token URL",
			endpoints:   &Endpoints{},
			grantType:   GrantTypeClientCreds,
			expectedErr: "token endpoint is not set",
		},
		{
			name:        "whitespace-only token URL is treated as missing",
			endpoints:   &Endpoints{TokenURL: "   "},
			grantType:   GrantTypeClientCreds,
			expectedErr: "token endpoint is not set",
		},
		{
			name:      "token URL alone satisfies client credentials",
			endpoints: &Endpoints{TokenURL: "http://token"},
			grantType: GrantTypeClientCreds,
		},
		{
			name:      "token URL alone satisfies refresh token",
			endpoints: &Endpoints{TokenURL: "http://token"},
			grantType: GrantTypeRefreshToken,
		},
		{
			name:      "token URL alone satisfies password",
			endpoints: &Endpoints{TokenURL: "http://token"},
			grantType: GrantTypePassword,
		},
		{
			name:      "token URL alone satisfies jwt bearer",
			endpoints: &Endpoints{TokenURL: "http://token"},
			grantType: GrantTypeJWTBearer,
		},
		{
			name:      "unknown grant type only needs the token URL",
			endpoints: &Endpoints{TokenURL: "http://token"},
			grantType: "something_else",
		},
		{
			name:        "authorization code requires the auth URL",
			endpoints:   &Endpoints{TokenURL: "http://token"},
			grantType:   GrantTypeAuthCode,
			expectedErr: "authorization endpoint is not set",
		},
		{
			name:        "authorization code rejects a whitespace-only auth URL",
			endpoints:   &Endpoints{TokenURL: "http://token", AuthURL: "  "},
			grantType:   GrantTypeAuthCode,
			expectedErr: "authorization endpoint is not set",
		},
		{
			name:      "authorization code with the auth URL set",
			endpoints: &Endpoints{TokenURL: "http://token", AuthURL: "http://auth"},
			grantType: GrantTypeAuthCode,
		},
		{
			name:        "device code requires the device URL",
			endpoints:   &Endpoints{TokenURL: "http://token"},
			grantType:   GrantTypeDeviceCode,
			expectedErr: "device authorization endpoint is not set",
		},
		{
			name:        "device code rejects a whitespace-only device URL",
			endpoints:   &Endpoints{TokenURL: "http://token", DeviceURL: "\t"},
			grantType:   GrantTypeDeviceCode,
			expectedErr: "device authorization endpoint is not set",
		},
		{
			name:      "device code with the device URL set",
			endpoints: &Endpoints{TokenURL: "http://token", DeviceURL: "http://device"},
			grantType: GrantTypeDeviceCode,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.endpoints.Validate(test.grantType)

			if test.expectedErr != "" {
				require.EqualError(t, err, test.expectedErr)

				return
			}

			require.NoError(t, err)
		})
	}
}
