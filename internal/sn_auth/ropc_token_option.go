package snauth

import "github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"

type ROPCTokenOption = oauth2.TokenOption[ROPCTokenConfig]

func WithClientID(id string) ROPCTokenOption {
	return func(rc ROPCTokenConfig) error {
		rc.clientID = id
		return nil
	}
}

func WithClientSecret(secret string) ROPCTokenOption {
	return func(rc ROPCTokenConfig) error {
		rc.clientID = secret
		return nil
	}
}
