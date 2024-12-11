package authentication

type responseType int64

const (
	responseTypeUnknown responseType = iota - 1
	responseTypeToken
	responseTypeCode
)

func (r responseType) String() string {
	grantTypeString, found := map[responseType]string{
		responseTypeUnknown: "unknown",
		responseTypeToken:   "token",
		responseTypeCode:    "code",
	}[r]
	if !found {
		return responseTypeUnknown.String()
	}
	return grantTypeString
}
