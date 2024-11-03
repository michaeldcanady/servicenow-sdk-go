package servicenowsdkgo

import (
	"net/http"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

type serviceNowClientConfig struct {
	middleware []nethttplibrary.Middleware
}

type serviceNowClientOption func(*serviceNowClientConfig) error

func getDefaultMiddleware() []nethttplibrary.Middleware {
	kiotaMiddlewares := nethttplibrary.GetDefaultMiddlewares()

	graphMiddlewares := []nethttplibrary.Middleware{
		//NewGraphTelemetryHandler(options),
	}
	graphMiddlewaresLen := len(graphMiddlewares)
	resultMiddlewares := make([]nethttplibrary.Middleware, len(kiotaMiddlewares)+graphMiddlewaresLen)
	copy(resultMiddlewares, graphMiddlewares)
	copy(resultMiddlewares[graphMiddlewaresLen:], kiotaMiddlewares)
	return resultMiddlewares
}

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

func WithMiddleware(middleware nethttplibrary.Middleware) serviceNowClientOption {
	return func(config *serviceNowClientConfig) error {
		if internal.IsNil(config.middleware) {
			config.middleware = []nethttplibrary.Middleware{}
		}
		config.middleware = append(config.middleware, middleware)
		return nil
	}
}
