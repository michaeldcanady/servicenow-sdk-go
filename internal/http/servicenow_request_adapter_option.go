package internalhttp

import (
	"errors"
	"net/http"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

// ServiceNowRequestAdapterOption represents an option for the ServiceNowRequestAdapter
type ServiceNowRequestAdapterOption func(*serviceNowRequestAdapterConfig) error

// WithClient provides a http.Client to be used by the ServiceNowRequestAdapter
func WithClient(client *http.Client) ServiceNowRequestAdapterOption {
	return func(config *serviceNowRequestAdapterConfig) error {
		if conversion.IsNil(client) {
			return errors.New("client is nil")
		}

		config.client = client
		return nil
	}
}

// WithParseNodeFactory provides a ParseNodeFactory to be used by the ServiceNowRequestAdapter
func WithParseNodeFactory(factory serialization.ParseNodeFactory) ServiceNowRequestAdapterOption {
	return func(config *serviceNowRequestAdapterConfig) error {
		if conversion.IsNil(factory) {
			return snerrors.ErrNilFactory
		}
		config.parseNodeFactory = factory
		return nil
	}
}

// WithServiceNowClientOptions provides options to create a ServiceNowClient to use a the http.Client
func WithServiceNowClientOptions(opts ...ServiceNowClientOption) ServiceNowRequestAdapterOption {
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
		if conversion.IsNil(factory) {
			return snerrors.ErrNilFactory
		}
		config.serializationWriterFactory = factory
		return nil
	}
}

// WithLogger provides an internal.Logger that receives request, response,
// retry and redirect diagnostics from the default client pipeline. It has no
// effect when a custom http.Client is supplied via WithClient.
func WithLogger(logger internal.Logger) ServiceNowRequestAdapterOption {
	return func(config *serviceNowRequestAdapterConfig) error {
		if conversion.IsNil(config) {
			return snerrors.ErrNilConfig
		}
		if conversion.IsNil(logger) {
			return errors.New("logger is nil")
		}

		config.logger = logger

		return nil
	}
}

// serviceNowRequestAdapterDefaultOptions configures default options if an option is not supplied for the ServiceNowRequestAdapterConfig
func serviceNowRequestAdapterDefaultOptions() ServiceNowRequestAdapterOption {
	return func(config *serviceNowRequestAdapterConfig) error {
		if conversion.IsNil(config.client) {
			config.client = nethttplibrary.GetDefaultClient(defaultMiddlewareWithLogger(config.logger)...)
		} else if len(config.middleware) > 0 {
			config.client = nethttplibrary.GetDefaultClient(config.middleware...)
		}

		if conversion.IsNil(config.serializationWriterFactory) {
			config.serializationWriterFactory = serialization.DefaultSerializationWriterFactoryInstance
		}
		if conversion.IsNil(config.parseNodeFactory) {
			config.parseNodeFactory = serialization.DefaultParseNodeFactoryInstance
		}
		return nil
	}
}
