package core

import (
	"context"
	"errors"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

var _ authentication.AuthenticationProvider = (*APIV1ClientAdapter)(nil)

// APIV1ClientAdapter an adapter to adapt from a V1 client to a kiota RequestAdapter
type APIV1ClientAdapter struct {
	client Client
}

// NewAPIV1ClientAdapter instantiates an authentication provider using a V1 client.
func NewAPIV1ClientAdapter(client Client) *APIV1ClientAdapter {
	return &APIV1ClientAdapter{
		client: client,
	}
}

// AuthenticateRequest authenticates the provided RequestInformation.
func (c *APIV1ClientAdapter) AuthenticateRequest(context context.Context, request *abstractions.RequestInformation, additionalAuthenticationContext map[string]interface{}) error {
	temp := NewRequestInformation()

	_, _ = c.client.Send(temp, nil)

	auth := temp.Headers.Clone().Get("Authorization")
	if auth == "" {
		return errors.New("failed to retrieve authorization")
	}

	request.Headers.Add("Authorization", auth)

	return nil
}
