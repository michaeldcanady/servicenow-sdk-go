package servicenowsdkgo

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/store"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

// serviceNowServiceClientOption is a function type that modifies the serviceNowServiceClientConfig.
// It returns an error if the modification is not successful.
type serviceNowServiceClientOption func(*serviceNowServiceClientConfig) error

// withURL creates an option to set the base URL for the requests.
// It returns an error if the provided configuration is nil or the URL is empty.
func withURL(uri string) serviceNowServiceClientOption {
	return func(config *serviceNowServiceClientConfig) error {
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

// withMiddleware creates an option to set the middleware used by the requests.
// It returns an error if the provided configuration is nil or the middleware slice is empty.
func withMiddleware(middleware ...nethttplibrary.Middleware) serviceNowServiceClientOption {
	return func(config *serviceNowServiceClientConfig) error {
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

// withInstance creates an option to set the instance of the default ServiceNow URL for the requests.
// It returns an error if the provided configuration is nil or the instance string is empty.
func withInstance(instance string) serviceNowServiceClientOption {
	return func(config *serviceNowServiceClientConfig) error {
		if internal.IsNil(config) {
			return errors.New("config is nil")
		}
		instance = strings.TrimSpace(instance)
		if instance == "" {
			return errors.New("instance is empty")
		}

		config.instance = instance

		return nil
	}
}

func withBackingStore(backingStoreFactory store.BackingStoreFactory) serviceNowServiceClientOption {
	return func(config *serviceNowServiceClientConfig) error {
		if internal.IsNil(backingStoreFactory) {
			return errors.New("backingStoreFactory is nil")
		}

		config.backingStoreFactory = backingStoreFactory

		return nil
	}
}
