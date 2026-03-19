package servicenowsdkgo

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/microsoft/kiota-abstractions-go/store"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

// ServiceNowServiceClientOption is a function type that modifies the ServiceNowServiceClientConfig.
// It returns an error if the modification is not successful.
type ServiceNowServiceClientOption func(*ServiceNowServiceClientConfig) error

// WithAuthenticationProvider sets the authentication provider for the ServiceNowServiceClient.
func WithAuthenticationProvider(authenticationProvider authentication.AuthenticationProvider) ServiceNowServiceClientOption {
	return func(config *ServiceNowServiceClientConfig) error {
		if internal.IsNil(authenticationProvider) {
			return errors.New("authenticationProvider is nil")
		}
		config.authenticationProvider = authenticationProvider
		return nil
	}
}

// WithRequestAdapter sets a pre-configured RequestAdapter for the ServiceNowServiceClient.
func WithRequestAdapter(requestAdapter abstractions.RequestAdapter) ServiceNowServiceClientOption {
	return func(config *ServiceNowServiceClientConfig) error {
		if internal.IsNil(requestAdapter) {
			return errors.New("requestAdapter is nil")
		}
		config.requestAdapter = requestAdapter
		return nil
	}
}

// WithURL creates an option to set the base URL for the requests.
// It returns an error if the provided configuration is nil or the URL is empty.
func WithURL(uri string) ServiceNowServiceClientOption {
	return func(config *ServiceNowServiceClientConfig) error {
		if internal.IsNil(config) {
			return errors.New("config is nil")
		}
		uri = strings.TrimSpace(uri)
		if uri == "" {
			return errors.New("url is empty")
		}

		if _, err := url.ParseRequestURI(uri); err != nil {
			return fmt.Errorf("%s", err)
		}

		config.rawURI = uri

		return nil
	}
}

// WithMiddleware creates an option to set the middleware used by the requests.
// It returns an error if the provided configuration is nil or the middleware slice is empty.
func WithMiddleware(middleware ...nethttplibrary.Middleware) ServiceNowServiceClientOption {
	return func(config *ServiceNowServiceClientConfig) error {
		if internal.IsNil(config) {
			return errors.New("config is nil")
		}
		if len(middleware) == 0 {
			return errors.New("middleware is empty")
		}

		config.middleware = append(config.middleware, middleware...)

		return nil
	}
}

// WithInstance creates an option to set the instance of the default ServiceNow URL for the requests.
// It returns an error if the provided configuration is nil or the instance string is empty.
func WithInstance(instance string) ServiceNowServiceClientOption {
	instance = strings.TrimSpace(instance)
	if instance == "" {
		return func(_ *ServiceNowServiceClientConfig) error {
			return errors.New("instance is empty")
		}
	}
	return WithURL(fmt.Sprintf("https://%s.%s", instance, defaultServiceNowHost))
}

// WithBackingStoreFactory creates an option to set the backingStoreFactory for the ServiceNowServiceClient.
// It returns an error if the provided factory is nil.
func WithBackingStoreFactory(backingStoreFactory store.BackingStoreFactory) ServiceNowServiceClientOption {
	return func(config *ServiceNowServiceClientConfig) error {
		if internal.IsNil(backingStoreFactory) {
			return errors.New("backingStoreFactory is nil")
		}

		config.backingStoreFactory = backingStoreFactory

		return nil
	}
}

// WithHTTPClient creates an option to set the HTTP client used by the requests.
func WithHTTPClient(client *http.Client) ServiceNowServiceClientOption {
	return func(config *ServiceNowServiceClientConfig) error {
		config.requestAdapterOptions = append(config.requestAdapterOptions, internalHttp.WithClient(client))

		return nil
	}
}
