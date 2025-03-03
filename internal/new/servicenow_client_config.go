package internal

import (
	"net/http"

	nethttplibrary "github.com/microsoft/kiota-http-go"
)

// serviceNowClientConfig represents configurations for a ServiceNowClient
type serviceNowClientConfig struct {
	// middleware the middleware for the client
	middleware []nethttplibrary.Middleware
}

// buildServiceNowClientConfig constructs new serviceNowClientConfig from provided options
func buildServiceNowClientConfig(opts ...serviceNowClientOption) (*serviceNowClientConfig, error) {
	opts = append(opts, serviceNowClientDefaultOptions())

	config := new(serviceNowClientConfig)

	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	return config, nil
}

// getDefaultMiddleware returns the default middleware for a ServiceNowClient
func getDefaultMiddleware() []nethttplibrary.Middleware {
	kiotaMiddlewares := nethttplibrary.GetDefaultMiddlewares()

	serviceNowMiddlewares := []nethttplibrary.Middleware{
		//NewGraphTelemetryHandler(options),
		nethttplibrary.NewHeadersInspectionHandler(),
	}
	graphMiddlewaresLen := len(serviceNowMiddlewares)
	resultMiddlewares := make([]nethttplibrary.Middleware, len(kiotaMiddlewares)+graphMiddlewaresLen)
	copy(resultMiddlewares, serviceNowMiddlewares)
	copy(resultMiddlewares[graphMiddlewaresLen:], kiotaMiddlewares)
	return resultMiddlewares
}

// GetDefaultClient constructs default client using provided options
func GetDefaultClient(opts ...serviceNowClientOption) (*http.Client, error) {
	config, err := buildServiceNowClientConfig(opts...)
	if err != nil {
		return nil, err
	}

	return nethttplibrary.GetDefaultClient(config.middleware...), nil
}
