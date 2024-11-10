package auth

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	grantTypeKey    = "grant_type"
	clientIDKey     = "client_id"
	clientSecretKey = "client_secret"
)

type Authenticatable interface {
	GetGrantType() (*GrantType, error)
	SetGrantType(grantType *GrantType) error
	GetClientID() (*string, error)
	SetClientID(clientID *string) error
	GetClientSecret() (*string, error)
	SetClientSecret(clientSecret *string) error
	serialization.Parsable
	store.BackedModel
}

type authenticate struct {
	backingStore store.BackingStore
}

func NewAuthenticate() Authenticatable {
	return &authenticate{
		backingStore: store.BackingStoreFactoryInstance(),
	}
}

// Serialize writes the objects properties to the current writer.
func (a *authenticate) Serialize(writer serialization.SerializationWriter) error {
	return errors.New("Serialize is not implemented")
}

func (a *authenticate) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		grantTypeKey: func(pn serialization.ParseNode) error {
			grantType, err := pn.GetEnumValue(ParseGrantTypeType)
			if err != nil {
				return err
			}
			typedGrantType, ok := grantType.(GrantType)
			if !ok {
				return errors.New("grantType is not GrantType")
			}
			return a.SetGrantType(&typedGrantType)
		},
		clientIDKey: func(pn serialization.ParseNode) error {
			clientID, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			return a.SetClientID(clientID)
		},
		clientSecretKey: func(pn serialization.ParseNode) error {
			clientSecret, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			return a.SetClientSecret(clientSecret)
		},
	}
}

// GetBackingStore returns the store that is backing the model.
func (a *authenticate) GetBackingStore() store.BackingStore {
	if core.IsNil(a) {
		return nil
	}

	if core.IsNil(a.backingStore) {
		a.backingStore = store.NewInMemoryBackingStore()
	}

	return a.backingStore
}

func (a *authenticate) GetGrantType() (*GrantType, error) {
	if core.IsNil(a) {
		return nil, nil
	}

	grantType, err := a.GetBackingStore().Get(grantTypeKey)
	if err != nil {
		return nil, err
	}
	typedGrantType, ok := grantType.(*GrantType)
	if !ok {
		return nil, errors.New("grantType is not *GrantType")
	}
	return typedGrantType, nil
}

func (a *authenticate) SetGrantType(grantType *GrantType) error {
	if core.IsNil(a) {
		return nil
	}
	return a.GetBackingStore().Set(grantTypeKey, grantType)
}

func (a *authenticate) GetClientID() (*string, error) {
	if core.IsNil(a) {
		return nil, nil
	}

	clientID, err := a.GetBackingStore().Get(clientIDKey)
	if err != nil {
		return nil, err
	}
	typedClientID, ok := clientID.(*string)
	if !ok {
		return nil, errors.New("clientID is not *string")
	}
	return typedClientID, nil
}

func (a *authenticate) SetClientID(clientID *string) error {
	if core.IsNil(a) {
		return nil
	}
	return a.GetBackingStore().Set(clientIDKey, clientID)
}

func (a *authenticate) GetClientSecret() (*string, error) {
	if core.IsNil(a) {
		return nil, nil
	}

	clientSecret, err := a.GetBackingStore().Get(clientSecretKey)
	if err != nil {
		return nil, err
	}
	typedclientSecret, ok := clientSecret.(*string)
	if !ok {
		return nil, errors.New("clientSecret is not *string")
	}
	return typedclientSecret, nil
}

func (a *authenticate) SetClientSecret(clientSecret *string) error {
	return a.GetBackingStore().Set(clientSecretKey, clientSecret)
}
