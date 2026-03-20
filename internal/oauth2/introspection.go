package oauth2

// IntrospectionResponse represents the response from the token introspection endpoint (RFC 7662).
type IntrospectionResponse struct {
	// Active is true if the token is valid and hasn't expired or been revoked.
	Active bool `json:"active"`
	// Scope is the list of scopes associated with the token.
	Scope string `json:"scope,omitempty"`
	// ClientID is the identifier for the client that requested the token.
	ClientID string `json:"client_id,omitempty"`
	// Username of the resource owner who authorized the token.
	Username string `json:"username,omitempty"`
	// TokenType is the type of the token (e.g., Bearer).
	TokenType string `json:"token_type,omitempty"`
	// Exp is the expiration time of the token (Unix timestamp).
	Exp int64 `json:"exp,omitempty"`
	// Iat is the time the token was issued (Unix timestamp).
	Iat int64 `json:"iat,omitempty"`
	// Nbf is the time before which the token must not be accepted (Unix timestamp).
	Nbf int64 `json:"nbf,omitempty"`
	// Sub is the subject of the token (usually the resource owner's ID).
	Sub string `json:"sub,omitempty"`
	// Aud is the intended audience for the token.
	Aud string `json:"aud,omitempty"`
	// Iss is the issuer of the token.
	Iss string `json:"iss,omitempty"`
	// Jti is the unique identifier for the token.
	Jti string `json:"jti,omitempty"`
	// Raw holds all returned metadata for the token.
	Raw map[string]any `json:"-"`
}
