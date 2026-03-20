package oauth2

// Token represents an OAuth2 access token and its associated metadata.
type Token struct {
	// AccessToken is the actual token used to access protected resources.
	AccessToken string `json:"access_token"` //nolint:gosec // G117: Not secret
	// TokenType is the type of the token (e.g., Bearer).
	TokenType string `json:"token_type"`
	// ExpiresIn is the lifetime of the access token in seconds.
	ExpiresIn int64 `json:"expires_in"`
	// RefreshToken is an optional token used to obtain new access tokens.
	RefreshToken string `json:"refresh_token,omitempty"` //nolint:gosec // G117: Not secret
	// Scope is the list of scopes granted by the token.
	Scope string `json:"scope,omitempty"`
	// Raw holds all returned metadata for the token.
	Raw map[string]any `json:"-"`
}
