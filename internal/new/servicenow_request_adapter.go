package internal

import (
	"errors"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

// ServiceNowRequestAdapter is the core service used by ServiceNowServiceClient to make requests to ServiceNow APIs.
type ServiceNowRequestAdapter struct {
	abstractions.RequestAdapter
}

// NewServiceNowRequestAdapter creates a new ServiceNowRequestAdapter using the provided authenticationProvider and options.
func NewServiceNowRequestAdapter(authenticationProvider authentication.AuthenticationProvider, opts ...ServiceNowRequestAdapterOption) (*ServiceNowRequestAdapter, error) {
	if authenticationProvider == nil {
		return nil, errors.New("authenticationProvider cannot be nil")
	}

	config, err := buildServiceNowRequestAdapterConfig(opts...)
	if err != nil {
		return nil, err
	}

	baseAdapter, err := nethttplibrary.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(authenticationProvider, config.parseNodeFactory, config.serializationWriterFactory, config.client)
	if err != nil {
		// can't test
		return nil, err
	}
	return &ServiceNowRequestAdapter{baseAdapter}, nil
}
