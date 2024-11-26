package authentication

type grantType int64

const (
	grantTypeUnknown grantType = iota - 1
	grantTypeAuthorizationCode
)

func (g grantType) String() string {
	grantTypeString, found := map[grantType]string{
		grantTypeUnknown:           "unknown",
		grantTypeAuthorizationCode: "authorization_code",
	}[g]
	if !found {
		return grantTypeUnknown.String()
	}
	return grantTypeString
}
