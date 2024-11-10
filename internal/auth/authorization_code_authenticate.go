package auth

import (
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	codeKey        = "code"
	redirectURIKey = "redirect_uri"
	stateKey       = "state"
)

type AuthorizationCodeAuthenticatable interface {
	Authenticatable
	GetCode() (*string, error)
	SetCode(code *string) error
	GetRedirectURI() (*string, error)
	SetRedirectURI(redirectURI *string) error
}

type authorizationCodeAuthenticate struct {
	Authenticatable
}

func NewAuthorizationCodeAuthenticate() AuthorizationCodeAuthenticatable {
	return &authorizationCodeAuthenticate{
		Authenticatable: NewAuthenticate(),
	}
}

// GetFieldDeserializers returns the deserialization information for this object.
func (a *authorizationCodeAuthenticate) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if core.IsNil(a) {
		return nil
	}
	fields := map[string]func(serialization.ParseNode) error{
		codeKey: func(pn serialization.ParseNode) error {
			code, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			return a.SetCode(code)
		},
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

func (a *authorizationCodeAuthenticate) GetCode() (*string, error) {
	if core.IsNil(a) {
		return nil, nil
	}

	code, err := a.GetBackingStore().Get(codeKey)
	if err != nil {
		return nil, err
	}
	typedCode, ok := code.(*string)
	if !ok {
		return nil, errors.New("redirectURI is not *string")
	}
	return typedCode, nil
}

func (a *authorizationCodeAuthenticate) SetCode(code *string) error {
	if core.IsNil(a) {
		return nil
	}

	return a.GetBackingStore().Set(codeKey, code)
}

func (a *authorizationCodeAuthenticate) GetRedirectURI() (*string, error) {
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

func (a *authorizationCodeAuthenticate) SetRedirectURI(redirectURI *string) error {
	if core.IsNil(a) {
		return nil
	}

	return a.GetBackingStore().Set(redirectURIKey, redirectURI)
}
