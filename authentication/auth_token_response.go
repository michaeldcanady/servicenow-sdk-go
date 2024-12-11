package authentication

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type authenticationTokenResponsable interface {
	GetAccessToken() (*string, error)
	setAccessToken(*string) error
	GetScope() (*string, error)
	setScope(*string) error
	GetTokenType() (*string, error)
	setTokenType(*string) error
	GetExpiresIn() (*serialization.ISODuration, error)
	setExpiresIn(*serialization.ISODuration) error
	serialization.Parsable
	store.BackedModel
}

type authenticationTokenResponse struct {
	backingStore store.BackingStore
}

func newAuthenticationTokenResponse() authenticationTokenResponsable {
	return &authenticationTokenResponse{
		backingStore: store.NewInMemoryBackingStore(),
	}
}

// TODO: make a refreshable variation

// CreateAuthenticationTokenResponseFromDiscriminatorValue is a parsable factory for creating a Fileable
func CreateAuthenticationTokenResponseFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return newAuthenticationTokenResponse(), nil
}

// GetBackingStore retrieves the backing store for the model.
func (rE *authenticationTokenResponse) GetBackingStore() store.BackingStore {
	if internal.IsNil(rE) {
		return nil
	}

	if internal.IsNil(rE.backingStore) {
		rE.backingStore = store.NewInMemoryBackingStore()
	}

	return rE.backingStore
}

// Serialize writes the objects properties to the current writer.
func (rE *authenticationTokenResponse) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(rE) {
		return nil
	}

	return errors.New("not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (rE *authenticationTokenResponse) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func (rE *authenticationTokenResponse) GetAccessToken() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	return nil, errors.New("not implemented")
}

func (rE *authenticationTokenResponse) setAccessToken(*string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return errors.New("not implemented")
}

func (rE *authenticationTokenResponse) GetScope() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	return nil, errors.New("not implemented")
}

func (rE *authenticationTokenResponse) setScope(*string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return errors.New("not implemented")
}

func (rE *authenticationTokenResponse) GetTokenType() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	return nil, errors.New("not implemented")
}

func (rE *authenticationTokenResponse) setTokenType(*string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return errors.New("not implemented")
}

func (rE *authenticationTokenResponse) GetExpiresIn() (*serialization.ISODuration, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	return nil, errors.New("not implemented")
}

func (rE *authenticationTokenResponse) setExpiresIn(*serialization.ISODuration) error {
	if internal.IsNil(rE) {
		return nil
	}

	return errors.New("not implemented")
}
