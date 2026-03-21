package servicenowsdkgo

import (
	"errors"
	"strings"

	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/microsoft/kiota-abstractions-go/store"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

// ServiceNowServiceClientConfig represents possible configurations for the Service-Now service client.
type ServiceNowServiceClientConfig struct {
	authenticationProvider authentication.AuthenticationProvider
	requestAdapter         abstractions.RequestAdapter
	middleware             []nethttplibrary.Middleware
	rawURI                 string
	backingStoreFactory    store.BackingStoreFactory
	requestAdapterOptions  []internalHttp.ServiceNowRequestAdapterOption
}

// newServiceNowClientConfig instantiates a new empty config with default values.
func newServiceNowClientConfig() *ServiceNowServiceClientConfig {
	return &ServiceNowServiceClientConfig{
		middleware:            make([]nethttplibrary.Middleware, 0),
		backingStoreFactory:   store.BackingStoreFactoryInstance,
		requestAdapterOptions: make([]internalHttp.ServiceNowRequestAdapterOption, 0),
	}
}

// buildServiceClientConfig assembles new client config using the provided options and validates it.
func buildServiceClientConfig(opts ...ServiceNowServiceClientOption) (*ServiceNowServiceClientConfig, error) {
	config := newServiceNowClientConfig()
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}

// Validate checks if the configuration is valid.
func (c *ServiceNowServiceClientConfig) Validate() error {
	if utils.IsNil(c.requestAdapter) && utils.IsNil(c.authenticationProvider) {
		return errors.New("must provide either an AuthenticationProvider or a RequestAdapter")
	}

	return nil
}

// getBaseURL resolves the base URL from the configuration.
func (c *ServiceNowServiceClientConfig) getBaseURL() string {
	return strings.TrimSpace(c.rawURI)
}

// getRequestAdapter resolves the request adapter and enables the backing store if configured.
func (c *ServiceNowServiceClientConfig) getRequestAdapter() (abstractions.RequestAdapter, error) {
	requestAdapter := c.requestAdapter
	var err error

	if utils.IsNil(requestAdapter) {
		requestAdapter, err = internalHttp.NewServiceNowRequestAdapter(c.authenticationProvider, c.requestAdapterOptions...)
		if err != nil {
			return nil, err
		}
	}

	if !utils.IsNil(c.backingStoreFactory) {
		requestAdapter.EnableBackingStore(c.backingStoreFactory)
	}

	baseURL := c.getBaseURL()
	requestAdapter.SetBaseUrl(baseURL)

	type preparable interface {
		Initialize(baseURL string)
	}

	if p, ok := c.authenticationProvider.(preparable); ok {
		p.Initialize(baseURL)
	}

	return requestAdapter, nil
}
