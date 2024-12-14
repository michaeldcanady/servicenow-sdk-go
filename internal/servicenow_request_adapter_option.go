package internal

import (
	"errors"
	"net/http"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type serviceNowRequestAdapterConfig struct {
	client                     *http.Client
	serializationWriterFactory serialization.SerializationWriterFactory
	parseNodeFactory           serialization.ParseNodeFactory
}

func buildConfig(opts ...ServiceNowRequestAdapterOption) (*serviceNowRequestAdapterConfig, error) {
	opts = append(opts, serviceNowRequestAdapterDefaultOptions())
	config := new(serviceNowRequestAdapterConfig)
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	return config, nil
}

type ServiceNowRequestAdapterOption func(*serviceNowRequestAdapterConfig) error

func WithClient(client *http.Client) ServiceNowRequestAdapterOption {
	if IsNil(client) {
		return func(config *serviceNowRequestAdapterConfig) error {
			return errors.New("WithClient client can't be nil")
		}
	}

	return func(config *serviceNowRequestAdapterConfig) error {
		config.client = client
		return nil
	}
}

func WithSerializationFactory(factory serialization.SerializationWriterFactory) ServiceNowRequestAdapterOption {
	if IsNil(factory) {
		return func(config *serviceNowRequestAdapterConfig) error {
			return errors.New("WithSerializationFactory factory can't be nil")
		}
	}

	return func(config *serviceNowRequestAdapterConfig) error {
		config.serializationWriterFactory = factory
		return nil
	}
}

func WithParseNodeFactory(factory serialization.ParseNodeFactory) ServiceNowRequestAdapterOption {
	if IsNil(factory) {
		return func(config *serviceNowRequestAdapterConfig) error {
			return errors.New("WithParseNodeFactory factory can't be nil")
		}
	}

	return func(config *serviceNowRequestAdapterConfig) error {
		config.parseNodeFactory = factory
		return nil
	}
}

func serviceNowRequestAdapterDefaultOptions() ServiceNowRequestAdapterOption {
	return func(config *serviceNowRequestAdapterConfig) error {
		if IsNil(config.client) {
			client, err := GetDefaultClient()
			if err != nil {
				return err
			}
			config.client = client
		}
		if IsNil(config.serializationWriterFactory) {
			config.serializationWriterFactory = serialization.DefaultSerializationWriterFactoryInstance
		}
		if IsNil(config.parseNodeFactory) {
			config.parseNodeFactory = serialization.DefaultParseNodeFactoryInstance
		}
		return nil
	}
}
