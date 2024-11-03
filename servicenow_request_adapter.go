package servicenowsdkgo

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/authentication"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

// ServiceNowRequestAdapter is the core service used by ServiceNowServiceClient to make requests to ServiceNow APIs.
type ServiceNowRequestAdapter struct {
	nethttplibrary.NetHttpRequestAdapter
}

func NewServiceNowRequestAdapterBase(authenticationProvider authentication.AuthenticationProvider, opts ...serviceNowRequestAdapterOption) (*ServiceNowRequestAdapter, error) {
	if authenticationProvider == nil {
		return nil, errors.New("authenticationProvider cannot be nil")
	}

	config, err := buildConfig(opts...)
	if err != nil {
		return nil, err
	}

	baseAdapter, err := nethttplibrary.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(authenticationProvider, config.parseNodeFactory, config.serializationWriterFactory, config.client)
	if err != nil {
		return nil, err
	}
	return &ServiceNowRequestAdapter{NetHttpRequestAdapter: *baseAdapter}, nil
}
