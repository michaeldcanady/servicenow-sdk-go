package oauth2

type SupportsRefresh interface {
	SetRefreshToken(Oauth2Token) error
}
