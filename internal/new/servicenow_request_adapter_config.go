package internal

import (
	"net/http"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// serviceNowRequestAdapterConfig represents configurations for ServiceNowRequestAdapter
type serviceNowRequestAdapterConfig struct {
	// client the HTTP client for the adapter
	client *http.Client
	// serializationWriterFactory the SerializationWriterFactory for the adapter
	serializationWriterFactory serialization.SerializationWriterFactory
	// parseNodeFactory the ParseNodeFactory for the adapter
	parseNodeFactory serialization.ParseNodeFactory
}

// buildServiceNowRequestAdapterConfig constructs new serviceNowRequestAdapterConfig from provided options
func buildServiceNowRequestAdapterConfig(opts ...ServiceNowRequestAdapterOption) (*serviceNowRequestAdapterConfig, error) {
	opts = append(opts, serviceNowRequestAdapterDefaultOptions())
	config := new(serviceNowRequestAdapterConfig)
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	return config, nil
}
