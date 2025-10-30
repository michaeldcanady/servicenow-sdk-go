package internal

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

// serviceNowClientOption represents options for the ServiceNowClient
type serviceNowClientOption func(*serviceNowClientConfig) error

// WithMiddleware adds supplied middleware to the ServiceNowClientConfig
func WithMiddleware(middleware ...nethttplibrary.Middleware) serviceNowClientOption {
	return func(config *serviceNowClientConfig) error {
		if len(middleware) == 0 {
			return errors.New("middleware is empty")
		}
		if internal.IsNil(config) {
			return errors.New("config is nil")
		}
		if internal.IsNil(config.middleware) {
			config.middleware = []nethttplibrary.Middleware{}
		}
		config.middleware = append(config.middleware, middleware...)
		return nil
	}
}

func serviceNowClientDefaultOptions() serviceNowClientOption {
	return func(config *serviceNowClientConfig) error {
		if len(config.middleware) == 0 {
			config.middleware = getDefaultMiddleware()
		}

		return nil
	}
}
