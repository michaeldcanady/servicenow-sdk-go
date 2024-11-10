package servicenowsdkgo

import (
	"context"
	"net/url"

	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type ServiceNowServiceClient struct {
	ServiceNowBaseServiceClient
}

// TODO: will be new base for auth provider (credentials)
type serviceNowAuthorizationProvider interface {
	GetAuthorization(context.Context, *url.URL, map[string]interface{}) (string, error)
}

func NewServiceNowServiceClient2(authenticationProvider authentication.AuthenticationProvider, opts ...serviceNowServiceClientOption) (*ServiceNowServiceClient, error) {

	// TODO: make credential to request adapter conversion
	requestAdapter, err := NewServiceNowRequestAdapterBase(authenticationProvider)
	if err != nil {
		return nil, err
	}

	return &ServiceNowServiceClient{
		ServiceNowBaseServiceClient: *NewServiceNowBaseServiceClient(requestAdapter, store.BackingStoreFactoryInstance),
	}, nil
}
