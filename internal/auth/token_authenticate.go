package auth

import (
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type TokenAuthenticatable interface {
	Authenticatable
	GetRedirectURI() (*string, error)
	SetRedirectURI(redirectURI *string) error
}

type tokenAuthenticate struct {
	Authenticatable
}

func NewTokenAuthenticate() TokenAuthenticatable {
	return &tokenAuthenticate{
		Authenticatable: NewAuthenticate(),
	}
}

// GetFieldDeserializers returns the deserialization information for this object.
func (a *tokenAuthenticate) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if core.IsNil(a) {
		return nil
	}
	fields := map[string]func(serialization.ParseNode) error{
		redirectURIKey: func(pn serialization.ParseNode) error {
			redirectURI, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			return a.SetRedirectURI(redirectURI)
		},
	}

	maps.Copy(fields, a.Authenticatable.GetFieldDeserializers())

	return fields
}

func (a *tokenAuthenticate) GetRedirectURI() (*string, error) {
	if core.IsNil(a) {
		return nil, nil
	}

	redirectURI, err := a.GetBackingStore().Get(refreshTokenKey)
	if err != nil {
		return nil, err
	}
	typedRedirectURI, ok := redirectURI.(*string)
	if !ok {
		return nil, errors.New("redirectURI is not *string")
	}
	return typedRedirectURI, nil
}

func (a *tokenAuthenticate) SetRedirectURI(redirectURI *string) error {
	if core.IsNil(a) {
		return nil
	}

	return a.GetBackingStore().Set(redirectURIKey, redirectURI)
}
