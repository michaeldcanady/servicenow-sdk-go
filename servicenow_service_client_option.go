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
	host                            string
	serviceNowRequestAdapterOptions []serviceNowRequestAdapterOption
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

func WithHost(host string) serviceNowServiceClientOption {
	host = strings.TrimSpace(host)
	if host == "" {
		return func(_ *serviceNowServiceClientConfig) error {
			return errors.New("WithHost host can't be nil")
		}
	}

	return func(config *serviceNowServiceClientConfig) error {
		config.host = host
		return nil
	}
}

func WithServiceNowRequestAdapterOptions(opts ...serviceNowRequestAdapterOption) serviceNowServiceClientOption {
	if opts == nil || len(opts) == 0 {
		return func(_ *serviceNowServiceClientConfig) error {
			return errors.New("WithServiceNowRequestAdapterOptions opts can't be nil or empty")
		}
	}

	return func(config *serviceNowServiceClientConfig) error {
		config.serviceNowRequestAdapterOptions = opts
		return nil
	}
}
