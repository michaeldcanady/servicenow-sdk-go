package oauth2

import "net/http"

type Token struct {
	AccessToken  string         `json:"access_token"` //nolint:gosec // G117: Not secret
	TokenType    string         `json:"token_type"`
	ExpiresIn    int64          `json:"expires_in"`
	RefreshToken string         `json:"refresh_token,omitempty"` //nolint:gosec // G117: Not secret
	Scope        string         `json:"scope,omitempty"`
	Raw          map[string]any `json:"-"`
	Headers      http.Header    `json:"-"`
}
