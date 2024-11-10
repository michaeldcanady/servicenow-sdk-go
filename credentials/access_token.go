package credentials

import "time"

// AccessToken represents an OAuth2 access token and associated metadata.
type AccessToken struct {
	// AccessToken is the actual access token issued by the authorization server.
	AccessToken string `json:"access_token"`
	// ExpiresIn is the duration in seconds for which the access token is valid.
	ExpiresIn int `json:"expires_in"`
	// ExpiresAt is the time at which the access token expires.
	ExpiresAt time.Time
	// RefreshToken is an optional token used to obtain a new access token without requiring reauthorization.
	RefreshToken string `json:"refresh_token"`
	// Scope is a space-delimited list of permissions or access rights granted by the token.
	Scope string `json:"scope"`
	// TokenType is the type of token, usually "Bearer" for OAuth2 access tokens.
	TokenType string `json:"token_type"`
}

// IsExpired Checks if the access token is expired
func (t *AccessToken) IsExpired() bool {
	return t.ExpiresAt.Before(time.Now())
}
