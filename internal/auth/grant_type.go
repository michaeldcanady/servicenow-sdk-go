package auth

import (
	"fmt"
	"strings"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

var _ serialization.EnumFactory

type GrantType int64

const (
	GrantTypeUnknown      GrantType = -1
	GrantTypeRefreshToken GrantType = iota
	GrantTypePassword
	GrantTypeAuthorizationCode
)

func (gT GrantType) String() string {
	return map[GrantType]string{
		GrantTypeRefreshToken:      "refresh_token",
		GrantTypePassword:          "password",
		GrantTypeAuthorizationCode: "authorization_code",
	}[gT]
}

func ParseGrantTypeType(v string) (interface{}, error) {
	switch strings.ToLower(v) {
	case GrantTypeRefreshToken.String():
		return pointer(GrantTypeRefreshToken), nil
	case GrantTypePassword.String():
		return pointer(GrantTypePassword), nil
	case GrantTypeAuthorizationCode.String():
		return pointer(GrantTypeAuthorizationCode), nil
	default:
		return pointer(GrantTypeUnknown), fmt.Errorf("Unknown GranType value: %s", v)
	}
}
