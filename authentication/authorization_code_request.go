package authentication

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	redirectURIKey       = "redirect_url"
	authorizationCodeKey = "authorization_code"
	codeKey              = "code"
	clientIDKey          = "client_id"
	clientSecretKey      = "client_secret"
)

type authorizationCodeRequestable interface {
	GetGrantType() (*string, error)
	setGrantType(*string) error
	GetClientID() (*string, error)
	SetClientID(*string) error
	GetClientSecret() (*string, error)
	SetClientSecret(*string) error
	GetCode() (*string, error)
	SetCode(*string) error
	GetRedirectURI() (*string, error)
	SetRedirectURI(*string) error
	serialization.Parsable
	store.BackedModel
}

type authorizationCodeRequest struct {
	grantTypeRequestable
}

type authorizationCodeRequestOption func(*authorizationCodeRequest)

func newAuthorizationCodeRequest(opts ...grantTypeRequestOption) authorizationCodeRequestable {
	req := &authorizationCodeRequest{
		grantTypeRequestable: newGrantTypeRequest(opts...),
	}

	return req
}

// Serialize writes the objects properties to the current writer.
func (request *authorizationCodeRequest) Serialize(writer serialization.SerializationWriter) error {
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
			code, err := request.GetCode()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(codeKey, code)
		},
		func(writer serialization.SerializationWriter) error {
			redirectURI, err := request.GetRedirectURI()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(redirectURIKey, redirectURI)
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
func (request *authorizationCodeRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func (request *authorizationCodeRequest) GetClientID() (*string, error) {
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

func (request *authorizationCodeRequest) SetClientID(clientID *string) error {
	if internal.IsNil(request) {
		return nil
	}

	return request.GetBackingStore().Set(clientIDKey, clientID)
}

func (request *authorizationCodeRequest) GetClientSecret() (*string, error) {
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

func (request *authorizationCodeRequest) SetClientSecret(clientSecret *string) error {
	if internal.IsNil(request) {
		return nil
	}

	return request.GetBackingStore().Set(clientSecretKey, clientSecret)
}

func (request *authorizationCodeRequest) GetCode() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	code, err := request.GetBackingStore().Get(codeKey)
	if err != nil {
		return nil, err
	}

	typedCode, ok := code.(*string)
	if !ok {
		return nil, errors.New("code is not *string")
	}

	return typedCode, nil
}

func (request *authorizationCodeRequest) SetCode(code *string) error {
	if internal.IsNil(request) {
		return nil
	}

	return request.GetBackingStore().Set(codeKey, code)
}

func (request *authorizationCodeRequest) GetRedirectURI() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	redirectURI, err := request.GetBackingStore().Get(redirectURIKey)
	if err != nil {
		return nil, err
	}

	typedRedirectURI, ok := redirectURI.(*string)
	if !ok {
		return nil, errors.New("redirectURI is not *string")
	}

	return typedRedirectURI, nil
}

func (request *authorizationCodeRequest) SetRedirectURI(redirectURI *string) error {
	if internal.IsNil(request) {
		return nil
	}

	return request.GetBackingStore().Set(redirectURIKey, redirectURI)
}
