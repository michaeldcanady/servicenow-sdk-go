package oauth2

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// defaultExpiryDelta determines how earlier a token should be considered
// expired than its actual expiration time. It is used to avoid late
// expirations due to client-server time mismatches.
const defaultExpiryDelta = 10 * time.Second

type Token struct {
	AccessToken  string         `json:"access_token"`
	TokenType    string         `json:"token_type"`
	ExpiresIn    int            `json:"expires_in"`
	ExpiresAt    time.Time      `json:"-"`
	RefreshToken string         `json:"refresh_token,omitempty"`
	Scope        string         `json:"scope,omitempty"`
	Raw          map[string]any `json:"-"`
	Headers      http.Header    `json:"-"`
}

func (t *Token) IsExpired() bool {
	now := time.Now().Round(0).UTC()
	fmt.Printf("ExpiresAt: %v, Now: %v\n", t.ExpiresAt, now)
	return t.ExpiresAt.Round(0).Add(-defaultExpiryDelta).Before(now)
}

func (t *Token) IsValid() bool {
	return strings.TrimSpace(t.AccessToken) == "" || t.IsExpired()
}
