package servicenowsdkgo

import (
	"errors"
	"strings"
)

const (
	defaultServiceNowHost = "service-now.com"
)

type serviceNowServiceClientConfig struct {
	instance                        string
	rawURL                          string
	serviceNowRequestAdapterOptions []serviceNowRequestAdapterOption
}

func (c serviceNowServiceClientConfig) validate() error {
	if c.instance != "" && c.rawURL != "" {
		return errors.New("can't use both WithInstance and WithRawURL")
	}

	if c.instance == "" && c.rawURL == "" {
		return errors.New("must use either WithInstance or WithRawURL")
	}

	return nil
}

type serviceNowServiceClientOption func(*serviceNowServiceClientConfig) error

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
