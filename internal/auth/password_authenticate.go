package auth

import (
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	usernameKey = "username"
	passwordKey = "password"
)

type PasswordAuthenticatable interface {
	Authenticatable
	GetUsername() (*string, error)
	SetUsername(username *string) error
	GetPassword() (*string, error)
	SetPassword(password *string) error
}

type passwordAuthenticate struct {
	Authenticatable
}

func NewPasswordAuthenticate() PasswordAuthenticatable {
	return &passwordAuthenticate{
		Authenticatable: NewAuthenticate(),
	}
}

// GetFieldDeserializers returns the deserialization information for this object.
func (a *passwordAuthenticate) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if core.IsNil(a) {
		return nil
	}

	fields := map[string]func(serialization.ParseNode) error{
		usernameKey: func(pn serialization.ParseNode) error {
			username, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			return a.SetUsername(username)
		},
		passwordKey: func(pn serialization.ParseNode) error {
			password, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			return a.SetPassword(password)
		},
	}

	maps.Copy(fields, a.Authenticatable.GetFieldDeserializers())

	return fields
}

func (a *passwordAuthenticate) GetUsername() (*string, error) {
	username, err := a.GetBackingStore().Get(usernameKey)
	if err != nil {
		return nil, err
	}
	typedUsername, ok := username.(*string)
	if !ok {
		return nil, errors.New("username is not *string")
	}
	return typedUsername, nil
}

func (a *passwordAuthenticate) SetUsername(username *string) error {
	if core.IsNil(a) {
		return nil
	}

	return a.GetBackingStore().Set(usernameKey, username)
}

func (a *passwordAuthenticate) GetPassword() (*string, error) {
	if core.IsNil(a) {
		return nil, nil
	}

	password, err := a.GetBackingStore().Get(passwordKey)
	if err != nil {
		return nil, err
	}
	typedPassword, ok := password.(*string)
	if !ok {
		return nil, errors.New("password is not *string")
	}
	return typedPassword, nil
}

func (a *passwordAuthenticate) SetPassword(password *string) error {
	if core.IsNil(a) {
		return nil
	}

	return a.GetBackingStore().Set(passwordKey, password)
}
