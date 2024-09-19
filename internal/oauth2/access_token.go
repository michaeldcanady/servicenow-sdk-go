package oauth2

import "time"

type Oauth2Token interface {
	IsExpired() bool
	// AccessToken is the actual access token issued by the authorization server.
	GetAccessToken() string
	// ExpiresIn is the duration in seconds for which the access token is valid.
	GetExpiresIn() int
	// ExpiresAt is the time at which the access token expires.
	GetExpiresAt() time.Time
	// TokenType is the type of token, usually "Bearer" for OAuth2 access tokens.
	GetTokenType() string
}

type HasRefreshToken interface {
	GetRefreshToken() string
}

type SupportsScopes interface {
	// Scope is a space-delimited list of permissions or access rights granted by the token.
	GetScopes() []string
}
