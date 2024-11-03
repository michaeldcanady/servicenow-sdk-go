package servicenowsdkgo

import (
	"errors"
	"net/http"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type serviceNowRequestAdapterConfig struct {
	client                     *http.Client
	serializationWriterFactory serialization.SerializationWriterFactory
	parseNodeFactory           serialization.ParseNodeFactory
}

func buildConfig(opts ...serviceNowRequestAdapterOption) (*serviceNowRequestAdapterConfig, error) {
	opts = append(opts, serviceNowRequestAdapterDefaultOptions())
	config := new(serviceNowRequestAdapterConfig)
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	return config, nil
}

type serviceNowRequestAdapterOption func(*serviceNowRequestAdapterConfig) error

func WithClient(client *http.Client) serviceNowRequestAdapterOption {
	if internal.IsNil(client) {
		return func(config *serviceNowRequestAdapterConfig) error {
			return errors.New("WithClient client can't be nil")
		}
	}

	return func(config *serviceNowRequestAdapterConfig) error {
		config.client = client
		return nil
	}
}

func WithSerializationFactory(factory serialization.SerializationWriterFactory) serviceNowRequestAdapterOption {
	if internal.IsNil(factory) {
		return func(config *serviceNowRequestAdapterConfig) error {
			return errors.New("WithSerializationFactory factory can't be nil")
		}
	}

	return func(config *serviceNowRequestAdapterConfig) error {
		config.serializationWriterFactory = factory
		return nil
	}
}

func WithParseNodeFactory(factory serialization.ParseNodeFactory) serviceNowRequestAdapterOption {
	if internal.IsNil(factory) {
		return func(config *serviceNowRequestAdapterConfig) error {
			return errors.New("WithParseNodeFactory factory can't be nil")
		}
	}

	return func(config *serviceNowRequestAdapterConfig) error {
		config.parseNodeFactory = factory
		return nil
	}
}

func serviceNowRequestAdapterDefaultOptions() serviceNowRequestAdapterOption {
	return func(config *serviceNowRequestAdapterConfig) error {
		if internal.IsNil(config.client) {
			client, err := GetDefaultClient()
			if err != nil {
				return err
			}
			config.client = client
		}
		if internal.IsNil(config.serializationWriterFactory) {
			config.serializationWriterFactory = serialization.DefaultSerializationWriterFactoryInstance
		}
		if internal.IsNil(config.parseNodeFactory) {
			config.parseNodeFactory = serialization.DefaultParseNodeFactoryInstance
		}
		return nil
	}
}
