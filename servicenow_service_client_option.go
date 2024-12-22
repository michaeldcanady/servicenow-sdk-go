package servicenowsdkgo

import (
	"errors"
	"strings"
)

const (
	// defaultServiceNowHost the default host for the Service-Now APIs
	defaultServiceNowHost = "service-now.com"
)

// serviceNowServiceClientConfig represents possible configurations for the ServiceNowServiceClient
type serviceNowServiceClientConfig struct {
	instance                        string
	rawURL                          string
	serviceNowRequestAdapterOptions []serviceNowRequestAdapterOption
}

// validate validates the configurations and returns an error if inappropriately set
func (c serviceNowServiceClientConfig) validate() error {
	if c.instance != "" && c.rawURL != "" {
		return errors.New("can't use both WithInstance and WithRawURL")
	}

	if c.instance == "" && c.rawURL == "" {
		return errors.New("must use either WithInstance or WithRawURL")
	}

	return nil
}

// serviceNowServiceClientOption represents possible options for the ServiceNowServiceClient
type serviceNowServiceClientOption func(*serviceNowServiceClientConfig) error

// WithInstance configures ServiceNowServiceClient to use the provide instance with the defaultServiceNowHost
func WithInstance(instance string) serviceNowServiceClientOption {
	instance = strings.TrimSpace(instance)
	if instance == "" {
		return func(_ *serviceNowServiceClientConfig) error {
			return errors.New("WithInstance instance can't be nil")
		}
	}

	return func(config *serviceNowServiceClientConfig) error {
		config.instance = instance
		return nil
	}
}

// WithRawURL configures ServiceNowServiceClient to use a custom host address, instead of the default
func WithRawURL(rawURL string) serviceNowServiceClientOption {
	rawURL = strings.TrimSpace(rawURL)
	if rawURL == "" {
		return func(_ *serviceNowServiceClientConfig) error {
			return errors.New("WithHost host can't be nil")
		}
	}

	return func(config *serviceNowServiceClientConfig) error {
		config.rawURL = rawURL
		return nil
	}
}

// WithServiceNowRequestAdapterOptions provides the ability to provide ServiceNowRequestAdapterOptions to the internal ServiceNowRequestAdapter
func WithServiceNowRequestAdapterOptions(opts ...serviceNowRequestAdapterOption) serviceNowServiceClientOption {
	if len(opts) == 0 {
		return func(_ *serviceNowServiceClientConfig) error {
			return errors.New("WithServiceNowRequestAdapterOptions opts can't be nil or empty")
		}
	}

	return func(config *serviceNowServiceClientConfig) error {
		config.serviceNowRequestAdapterOptions = opts
		return nil
	}
}

// buildServiceClientConfig creates a ServiceNowClientConfig using the provided options
func buildServiceClientConfig(opts ...serviceNowServiceClientOption) (*serviceNowServiceClientConfig, error) {
	config := &serviceNowServiceClientConfig{}

	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}
