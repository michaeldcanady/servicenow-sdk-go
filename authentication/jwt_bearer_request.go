package authentication

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type JWTBearerRequestable interface {
	grantTypeRequestable
	GetClientID() (*string, error)
	SetClientID(*string) error
	GetClientSecret() (*string, error)
	SetClientSecret(*string) error
	GetAssertion() (*string, error)
	SetAssertion(*string) error
	serialization.Parsable
	store.BackedModel
}

type JWTBearerRequest struct {
	grantTypeRequestable
}

type JWTBearerRequestOption func(*JWTBearerRequest)

func newJWTBearerRequest(opts ...grantTypeRequestOption) JWTBearerRequestable {
	req := &JWTBearerRequest{
		grantTypeRequestable: newGrantTypeRequest(opts...),
	}

	return req
}

// Serialize writes the objects properties to the current writer.
func (request *JWTBearerRequest) Serialize(writer serialization.SerializationWriter) error {
	return errors.New("not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (request *JWTBearerRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func (request *JWTBearerRequest) GetClientID() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	return nil, errors.New("not implemented")
}

func (request *JWTBearerRequest) SetClientID(*string) error {
	if internal.IsNil(request) {
		return nil
	}

	return errors.New("not implemented")
}

func (request *JWTBearerRequest) GetClientSecret() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	return nil, errors.New("not implemented")
}

func (request *JWTBearerRequest) SetClientSecret(*string) error {
	if internal.IsNil(request) {
		return nil
	}

	return errors.New("not implemented")
}

func (request *JWTBearerRequest) GetAssertion() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	return nil, errors.New("not implemented")
}

func (request *JWTBearerRequest) SetAssertion(*string) error {
	if internal.IsNil(request) {
		return nil
	}

	return errors.New("not implemented")
}
