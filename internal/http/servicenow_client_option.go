package internalhttp

import (
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

// ServiceNowClientOption represents options for the ServiceNowClient
type ServiceNowClientOption func(*serviceNowClientConfig) error

// WithMiddleware adds supplied middleware to the ServiceNowClientConfig
func WithMiddleware(middleware ...nethttplibrary.Middleware) ServiceNowClientOption {
	return func(config *serviceNowClientConfig) error {
		if len(middleware) == 0 {
			return snerrors.ErrEmptyMiddleware
		}
		if conversion.IsNil(config) {
			return snerrors.ErrNilConfig
		}
		if conversion.IsNil(config.middleware) {
			config.middleware = []nethttplibrary.Middleware{}
		}
		config.middleware = append(config.middleware, middleware...)
		return nil
	}
}

func serviceNowClientDefaultOptions() ServiceNowClientOption {
	return func(config *serviceNowClientConfig) error {
		if len(config.middleware) == 0 {
			config.middleware = getDefaultMiddleware()
		}

		return nil
	}
}
