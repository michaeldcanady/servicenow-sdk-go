package servicenowsdkgo

import (
	"context"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type ServiceNowServiceClient struct {
	ServiceNowBaseServiceClient
}

// TODO: will be new base for auth provider (credentials)
type serviceNowAuthorizationProvider interface {
	GetAuthorization(context.Context, *url.URL, map[string]interface{}) (string, error)
}

func NewServiceNowServiceClient2(credentials internal.Credential, opts ...serviceNowServiceClientOption) (*ServiceNowServiceClient, error) {

	// TODO: make credential to request adapter convertion
	requestAdapter, err := NewServiceNowRequestAdapterBase()
	if err != nil {
		return nil, err
	}

	return &ServiceNowServiceClient{
		ServiceNowBaseServiceClient: *NewServiceNowBaseServiceClient(requestAdapter, store.BackingStoreFactoryInstance),
	}, nil
}
