package authentication

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type externalJWTRequestable interface {
	GetGrantType() (*string, error)
	setGrantType(*string) error
	GetClientID() (*string, error)
	SetClientID(*string) error
	GetClientSecret() (*string, error)
	SetClientSecret(*string) error
	GetAssertion() (*string, error)
	SetAssertion(*string) error
	serialization.Parsable
	store.BackedModel
}

type externalJWTRequest struct {
	grantTypeRequestable
}

type externalJWTRequestOption func(*externalJWTRequest)

func newExternalJWTRequest(opts ...grantTypeRequestOption) externalJWTRequestable {
	req := &externalJWTRequest{
		grantTypeRequestable: newGrantTypeRequest(opts...),
	}

	return req
}

// Serialize writes the objects properties to the current writer.
func (request *externalJWTRequest) Serialize(writer serialization.SerializationWriter) error {
	return errors.New("not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (request *externalJWTRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func (request *externalJWTRequest) GetClientID() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	return nil, errors.New("not implemented")
}
func (request *externalJWTRequest) SetClientID(*string) error {
	if internal.IsNil(request) {
		return nil
	}

	return errors.New("not implemented")
}
func (request *externalJWTRequest) GetClientSecret() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	return nil, errors.New("not implemented")
}
func (request *externalJWTRequest) SetClientSecret(*string) error {
	if internal.IsNil(request) {
		return nil
	}

	return errors.New("not implemented")
}
func (request *externalJWTRequest) GetAssertion() (*string, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	return nil, errors.New("not implemented")
}
func (request *externalJWTRequest) SetAssertion(*string) error {
	if internal.IsNil(request) {
		return nil
	}

	return errors.New("not implemented")
}
