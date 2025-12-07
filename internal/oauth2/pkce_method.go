package oauth2

type PKCEMethod int

const (
	PKCEMethodUnknown PKCEMethod = iota - 1
	PKCEMethodPlain
	PKCEMethodS256
)

func (e PKCEMethod) String() string {
	str, ok := map[PKCEMethod]string{
		PKCEMethodPlain:   "plain",
		PKCEMethodS256:    "S256",
		PKCEMethodUnknown: "unknown",
	}[e]
	if !ok {
		return PKCEMethodUnknown.String()
	}
	return str
}
