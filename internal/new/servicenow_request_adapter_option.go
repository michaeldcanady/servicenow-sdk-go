package internal

import (
	"errors"
	"net/http"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceNowRequestAdapterOption represents an option for the ServiceNowRequestAdapter
type ServiceNowRequestAdapterOption func(*serviceNowRequestAdapterConfig) error

// WithClient provides a http.Client to be used by the ServiceNowRequestAdapter
func WithClient(client *http.Client) ServiceNowRequestAdapterOption {
	return func(config *serviceNowRequestAdapterConfig) error {
		if IsNil(client) {
			return errors.New("client is nil")
		}

		config.client = client
		return nil
	}
}

// WithParseNodeFactory provides a ParseNodeFactory to be used by the ServiceNowRequestAdapter
func WithParseNodeFactory(factory serialization.ParseNodeFactory) ServiceNowRequestAdapterOption {
	return func(config *serviceNowRequestAdapterConfig) error {
		if IsNil(factory) {
			return errors.New("factory is nil")
		}
		config.parseNodeFactory = factory
		return nil
	}
}

// WithServiceNowClientOptions provides options to create a ServiceNowClient to use a the http.Client
func WithServiceNowClientOptions(opts ...serviceNowClientOption) ServiceNowRequestAdapterOption {
	client, err := GetDefaultClient(opts...)
	if err != nil {
		return func(_ *serviceNowRequestAdapterConfig) error {
			return err
		}
	}

	return WithClient(client)
}

// WithSerializationFactory provides a SerializationFactory to be used by the ServiceNowRequestAdapter
func WithSerializationFactory(factory serialization.SerializationWriterFactory) ServiceNowRequestAdapterOption {
	return func(config *serviceNowRequestAdapterConfig) error {
		if IsNil(factory) {
			return errors.New("factory is nil")
		}
		config.serializationWriterFactory = factory
		return nil
	}
}

// serviceNowRequestAdapterDefaultOptions configures default options if an option is not supplied for the ServiceNowRequestAdapterConfig
func serviceNowRequestAdapterDefaultOptions() ServiceNowRequestAdapterOption {
	return func(config *serviceNowRequestAdapterConfig) error {
		if IsNil(config.client) {
			client, err := GetDefaultClient()
			// can't test since an error can't be forced
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
