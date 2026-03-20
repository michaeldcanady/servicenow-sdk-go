package credentials

import "time"

// AccessToken represents an OAuth2 access token and associated metadata.
type AccessToken struct {
	// AccessToken is the actual access token issued by the authorization server.
	AccessToken string `json:"access_token"` //nolint:gosec // G117: Access token key, not secret
	// ExpiresIn is the duration in seconds for which the access token is valid.
	ExpiresIn int `json:"expires_in"`
	// ExpiresAt is the time at which the access token expires.
	ExpiresAt time.Time
	// RefreshToken is an optional token used to obtain a new access token without requiring reauthorization.
	RefreshToken string `json:"refresh_token"` //nolint:gosec // G117: Refresh token key, not secret
	// Scope is a space-delimited list of permissions or access rights granted by the token.
	Scope string `json:"scope"`
	// TokenType is the type of token, usually "Bearer" for OAuth2 access tokens.
	TokenType string `json:"token_type"`
}

// IsExpired checks if the access token is expired, including a 20-second buffer for clock skew.
func (t *AccessToken) IsExpired() bool {
	return time.Now().Add(20 * time.Second).After(t.ExpiresAt)
}
