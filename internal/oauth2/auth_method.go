package oauth2

type AuthMethod int

const (
	AuthMethodUnknown AuthMethod = iota - 1
	AuthMethodClientSecretPost
	AuthMethodClientSecretBasic
)

func (e AuthMethod) String() string {
	str, ok := map[AuthMethod]string{
		AuthMethodClientSecretPost:  "client_secret_post",
		AuthMethodClientSecretBasic: "client_secret_basic",
		AuthMethodUnknown:           "unknown",
	}[e]
	if !ok {
		return AuthMethodUnknown.String()
	}
	return str
}
