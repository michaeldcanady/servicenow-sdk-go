package servicenowsdkgo

import (
	"context"
	"errors"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const authorizationHeader = "Authorization"

type ServiceNowServiceClient struct {
	ServiceNowBaseServiceClient
}

// TODO: will be new base for auth provider (credentials)
type credentialAuthorizationProviderAdapter struct {
	cred core.Credential
}

func (provider *credentialAuthorizationProviderAdapter) AuthenticateRequest(ctx context.Context, request *abstractions.RequestInformation, _ map[string]interface{}) error {
	if request == nil {
		return errors.New("request is nil")
	}
	if request.Headers == nil {
		request.Headers = abstractions.NewRequestHeaders()
	}
	if provider.cred == nil {
		return errors.New("this class needs to be initialized with a credential")
	}
	if !request.Headers.ContainsKey(authorizationHeader) {
		authString, err := provider.cred.GetAuthentication()
		if err != nil {
			return err
		}
		if authString != "" {
			request.Headers.Add(authorizationHeader, authString)
		}
	}

	return nil
}

func NewServiceNowServiceClient2WithCredential(credential core.Credential, opts ...serviceNowServiceClientOption) (*ServiceNowServiceClient, error) {
	authenticationProvider := &credentialAuthorizationProviderAdapter{
		cred: credential,
	}

	return NewServiceNowServiceClient2(authenticationProvider, opts...)
}

func NewServiceNowServiceClient2(authenticationProvider authentication.AuthenticationProvider, opts ...serviceNowServiceClientOption) (*ServiceNowServiceClient, error) {
	requestAdapter, err := NewServiceNowRequestAdapterBase(authenticationProvider)
	if err != nil {
		return nil, err
	}

	config, err := buildServiceClientConfig(opts...)
	if err != nil {
		return nil, err
	}

	var baseURL = config.rawURL
	if config.instance != "" {
		baseURL = fmt.Sprintf("https://%s.%s", config.instance, defaultServiceNowHost)
	}

	return &ServiceNowServiceClient{
		ServiceNowBaseServiceClient: *NewServiceNowBaseServiceClient(requestAdapter, store.BackingStoreFactoryInstance, baseURL),
	}, nil
}
