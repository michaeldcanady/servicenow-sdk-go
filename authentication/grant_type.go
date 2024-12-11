package authentication

type grantType int64

const (
	grantTypeUnknown grantType = iota - 1
	grantTypeAuthorizationCode
	grantTypeRefreshToken
	grantTypeJWTBearer
	grantTypePassword
)

func (g grantType) String() string {
	grantTypeString, found := map[grantType]string{
		grantTypeUnknown:           "unknown",
		grantTypeAuthorizationCode: "authorization_code",
		grantTypeRefreshToken:      "refresh_token",
		grantTypeJWTBearer:         "urn:ietf:params:oauth:grant-type:jwt-bearer",
		grantTypePassword:          "password",
	}[g]
	if !found {
		return grantTypeUnknown.String()
	}
	return grantTypeString
}
