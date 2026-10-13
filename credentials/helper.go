// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
)

func convertToken(t *oauth2.Token) *AccessToken {
	if t == nil {
		return nil
	}
	return &AccessToken{
		AccessToken:  t.AccessToken,
		ExpiresIn:    int(t.ExpiresIn),
		ExpiresAt:    time.Now().Add(time.Duration(t.ExpiresIn) * time.Second),
		RefreshToken: t.RefreshToken,
		Scope:        t.Scope,
		TokenType:    t.TokenType,
	}
}
