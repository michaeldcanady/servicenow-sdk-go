package authentication

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type resourceOwnerRequestable interface {
	grantTypeRequestable
	GetClientID() (*string, error)
	SetClientID(*string) error
	GetClientSecret() (*string, error)
	SetClientSecret(*string) error
	GetUsername() (*string, error)
	SetUsername(*string) error
	GetPassword() (*string, error)
	SetPassword(*string) error
}

type resourceOwnerRequest struct {
	grantTypeRequestable
}

func newResourceOwnerRequest() resourceOwnerRequestable {
	return &resourceOwnerRequest{
		grantTypeRequestable: newGrantTypeRequest(),
	}
}

func (request *resourceOwnerRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}
func (request *resourceOwnerRequest) Serialize(writer serialization.SerializationWriter) error {
	return errors.New("not implemented")
}

func (request *resourceOwnerRequest) GetClientID() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	clientID, err := request.GetBackingStore().Get("client_id")
	if err != nil {
		return nil, err
	}

	typedClientID, ok := clientID.(*string)
	if !ok {
		return nil, errors.New("clientID is not *string")
	}

	return typedClientID, nil
}

func (request *resourceOwnerRequest) SetClientID(clientID *string) error {
	if internal.IsNil(request) {
		return nil
	}

	return request.GetBackingStore().Set("client_id", clientID)
}

func (request *resourceOwnerRequest) GetClientSecret() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	clientSecret, err := request.GetBackingStore().Get("client_secret")
	if err != nil {
		return nil, err
	}

	typedClientSecret, ok := clientSecret.(*string)
	if !ok {
		return nil, errors.New("clientSecret is not *string")
	}

	return typedClientSecret, nil
}

func (request *resourceOwnerRequest) SetClientSecret(clientSecret *string) error {
	if internal.IsNil(request) {
		return nil
	}

	return request.GetBackingStore().Set("client_secret", clientSecret)
}

func (request *resourceOwnerRequest) GetUsername() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	username, err := request.GetBackingStore().Get("username")
	if err != nil {
		return nil, err
	}

	typedUsername, ok := username.(*string)
	if !ok {
		return nil, errors.New("username is not *string")
	}

	return typedUsername, nil
}

func (request *resourceOwnerRequest) SetUsername(username *string) error {
	if internal.IsNil(request) {
		return nil
	}

	return request.GetBackingStore().Set("username", username)
}

func (request *resourceOwnerRequest) GetPassword() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	password, err := request.GetBackingStore().Get("password")
	if err != nil {
		return nil, err
	}

	typedPassword, ok := password.(*string)
	if !ok {
		return nil, errors.New("password is not *string")
	}

	return typedPassword, nil
}

func (request *resourceOwnerRequest) SetPassword(password *string) error {
	if internal.IsNil(request) {
		return nil
	}

	return request.GetBackingStore().Set("password", password)
}
