package credentials

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const authorizationHeaderKey = "Authorization"

type RequestAdapterAdapter struct {
	credential Credential
}

func NewRequestAdapterAdapter(credential Credential) *RequestAdapterAdapter {
	return &RequestAdapterAdapter{
		credential: credential,
	}
}

// AuthenticateRequest authenticates the provided RequestInformation.
func (c *RequestAdapterAdapter) AuthenticateRequest(context context.Context, request *abstractions.RequestInformation, additionalAuthenticationContext map[string]interface{}) error {
	if internal.IsNil(c) || internal.IsNil(c.credential) {
		return errors.New("credential is nil")
	}

	auth, err := c.credential.GetAuthentication()
	if err != nil {
		return err
	}

	request.Headers.Add(authorizationHeaderKey, auth)
	return nil
}
