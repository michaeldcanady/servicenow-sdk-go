package servicenowsdkgo

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/store"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

// serviceNowServiceClientConfig represents possible configurations for the Service-Now service client.
type serviceNowServiceClientConfig struct {
	middleware          []nethttplibrary.Middleware
	rawURI              string
	instance            string
	backingStoreFactory store.BackingStoreFactory
}

// newServiceNowClientConfig instantiates a new empty config.
func newServiceNowClientConfig() *serviceNowServiceClientConfig {
	return &serviceNowServiceClientConfig{
		middleware: make([]nethttplibrary.Middleware, 0),
		rawURI:     "",
		instance:   "",
	}
}

// buildServiceClientConfig assembles new client config using the provided options.
func buildServiceClientConfig(opts ...serviceNowServiceClientOption) (*serviceNowServiceClientConfig, error) {
	config := newServiceNowClientConfig()
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	if config.rawURI != "" && config.instance != "" {
		return nil, errors.New("rawURL and instance cannot be used together")
	}

	return config, nil
}
