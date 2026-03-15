package servicenowsdkgo

import (
	"errors"
	"net/http"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/store"
	nethttplibrary "github.com/microsoft/kiota-http-go"
	"github.com/stretchr/testify/assert"
)

func TestWithAuthenticationProvider(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				authProvider := mocking.NewMockAuthenticationProvider()
				option := WithAuthenticationProvider(authProvider)
				config := &ServiceNowServiceClientConfig{}
				err := option(config)
				assert.Nil(t, err)
				assert.Equal(t, authProvider, config.authenticationProvider)
			},
		},
		{
			name: "nil auth",
			test: func(t *testing.T) {
				option := WithAuthenticationProvider(nil)
				config := &ServiceNowServiceClientConfig{}
				err := option(config)
				assert.Equal(t, errors.New("authenticationProvider is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestWithRequestAdapter(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				adapter := mocking.NewMockRequestAdapter()
				option := WithRequestAdapter(adapter)
				config := &ServiceNowServiceClientConfig{}
				err := option(config)
				assert.Nil(t, err)
				assert.Equal(t, adapter, config.requestAdapter)
			},
		},
		{
			name: "nil adapter",
			test: func(t *testing.T) {
				option := WithRequestAdapter(nil)
				config := &ServiceNowServiceClientConfig{}
				err := option(config)
				assert.Equal(t, errors.New("requestAdapter is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestWithURL(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				input := "https://exampleurl.com"
				option := WithURL(input)
				config := &ServiceNowServiceClientConfig{}
				err := option(config)
				assert.Nil(t, err)
				assert.Equal(t, input, config.rawURI)
			},
		},
		{
			name: "empty uri",
			test: func(t *testing.T) {
				input := " "
				option := WithURL(input)
				config := &ServiceNowServiceClientConfig{}
				err := option(config)
				assert.Equal(t, errors.New("url is empty"), err)
				assert.Equal(t, "", config.rawURI)
			},
		},
		{
			name: "invalid uri",
			test: func(t *testing.T) {
				input := "https://example url.com"
				option := WithURL(input)
				config := &ServiceNowServiceClientConfig{}
				err := option(config)
				assert.Equal(t, errors.New("parse \"https://example url.com\": invalid character \" \" in host name"), err)
				assert.Equal(t, "", config.rawURI)
			},
		},
		{
			name: "nil config",
			test: func(t *testing.T) {
				input := "https://exampleurl.com"
				option := WithURL(input)
				config := (*ServiceNowServiceClientConfig)(nil)
				err := option(config)
				assert.Equal(t, errors.New("config is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestWithMiddleware(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				config := &ServiceNowServiceClientConfig{}
				middleware := mocking.NewMockMiddleware()

				option := WithMiddleware(middleware)

				err := option(config)

				assert.Nil(t, err)
				assert.Equal(t, middleware, config.middleware[0])
			},
		},
		{
			name: "nil config",
			test: func(t *testing.T) {
				middleware := mocking.NewMockMiddleware()
				config := (*ServiceNowServiceClientConfig)(nil)

				option := WithMiddleware(middleware)
				err := option(config)
				assert.Equal(t, errors.New("config is nil"), err)
			},
		},
		{
			name: "no middleware",
			test: func(t *testing.T) {
				config := &ServiceNowServiceClientConfig{}

				option := WithMiddleware()

				err := option(config)

				assert.Equal(t, errors.New("middleware is empty"), err)
				assert.Equal(t, ([]nethttplibrary.Middleware)(nil), config.middleware)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestWithInstance(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				instance := "test"
				config := &ServiceNowServiceClientConfig{}

				option := WithInstance(instance)
				err := option(config)

				assert.Nil(t, err)
				assert.Equal(t, instance, config.instance)
			},
		},
		{
			name: "nil config",
			test: func(t *testing.T) {
				instance := "https://exampleurl.com"
				config := (*ServiceNowServiceClientConfig)(nil)

				option := WithInstance(instance)

				err := option(config)
				assert.Equal(t, errors.New("config is nil"), err)
			},
		},
		{
			name: "empty instance",
			test: func(t *testing.T) {
				instance := " "
				config := &ServiceNowServiceClientConfig{}

				option := WithInstance(instance)
				err := option(config)

				assert.Equal(t, errors.New("instance is empty"), err)
				assert.Equal(t, "", config.instance)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestWithBackingStore(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				backingStore := store.BackingStoreFactoryInstance
				config := &ServiceNowServiceClientConfig{}

				option := WithBackingStoreFactory(backingStore)
				err := option(config)

				assert.Nil(t, err)
				// can't verify they're the same because they're functions
				assert.NotNil(t, config.backingStoreFactory)
			},
		},
		{
			name: "nil factory",
			test: func(t *testing.T) {
				config := &ServiceNowServiceClientConfig{}

				option := WithBackingStoreFactory(nil)
				err := option(config)

				assert.Equal(t, errors.New("backingStoreFactory is nil"), err)
				assert.Nil(t, config.backingStoreFactory)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestWithHTTPClient(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				client := &http.Client{}
				config := &ServiceNowServiceClientConfig{}

				option := WithHTTPClient(client)
				err := option(config)

				assert.Nil(t, err)
				assert.NotEmpty(t, config.requestAdapterOptions)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
