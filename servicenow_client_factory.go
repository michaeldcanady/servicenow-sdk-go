package servicenowsdkgo

import (
	"net/http"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

// serviceNowClientConfig represents configurations for the ServiceNowClient
type serviceNowClientConfig struct {
	// middleware middleware used by client
	middleware []nethttplibrary.Middleware
}

// serviceNowClientOption represent options for the ServiceNowClient
type serviceNowClientOption func(*serviceNowClientConfig) error

// getDefaultMiddleware returns a slice of the default middleware for the ServiceNowClient
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

// GetDefaultClient creates new default http.Client using the provided options
func GetDefaultClient(opts ...serviceNowClientOption) (*http.Client, error) {
	config := serviceNowClientConfig{
		middleware: []nethttplibrary.Middleware{},
	}

	for _, opt := range opts {
		if err := opt(&config); err != nil {
			return nil, err
		}
	}

	if len(config.middleware) == 0 {
		config.middleware = getDefaultMiddleware()
	}

	return nethttplibrary.GetDefaultClient(config.middleware...), nil
}

// WithMiddleware represents client option to apply desired middleware
func WithMiddleware(middleware nethttplibrary.Middleware) serviceNowClientOption {
	return func(config *serviceNowClientConfig) error {
		if internal.IsNil(config.middleware) {
			config.middleware = []nethttplibrary.Middleware{}
		}
		config.middleware = append(config.middleware, middleware)
		return nil
	}
}
