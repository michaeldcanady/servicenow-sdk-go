package authentication

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	refreshTokenKey = "refresh_token"
)

type refreshTokenRequestable interface {
	grantTypeRequestable
	GetClientID() (*string, error)
	SetClientID(*string) error
	GetClientSecret() (*string, error)
	SetClientSecret(*string) error
	GetRefreshToken() (*string, error)
	SetRefreshToken(*string) error
	serialization.Parsable
	store.BackedModel
}

type refreshTokenRequest struct {
	grantTypeRequestable
}

func newRefreshTokenRequest(opts ...grantTypeRequestOption) refreshTokenRequestable {
	req := &refreshTokenRequest{
		grantTypeRequestable: newGrantTypeRequest(opts...),
	}

	return req
}

// Serialize writes the objects properties to the current writer.
func (request *refreshTokenRequest) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(request) {
		return nil
	}

	fieldSerializers := []func(serialization.SerializationWriter) error{
		func(writer serialization.SerializationWriter) error {
			clientID, err := request.GetClientID()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(clientIDKey, clientID)
		},
		func(writer serialization.SerializationWriter) error {
			clientSecret, err := request.GetClientSecret()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(clientSecretKey, clientSecret)
		},
		func(writer serialization.SerializationWriter) error {
			code, err := request.GetRefreshToken()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(refreshTokenKey, code)
		},
	}

	for _, fieldSerializer := range fieldSerializers {
		if err := fieldSerializer(writer); err != nil {
			return err
		}
	}

	if err := request.grantTypeRequestable.Serialize(writer); err != nil {
		return err
	}

	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (request *refreshTokenRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func (request *refreshTokenRequest) GetClientID() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	clientID, err := request.GetBackingStore().Get(clientIDKey)
	if err != nil {
		return nil, err
	}

	typedClientID, ok := clientID.(*string)
	if !ok {
		return nil, errors.New("clientID is not *string")
	}

	return typedClientID, nil
}

func (request *refreshTokenRequest) SetClientID(clientID *string) error {
	if internal.IsNil(request) {
		return nil
	}

	return request.GetBackingStore().Set(clientIDKey, clientID)
}

func (request *refreshTokenRequest) GetClientSecret() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	clientSecret, err := request.GetBackingStore().Get(clientSecretKey)
	if err != nil {
		return nil, err
	}

	typedClientSecret, ok := clientSecret.(*string)
	if !ok {
		return nil, errors.New("clientSecret is not *string")
	}

	return typedClientSecret, nil
}

func (request *refreshTokenRequest) SetClientSecret(clientSecret *string) error {
	if internal.IsNil(request) {
		return nil
	}

	return request.GetBackingStore().Set(clientSecretKey, clientSecret)
}

func (request *refreshTokenRequest) GetRefreshToken() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	code, err := request.GetBackingStore().Get(refreshTokenKey)
	if err != nil {
		return nil, err
	}

	typedRefreshToken, ok := code.(*string)
	if !ok {
		return nil, errors.New("code is not *string")
	}

	return typedRefreshToken, nil
}

func (request *refreshTokenRequest) SetRefreshToken(code *string) error {
	if internal.IsNil(request) {
		return nil
	}

	return request.GetBackingStore().Set(refreshTokenKey, code)
}
