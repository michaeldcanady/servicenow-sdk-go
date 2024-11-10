package auth

import (
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type RefreshTokenAuthenticatable interface {
	Authenticatable
	GetRefreshToken() (*string, error)
	SetRefreshToken(refreshToken *string) error
}

type refreshTokenAuthenticate struct {
	Authenticatable
}

func NewRefreshTokenAuthenticate() RefreshTokenAuthenticatable {
	return &refreshTokenAuthenticate{
		Authenticatable: NewAuthenticate(),
	}
}

// GetFieldDeserializers returns the deserialization information for this object.
func (a *refreshTokenAuthenticate) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if core.IsNil(a) {
		return nil
	}
	fields := map[string]func(serialization.ParseNode) error{
		refreshTokenKey: func(pn serialization.ParseNode) error {
			username, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			return a.SetRefreshToken(username)
		},
	}

	maps.Copy(fields, a.Authenticatable.GetFieldDeserializers())

	return fields
}

func (a *refreshTokenAuthenticate) GetRefreshToken() (*string, error) {
	if core.IsNil(a) {
		return nil, nil
	}

	refreshToken, err := a.GetBackingStore().Get(refreshTokenKey)
	if err != nil {
		return nil, err
	}
	typedRefreshToken, ok := refreshToken.(*string)
	if !ok {
		return nil, errors.New("password is not *string")
	}
	return typedRefreshToken, nil
}
func (a *refreshTokenAuthenticate) SetRefreshToken(refreshToken *string) error {
	if core.IsNil(a) {
		return nil
	}

	return a.GetBackingStore().Set(refreshTokenKey, refreshToken)
}
